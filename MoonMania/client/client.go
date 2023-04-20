package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ClientApiRequest struct {
	URL     string
	Method  string
	Headers map[string]string
	Body    interface{}
}

// NewClientAPIRequest constructs an HTTP request using the parameters of ClientAPIRequest
// It returns an http.Response struct, which can be used for streaming data processing in other functions.
func NewClientAPIRequest(req *ClientApiRequest) (*http.Response, error) {
	client := &http.Client{}

	requestBody, err := json.Marshal(req.Body)
	if err != nil {
		return nil, fmt.Errorf("Error marshaling request body: %v", err)
	}
	body := bytes.NewBuffer(requestBody)

	httpRequest, err := http.NewRequest(req.Method, req.URL, body)
	if err != nil {
		return nil, fmt.Errorf("Error creating HTTP request: %v", err)
	}

	for key, value := range req.Headers {
		httpRequest.Header.Set(key, value)
	}

	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("Error making HTTP request: %v", err)
	}
	defer httpResponse.Body.Close()

	return httpResponse, nil
}
