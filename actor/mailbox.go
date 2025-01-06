package actor

import (
	"runtime"
	"sync"
	"sync/atomic"
)

var nodes = sync.Pool{New: func() interface{} { return new(queueNode) }}

type queueNode struct {
	envelope Envelope
	next     atomic.Pointer[queueNode]
}

type Invoker interface {
	Invoke(envelope Envelope)
}

type mailbox struct {
	invoker Invoker

	head *queueNode
	tail atomic.Pointer[queueNode]
}

func newMailbox(invoker Invoker) *mailbox {
	m := &mailbox{
		invoker: invoker,
	}
	return m
}

func (m *mailbox) Send(envelope Envelope) {
	node := nodes.Get().(*queueNode)
	node.envelope = envelope
	tail := m.tail.Swap(node)
	if tail != nil {
		tail.next.Store(node)
	} else {
		m.head = node
		go m.process()
	}
}

func (m *mailbox) process() {
	for {
		node := m.head
		m.invoker.Invoke(node.envelope)

		if m.head = node.next.Load(); m.head != nil {
			*node = queueNode{}
			nodes.Put(node)
			continue
		}

		if m.tail.CompareAndSwap(node, nil) {
			*node = queueNode{}
			nodes.Put(node)
			return
		}

		for m.head == nil {
			runtime.Gosched()
			m.head = node.next.Load()
		}
		*node = queueNode{}
		nodes.Put(node)
	}
}
