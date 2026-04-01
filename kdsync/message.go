package kdsync

import (
	"github.com/iakud/knoll/kdsync/wire"
)

type Message[T any] interface {
	*T
	wire.Marshaler
	wire.Unmarshaler
	MarkDirty()
	ClearDirty()
}

type DirtyFunc func()

func (f DirtyFunc) Invoke() {
	if f == nil {
		return
	}
	f()
}

type MessageType[T any, M Message[T]] struct {
	New              func() M
	CheckDirtyParent func(M) bool
	SetDirtyParent   func(M, DirtyFunc)
	ClearDirtyParent func(M)
}
