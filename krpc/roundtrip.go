package krpc

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
)

type roundTrip struct {
	reqId atomic.Uint32

	locker   sync.Mutex
	requests map[uint32]chan Message
}

func newRoundTrip() *roundTrip {
	r := &roundTrip{
		requests: make(map[uint32]chan Message),
	}
	return r
}

func (r *roundTrip) request(ctx context.Context, conn Conn, m Message) (Message, error) {
	reqId := r.reqId.Add(1)
	m.Header().setFlagRequest()
	m.Header().setReqId(reqId)

	rc := make(chan Message, 1)
	r.locker.Lock()
	r.requests[reqId] = rc
	r.locker.Unlock()

	if err := conn.Send(m); err != nil {
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

func (r *roundTrip) handleReply(m Message) error {
	r.locker.Lock()
	rc, ok := r.requests[m.Header().ReqId()]
	if !ok {
		r.locker.Unlock()
		return errors.New("invalid reply detected")
	}
	delete(r.requests, m.Header().ReqId())
	r.locker.Unlock()
	rc <- m
	return nil
}
