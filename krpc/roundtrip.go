package krpc

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
)

type RoundTrip struct {
	reqId atomic.Uint32

	locker   sync.Mutex
	requests map[uint32]chan Msg
}

func NewRoundTrip() *RoundTrip {
	r := &RoundTrip{
		requests: make(map[uint32]chan Msg),
	}
	return r
}

func (r *RoundTrip) Request(ctx context.Context, conn Conn, msg Msg) (Msg, error) {
	reqId := r.reqId.Add(1)
	msg.setFlagRequest()
	msg.setReqId(reqId)

	rc := make(chan Msg, 1)
	r.locker.Lock()
	r.requests[reqId] = rc
	r.locker.Unlock()

	if err := conn.Send(msg); err != nil {
		r.locker.Lock()
		delete(r.requests, reqId)
		r.locker.Unlock()
		return nil, err
	}

	select {
	case reply := <-rc:
		return reply, nil
	case <-ctx.Done():
		r.locker.Lock()
		delete(r.requests, reqId)
		r.locker.Unlock()
		return nil, ctx.Err()
	}
}

func (r *RoundTrip) HandleReply(msg Msg) error {
	r.locker.Lock()
	rc, ok := r.requests[msg.ReqId()]
	if !ok {
		r.locker.Unlock()
		return errors.New("invalid reply detected")
	}
	delete(r.requests, msg.ReqId())
	r.locker.Unlock()
	rc <- msg
	return nil
}
