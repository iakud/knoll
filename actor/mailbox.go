package actor

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"runtime"
	"sync/atomic"
)

var (
	ErrMailboxStopped = errors.New("Mailbox is stopped")
)

type Invoker interface {
	Invoke(envelope Envelope)
}

type mailbox struct {
	messageC chan Envelope
	invoker  Invoker
	stopped  atomic.Bool
	ctx      context.Context
	cancel   context.CancelFunc

	head *queueElem                // Used carefully to avoid needing atomics
	tail atomic.Pointer[queueElem] // *queueElem, accessed atomically
}

func newMailbox(invoker Invoker) *mailbox {
	ctx, cancel := context.WithCancel(context.Background())
	m := &mailbox{
		messageC: make(chan Envelope, 1024),
		invoker:  invoker,
		ctx:      ctx,
		cancel:   cancel,
	}
	return m
}

func (m *mailbox) Start() {
	// var i atomic.Int32
	go m.process(m.ctx)
}

func (m *mailbox) process(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			slog.Info("Receive recovered in %v", r)
		}
	}()
	for {
		select {
		case envelope := <-m.messageC:
			m.invoker.Invoke(envelope)
		case <-ctx.Done():
			return
		}
	}
}

func (m *mailbox) Send(ctx context.Context, envelope Envelope) error {
	if m.stopped.Load() {
		return fmt.Errorf("Mailbox.Send failed: %w", ErrMailboxStopped)
	}

	select {
	case <-ctx.Done():
		return fmt.Errorf("Mailbox.Send canceled: %w", ctx.Err())
	case m.messageC <- envelope:
		return nil
	}
}

func (m *mailbox) Stop() {
	if !m.stopped.CompareAndSwap(false, true) {
		return
	}
	m.cancel()
}

type queueElem struct {
	envelope Envelope
	next     atomic.Pointer[queueElem] // *queueElem, accessed atomically
}

func (m *mailbox) run() {
	for {
		head := m.head
		m.invoker.Invoke(head.envelope)

		if !m.next(head) {
			*head = queueElem{}
			return
		}
		*head = queueElem{}
	}

}

func (m *mailbox) next(head *queueElem) bool {
	if m.head = head.next.Load(); m.head != nil {
		return true
	}

	if m.tail.CompareAndSwap(head, nil) {
		return false
	}

	for m.head == nil {
		runtime.Gosched()
		m.head = head.next.Load()
	}
	return true
}
