package kdsync

import (
	"github.com/iakud/knoll/kdsync/wire"
)

type Message[T any] interface {
	*T
	wire.Marshaler
	wire.Unmarshaler
	ClearDirty()
	ClearPersistDirty()
}

type DirtyFunc func()

func (f DirtyFunc) Invoke() {
	if f == nil {
		return
	}
	f()
}

func NoSync()    {}
func NoPersist() {}

type MessageType[T any, M Message[T]] struct {
	New              func() M
	CheckDirtyParent func(M) bool
	SetDirtyParent   func(M, DirtyFunc, DirtyFunc)
	ClearDirtyParent func(M)
}
