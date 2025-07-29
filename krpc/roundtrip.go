package krpc

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
)

type RoundTrip struct {
	reqId atomic.Uint32

	lock     sync.Mutex
	requests map[uint32]chan Msg
}

func NewRoundTrip() *RoundTrip {
	r := &RoundTrip{}
	return r
}

func (r *RoundTrip) Request(ctx context.Context, conn Conn, msg Msg) (Msg, error) {
	reqId := r.reqId.Add(1)
	msg.setFlagRequest()
	msg.setReqId(reqId)

	rc := make(chan Msg, 1)
	r.lock.Lock()
	r.requests[reqId] = rc
	r.lock.Unlock()

	if err := Send(conn, msg); err != nil {
		r.lock.Lock()
		delete(r.requests, reqId)
		r.lock.Unlock()
		return nil, err
	}

	select {
	case reply := <-rc:
		return reply, nil
	case <-ctx.Done():
		r.lock.Lock()
		delete(r.requests, reqId)
		r.lock.Unlock()
		return nil, ctx.Err()
	}
}

func (r *RoundTrip) HandleReply(msg Msg) error {
	r.lock.Lock()
	rc, ok := r.requests[msg.ReqId()]
	if !ok {
		r.lock.Unlock()
		return errors.New("invalid reply detected")
	}
	delete(r.requests, msg.ReqId())
	r.lock.Unlock()
	rc <- msg
	return nil
}
