package kdsync

import (
	"github.com/iakud/knoll/kdsync/wire"
)

type MessageState struct {
	dirtyParentFunc *DirtyFunc
}

func (ms *MessageState) Init(dirtyParentFunc *DirtyFunc) {
	ms.dirtyParentFunc = dirtyParentFunc
}

func (ms *MessageState) checkDirtyParentFunc() bool {
	return *ms.dirtyParentFunc == nil
}

func (ms *MessageState) setDirtyParentFunc(f DirtyFunc) {
	*ms.dirtyParentFunc = f
}

type Message[T any] interface {
	*T
	wire.Marshaler
	wire.Unmarshaler
	MessageState() *MessageState
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
