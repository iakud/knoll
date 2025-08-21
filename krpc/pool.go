package krpc

import "sync"

var cPool = sync.Pool{New: func() any { return &CMessage{} }}

func NewCMessage() Message {
	return cPool.Get().(Message)
}
