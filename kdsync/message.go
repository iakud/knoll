package kdsync

import (
	"time"

	"github.com/iakud/knoll/kdsync/wire"
)

type Field interface {
	bool | int32 | uint32 | int64 | uint64 | float32 | float64 | string | time.Duration | struct{}
}

type MessageState struct {
	setParent    func(f DirtyFunc)
	removeParent func()
}

type Message[T any] interface {
	*T
	wire.Marshaler
	wire.Unmarshaler
	MessageState() *MessageState
	MarshalDirty(b []byte) ([]byte, error)
	SetDirtyParent(f DirtyFunc)
	GetDirtyParent() DirtyFunc
	MarkDirtyAll()
	ClearDirth()

	String(indent string) string
}

type DirtyFunc func()

func (f DirtyFunc) Invoke() {
	if f == nil {
		return
	}
	f()
}
