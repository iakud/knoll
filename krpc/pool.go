package krpc

import "sync"

var cPool = sync.Pool{New: func() any { return &CMessage{} }}

var bPool = sync.Pool{New: func() any { return &BMessage{} }}

func NewCMessage() Message {
	return cPool.Get().(Message)
}

func NewBMessage() Message {
	return bPool.Get().(Message)
}
