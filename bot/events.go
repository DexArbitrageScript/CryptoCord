package bot

import (
	"context"
	"sync"
)

type EventMessage struct {
	Ticket chan string
	mux    sync.Mutex
	ctx    context.Context
	cancel context.CancelFunc
}

var EventMsg *EventMessage

func NewEvent(ctx context.Context) *EventMessage {
	ctx, cancel := context.WithCancel(ctx)

	EventMsg = &EventMessage{
		Ticket: make(chan string),
		ctx:    ctx,
		cancel: cancel,
	}

	return &EventMessage{
		Ticket: make(chan string),
		ctx:    ctx,
		cancel: cancel,
	}
}

func (e *EventMessage) WriteAutomaticEvent(content string) {
	e.mux.Lock()
	defer e.mux.Unlock()
	select {
	case <-e.ctx.Done():
		return // context canceled
	default:
		e.Ticket <- content
	}
}

func (e *EventMessage) RunAutomaticEvent(f func()) {
	for {
		select {
		case <-e.ctx.Done():
			return
		case <-e.Ticket:
			// Start a new goroutine to execute f() and release EventMessage
			go func() {
				// start task
				f()
				e.mux.Unlock()
			}()
		}
	}
}

func (e *EventMessage) StopAutomaticEvent() {
	e.cancel()
	close(e.Ticket)
}
