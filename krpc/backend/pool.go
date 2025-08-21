package backend

import (
	"sync"

	"github.com/iakud/knoll/krpc"
)

var pool = sync.Pool{New: func() any { return &Message{} }}

func New() krpc.Message {
	return pool.Get().(krpc.Message)
}
