package kdsync

import (
	"github.com/iakud/knoll/kdsync/wire"
)

type MessageState struct {
	setParent    func(f DirtyFunc)
	removeParent func()
}

type Message[T any] interface {
	*T
	wire.Marshaler
	wire.Unmarshaler
	// MessageState() *MessageState
	MarshalDirty(b []byte) ([]byte, error)
	SetDirtyParent(f DirtyFunc)
	GetDirtyParent() DirtyFunc
	MarkDirtyAll()
	ClearDirty()
}

type DirtyFunc func()

func (f DirtyFunc) Invoke() {
	if f == nil {
		return
	}
	f()
}
