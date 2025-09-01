package krpc

import "sync"

var cPool = sync.Pool{New: func() any { return &userMsg{} }}

func NewCMessage() Msg {
	return cPool.Get().(Msg)
}
