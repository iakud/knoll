package actor

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
)

var (
	ErrMailboxStopped = errors.New("Mailbox is stopped")
)

func newMailbox() *mailbox {
	return &mailbox{
		messageC: make(chan Envelope, 1024),
		done:     make(chan struct{}, 1),
	}
}

type mailbox struct {
	messageC chan Envelope
	closed   atomic.Bool
	done     chan struct{}
}

func (m *mailbox) Close() {
	if !m.closed.CompareAndSwap(false, true) {
		return
	}
	close(m.done)
}

func (m *mailbox) Send(ctx context.Context, envelope Envelope) error {
	if m.closed.Load() {
		return fmt.Errorf("Mailbox.Send failed: %w", ErrMailboxStopped)
	}

	select {
	case <-ctx.Done():
		return fmt.Errorf("Mailbox.Send canceled: %w", ctx.Err())
	case m.messageC <- envelope:
		return nil
	}
}

func (m *mailbox) MessageC() <-chan Envelope {
	return m.messageC
}

func (m *mailbox) Done() <-chan struct{} {
	return m.done
}
