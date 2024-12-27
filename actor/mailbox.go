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

func NewMailbox[T any]() *mailbox[T] {
	return &mailbox[T]{
		messageC: make(chan T, 1024),
		done:     make(chan struct{}, 1),
	}
}

type mailbox[T any] struct {
	messageC chan T
	closed   atomic.Bool
	done     chan struct{}
}

func (m *mailbox[T]) Close() {
	if !m.closed.CompareAndSwap(false, true) {
		return
	}
	close(m.done)
}

func (m *mailbox[T]) Send(ctx context.Context, message T) error {
	if m.closed.Load() {
		return fmt.Errorf("Mailbox.Send failed: %w", ErrMailboxStopped)
	}

	select {
	case <-ctx.Done():
		return fmt.Errorf("Mailbox.Send canceled: %w", ctx.Err())
	case m.messageC <- message:
		return nil
	}
}

func (m *mailbox[T]) MessageC() <-chan T {
	return m.messageC
}

func (m *mailbox[T]) Done() <-chan struct{} {
	return m.done
}
