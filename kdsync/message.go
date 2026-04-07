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

type DirtyType byte

const (
	DirtyType_None           DirtyType = 0x00
	DirtyType_Sync           DirtyType = 0x01
	DirtyType_Persist        DirtyType = 0x02
	DirtyType_SyncAndPersist DirtyType = DirtyType_Sync | DirtyType_Persist
)

type DirtyFunc func(DirtyType)

func (f DirtyFunc) Invoke(dirtyType DirtyType) {
	if f == nil {
		return
	}
	f(dirtyType)
}

func NoSync()    {}
func NoPersist() {}

type MessageType[T any, M Message[T]] struct {
	New              func() M
	CheckDirtyParent func(M) bool
	SetDirtyParent   func(M, DirtyFunc)
	ClearDirtyParent func(M)
}
