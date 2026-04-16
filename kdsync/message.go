package kdsync

import (
	"github.com/iakud/knoll/kdsync/kdsjson"
	"github.com/iakud/knoll/kdsync/wire"
)

type Message[T any] interface {
	*T
	wire.Marshaler
	wire.Unmarshaler
	ClearDirty()
	ClearPersistDirty()
	WriteJSON(*kdsjson.Encoder) error
}

type DirtyType byte

const (
	DirtyType_None           DirtyType = 0x00
	DirtyType_Sync           DirtyType = 0x01
	DirtyType_Persist        DirtyType = 0x02
	DirtyType_SyncAndPersist DirtyType = DirtyType_Sync | DirtyType_Persist
)

type DirtyFunc func(DirtyType)

func (f DirtyFunc) Invoke(t DirtyType) {
	if f == nil {
		return
	}
	f(t)
}

type MessageType[T any, M Message[T]] struct {
	New              func() M
	CheckDirtyParent func(M) bool
	SetDirtyParent   func(M, DirtyFunc)
	ClearDirtyParent func(M)
}
