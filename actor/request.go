package actor

import (
	"context"
)

type request chan any

func newRequest() request {
	return make(chan any, 1)
}

func (r request) Start()                       {}
func (r request) Stop()                        {}
func (r request) Shutdown(ctx context.Context) {}

func (r request) Send(message any, _ *PID) {
	select {
	case r <- message:
	default:
	}
}

func (r request) Result(ctx context.Context) (any, error) {
	select {
	case resp := <-r:
		return resp, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
