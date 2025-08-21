package backend

import (
	"sync"

	"github.com/iakud/knoll/krpc"
)

var bPool = sync.Pool{New: func() any { return &Message{} }}

func NewBMessage() krpc.Message {
	return bPool.Get().(krpc.Message)
}
