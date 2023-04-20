package stream

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"encoding/json"
	"fmt"
	"io"
	"sync"
)

type ProcessOptions struct {
	Filter        func(interface{}) bool
	MaxWorkers    int
	BatchSize     int
	UseCompressor bool
	Compression   CompressionType
	BufferPool    *sync.Pool
}

type CompressionType int

const (
	NoCompression CompressionType = iota
	GzipCompression
	ZlibCompression
)

func ProcessStream(r io.Reader, options ProcessOptions, out interface{}) error {
	var err error
	var buffer bytes.Buffer

	// compress data if compression is enabled
	if options.UseCompressor {
		switch options.Compression {
		case GzipCompression:
			// create gzip writer
			gw := gzip.NewWriter(&buffer)
			defer gw.Close()

			// copy data to gzip writer
			_, err = io.Copy(gw, r)
			if err != nil {
				return fmt.Errorf("error compressing data: %s", err)
			}
		case ZlibCompression:
			// create zlib writer
			zw := zlib.NewWriter(&buffer)
			defer zw.Close()

			// copy data to zlib writer
			_, err = io.Copy(zw, r)
			if err != nil {
				return fmt.Errorf("error compressing data: %s", err)
			}
		default:
			// copy data to buffer without compression
			_, err = io.Copy(&buffer, r)
			if err != nil {
				return fmt.Errorf("error copying data to buffer: %s", err)
			}
		}
	} else {
		// copy data to buffer without compression
		_, err = io.Copy(&buffer, r)
		if err != nil {
			return fmt.Errorf("error copying data to buffer: %s", err)
		}
	}

	// create channel to receive results
	ch := make(chan interface{}, options.MaxWorkers*options.BatchSize)

	// create wait group to wait for all workers to finish
	var wg sync.WaitGroup

	// create pool of buffers for workers
	bufferPool := options.BufferPool
	if bufferPool == nil {
		bufferPool = &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		}
	}

	// launch workers
	for i := 0; i < options.MaxWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				// get buffer from pool
				buffer := bufferPool.Get().(*bytes.Buffer)
				buffer.Reset()

				// read batch of data from channel
				batch := make([]interface{}, options.BatchSize)
				for j := 0; j < options.BatchSize; j++ {
					s, ok := <-ch
					if !ok {
						break
					}
					if options.Filter != nil && !options.Filter(s) {
						continue
					}
					batch[j] = s
				}

				// check if channel is closed
				if len(batch) == 0 {
					break
				}

				// marshal batch of data to JSON
				encoder := json.NewEncoder(buffer)
				for _, s := range batch {
					if s != nil {
						if err := encoder.Encode(s); err != nil {
							break
						}
					}
				}

				// check for errors while encoding
				if err != nil {
					break
				}

				// send encoded batch to output
				_, err = buffer.WriteTo(out)
				if err != nil {
					break
				}

				// return buffer to pool
				bufferPool.Put(buffer)
			}
		}()
	}

	// decompress and decode JSON data
	switch options.Compression {
	case GzipCompression:
		// create gzip reader
		gzr, err := gzip.NewReader(buffer)
		if err != nil {
			return fmt.Errorf("error creating gzip reader: %s", err)
		}
		defer gzr.Close()

		// decompress and decode JSON data
		decoder := json.NewDecoder(gzr)
		for {
			if err := decoder.Decode(&out); err == io.EOF {
				break
			} else if err != nil {
				return fmt.Errorf("error decoding response: %s", err)
			}
			// send result to channel
			ch <- out
		}
	case ZlibCompression:
		// create zlib reader
		zr, err := zlib.NewReader(buffer)
		if err != nil {
			return fmt.Errorf("error creating zlib reader: %s", err)
		}
		defer zr.Close()

		// decompress and decode JSON data
		decoder := json.NewDecoder(zr)
		for {
			if err := decoder.Decode(&out); err == io.EOF {
				break
			} else if err != nil {
				return fmt.Errorf("error decoding response: %s", err)
			}
			// send result to channel
			ch <- out
		}
	default:
		// decode JSON data without compression
		decoder := json.NewDecoder(buffer)
		for {
			if err := decoder.Decode(&out); err == io.EOF {
				break
			} else if err != nil {
				return fmt.Errorf("error decoding response: %s", err)
			}
			// send result to channel
			ch <- out
		}
	}
	return nil
}
