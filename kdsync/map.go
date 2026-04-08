package kdsync

import (
	"fmt"
	"iter"
	"maps"
	"slices"
	"time"

	"github.com/iakud/knoll/kdsync/wire"
)

type Map[K comparable, V any] interface {
	Len() int
	Clear()
	Get(k K) (V, bool)
	Set(k K, v V)
	Delete(k K)
	All() iter.Seq2[K, V]
	Keys() iter.Seq[K]
	Values() iter.Seq[V]

	ClearDirty()
	ClearPersistDirty()
	Marshal(b []byte) ([]byte, error)
	MarshalDirty(b []byte) ([]byte, error)
	Unmarshal(b []byte) error
	MarshalJSONIndent(b []byte, prefix string, indent string) ([]byte, error)
}

// FieldCodec map check
var _ Map[bool, bool] = (*MapField[bool, bool])(nil)
var _ Map[int32, int32] = (*MapField[int32, int32])(nil)
var _ Map[uint32, uint32] = (*MapField[uint32, uint32])(nil)
var _ Map[int64, int64] = (*MapField[int64, int64])(nil)
var _ Map[uint64, uint64] = (*MapField[uint64, uint64])(nil)
var _ Map[float32, float32] = (*MapField[float32, float32])(nil)
var _ Map[float64, float64] = (*MapField[float64, float64])(nil)
var _ Map[string, string] = (*MapField[string, string])(nil)
var _ Map[string, []byte] = (*MapField[string, []byte])(nil)
var _ Map[string, time.Time] = (*MapField[string, time.Time])(nil)
var _ Map[string, time.Duration] = (*MapField[string, time.Duration])(nil)

// Field map
type MapField[K comparable, V any] struct {
	data map[K]V

	syncCleared bool
	syncUpdated map[K]V
	syncDeleted map[K]struct{}

	persistCleared bool
	persistUpdated map[K]V
	persistDeleted map[K]struct{}

	dirtyParent DirtyFunc

	keyCodec   FieldCodec[K]
	valueCodec FieldCodec[V]
}

func (x *MapField[K, V]) Init(dirtyParent DirtyFunc, keyCodec FieldCodec[K], valueCodec FieldCodec[V]) {
	x.data = make(map[K]V)
	x.syncUpdated = make(map[K]V)
	x.syncDeleted = make(map[K]struct{})
	x.persistUpdated = make(map[K]V)
	x.persistDeleted = make(map[K]struct{})
	x.dirtyParent = dirtyParent
	x.keyCodec = keyCodec
	x.valueCodec = valueCodec
}

func (x *MapField[K, V]) Len() int {
	return len(x.data)
}

func (x *MapField[K, V]) Clear() {
	if len(x.data) == 0 && len(x.syncDeleted) == 0 {
		return
	}
	clear(x.data)
	x.updateDirtyCleared()
}

func (x *MapField[K, V]) Get(k K) (V, bool) {
	v, ok := x.data[k]
	return v, ok
}

func (x *MapField[K, V]) Set(k K, v V) {
	if e, ok := x.data[k]; ok {
		if x.valueCodec.compareFunc(v, e) == 0 {
			return
		}
	}
	x.data[k] = v
	x.updateDirtyUpdated(k, v)
}

func (x *MapField[K, V]) Delete(k K) {
	if _, ok := x.data[k]; !ok {
		return
	}
	delete(x.data, k)
	x.updateDirtyDeleted(k)
}

func (x *MapField[K, V]) All() iter.Seq2[K, V] {
	return maps.All(x.data)
}

func (x *MapField[K, V]) Keys() iter.Seq[K] {
	return maps.Keys(x.data)
}

func (x *MapField[K, V]) Values() iter.Seq[V] {
	return maps.Values(x.data)
}

func (x *MapField[K, V]) updateDirtyCleared() {
	if len(x.syncUpdated) == 0 && len(x.syncDeleted) == 0 && len(x.persistUpdated) == 0 && len(x.persistDeleted) == 0 {
		return
	}

	x.syncCleared = true
	clear(x.syncUpdated)
	clear(x.syncDeleted)
	x.persistCleared = true
	clear(x.persistUpdated)
	clear(x.persistDeleted)
	x.dirtyParent.Invoke(DirtyType_SyncAndPersist)
}

func (x *MapField[K, V]) updateDirtyUpdated(k K, v V) {
	_, syncOk := x.syncUpdated[k]
	_, persistOk := x.persistUpdated[k]
	if syncOk && persistOk {
		return
	}
	x.syncUpdated[k] = v
	delete(x.syncDeleted, k)
	x.persistUpdated[k] = v
	delete(x.persistDeleted, k)
	x.dirtyParent.Invoke(DirtyType_SyncAndPersist)
}

func (x *MapField[K, V]) updateDirtyDeleted(k K) {
	_, syncOk := x.syncUpdated[k]
	_, persistOk := x.persistDeleted[k]
	if syncOk && persistOk {
		return
	}
	delete(x.syncUpdated, k)
	x.syncDeleted[k] = struct{}{}
	delete(x.persistUpdated, k)
	x.persistDeleted[k] = struct{}{}
	x.dirtyParent.Invoke(DirtyType_SyncAndPersist)
}

func (x *MapField[K, V]) ClearDirty() {
	if !x.syncCleared && len(x.syncUpdated) == 0 && len(x.syncDeleted) == 0 {
		return
	}
	x.syncCleared = false
	clear(x.syncUpdated)
	clear(x.syncDeleted)
}

func (x *MapField[K, V]) ClearPersistDirty() {
	if !x.persistCleared && len(x.persistUpdated) == 0 && len(x.persistDeleted) == 0 {
		return
	}
	x.persistCleared = false
	clear(x.persistUpdated)
	clear(x.persistDeleted)
}

func (x *MapField[K, V]) Marshal(b []byte) ([]byte, error) {
	var pos int
	var err error
	if b, err = wire.MarshalBool(b, wire.MapClearFieldNumber, true); err != nil {
		return b, err
	}
	for k, v := range x.data {
		b = wire.AppendTag(b, wire.MapEntryFieldNumber, wire.BytesType)
		b, pos = wire.AppendSpeculativeLength(b)
		b = wire.AppendTag(b, wire.MapEntryKeyFieldNumber, x.keyCodec.wireType)
		b = x.keyCodec.marshalFunc(b, k)
		b = wire.AppendTag(b, wire.MapEntryValueFieldNumber, x.valueCodec.wireType)
		b = x.valueCodec.marshalFunc(b, v)
		b = wire.FinishSpeculativeLength(b, pos)
	}
	return b, err
}

func (x *MapField[K, V]) MarshalDirty(b []byte) ([]byte, error) {
	var pos int
	var err error
	if x.syncCleared {
		if b, err = wire.MarshalBool(b, wire.MapClearFieldNumber, true); err != nil {
			return b, err
		}
	}
	if len(x.syncDeleted) > 0 {
		b = wire.AppendTag(b, wire.MapDeleteFieldNumber, wire.BytesType)
		b, pos = wire.AppendSpeculativeLength(b)
		for k := range x.syncDeleted {
			b = x.keyCodec.marshalFunc(b, k)
		}
		b = wire.FinishSpeculativeLength(b, pos)
	}
	for k, v := range x.syncUpdated {
		b = wire.AppendTag(b, wire.MapEntryFieldNumber, wire.BytesType)
		b, pos = wire.AppendSpeculativeLength(b)
		b = wire.AppendTag(b, wire.MapEntryKeyFieldNumber, x.keyCodec.wireType)
		b = x.keyCodec.marshalFunc(b, k)
		b = wire.AppendTag(b, wire.MapEntryValueFieldNumber, x.valueCodec.wireType)
		b = x.valueCodec.marshalFunc(b, v)
		b = wire.FinishSpeculativeLength(b, pos)
	}
	return b, err
}

func (x *MapField[K, V]) Unmarshal(b []byte) error {
	var clear bool
	var deletes []byte
	var entries [][]byte
	for len(b) > 0 {
		num, wtyp, tagLen, err := wire.ConsumeTag(b)
		if err != nil {
			return err
		}
		var valLen int
		err = wire.ErrUnknown
		switch num {
		case wire.MapClearFieldNumber:
			clear, valLen, err = wire.UnmarshalBool(b[tagLen:], wtyp)
		case wire.MapDeleteFieldNumber:
			deletes, valLen, err = wire.UnmarshalBytes(b[tagLen:], wtyp)
		case wire.MapEntryFieldNumber:
			var entry []byte
			if entry, valLen, err = wire.UnmarshalBytes(b[tagLen:], wtyp); err != nil {
				break
			}
			entries = append(entries, entry)
		}
		if err == wire.ErrUnknown {
			if valLen, err = wire.ConsumeFieldValue(num, wtyp, b[tagLen:]); err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
		b = b[tagLen+valLen:]
	}
	if clear {
		x.Clear()
	}
	for b := deletes; len(b) > 0; {
		k, n, err := x.keyCodec.unmarshalFunc(b)
		if err != nil {
			return err
		}
		b = b[n:]
		x.Delete(k)
	}
	for _, b := range entries {
		var k K
		var v V
		for len(b) > 0 {
			num, wtyp, tagLen, err := wire.ConsumeTag(b)
			if err != nil {
				return err
			}
			var valLen int
			err = wire.ErrUnknown
			switch num {
			case wire.MapEntryKeyFieldNumber:
				if wtyp != x.keyCodec.wireType {
					break
				}
				k, valLen, err = x.keyCodec.unmarshalFunc(b[tagLen:])
			case wire.MapEntryValueFieldNumber:
				if wtyp != x.valueCodec.wireType {
					break
				}
				v, valLen, err = x.valueCodec.unmarshalFunc(b[tagLen:])
			}
			if err == wire.ErrUnknown {
				if valLen, err = wire.ConsumeFieldValue(num, wtyp, b[tagLen:]); err != nil {
					return err
				}
			} else if err != nil {
				return err
			}
			b = b[tagLen+valLen:]
		}
		x.Set(k, v)
	}
	return nil
}

func (x *MapField[K, V]) MarshalJSONIndent(b []byte, prefix string, indent string) ([]byte, error) {
	if len(x.data) == 0 {
		return append(b, "{}"...), nil
	}

	keys := slices.SortedFunc(maps.Keys(x.data), x.keyCodec.compareFunc)
	var err error
	b = append(b, "{\n"...)
	for i, k := range keys {
		b = append(b, prefix+indent...)
		b = append(b, '"')
		b = fmt.Append(b, k)
		b = append(b, '"')
		b = append(b, ": "...)
		b, err = MarshalJSONIndent(b, x.data[k], prefix+indent, indent)
		if err != nil {
			return nil, err
		}
		if i+1 < len(keys) {
			b = append(b, ',')
		}
		b = append(b, '\n')
	}
	b = append(b, prefix...)
	b = append(b, '}')
	return b, nil
}

// Message map
type MapMessage[K comparable, T any, V Message[T]] struct {
	data map[K]V

	syncCleared bool
	syncUpdated map[K]V
	syncDeleted map[K]struct{}

	persistCleared bool
	persistUpdated map[K]V
	persistDeleted map[K]struct{}

	dirtyParent DirtyFunc

	keyCodec  FieldCodec[K]
	valueType *MessageType[T, V]
}

func (x *MapMessage[K, T, V]) Init(dirtyParent DirtyFunc, keyCodec FieldCodec[K], valueType *MessageType[T, V]) {
	x.data = make(map[K]V)
	x.syncUpdated = make(map[K]V)
	x.syncDeleted = make(map[K]struct{})
	x.persistUpdated = make(map[K]V)
	x.persistDeleted = make(map[K]struct{})
	x.dirtyParent = dirtyParent
	x.keyCodec = keyCodec
	x.valueType = valueType
}

func (x *MapMessage[K, T, V]) Len() int {
	return len(x.data)
}

func (x *MapMessage[K, T, V]) Clear() {
	if len(x.data) == 0 && len(x.syncDeleted) == 0 {
		return
	}
	for _, v := range x.data {
		if v != nil {
			x.valueType.ClearDirtyParent(v)
		}
	}
	clear(x.data)
	x.updateDirtyCleared()
}

func (x *MapMessage[K, T, V]) Get(k K) (V, bool) {
	v, ok := x.data[k]
	return v, ok
}

func (x *MapMessage[K, T, V]) Set(k K, v V) {
	if v != nil {
		if x.valueType.CheckDirtyParent(v) {
			panic("the component should be removed from its original place first")
		}
	}
	if e, ok := x.data[k]; ok {
		if e == v {
			return
		}
		if e != nil {
			x.valueType.ClearDirtyParent(e)
		}
	}
	if v != nil {
		x.valueType.SetDirtyParent(v, func(dirtyType DirtyType) {
			x.updateDirtyUpdated(k, v, dirtyType)
		})
	}
	x.data[k] = v
	x.updateSyncAndPersistUpdated(k, v)
}

func (x *MapMessage[K, T, V]) Delete(k K) {
	v, ok := x.data[k]
	if !ok {
		return
	}
	if v != nil {
		x.valueType.ClearDirtyParent(v)
	}
	delete(x.data, k)
	x.updateDirtyDeleted(k)
}

func (x *MapMessage[K, T, V]) All() iter.Seq2[K, V] {
	return maps.All(x.data)
}

func (x *MapMessage[K, T, V]) Keys() iter.Seq[K] {
	return maps.Keys(x.data)
}

func (x *MapMessage[K, T, V]) Values() iter.Seq[V] {
	return maps.Values(x.data)
}

func (x *MapMessage[K, T, V]) updateDirtyCleared() {
	if len(x.syncUpdated) == 0 && len(x.syncDeleted) == 0 && len(x.persistUpdated) == 0 && len(x.persistDeleted) == 0 {
		return
	}

	x.syncCleared = true
	clear(x.syncUpdated)
	clear(x.syncDeleted)
	x.persistCleared = true
	clear(x.persistUpdated)
	clear(x.persistDeleted)
	x.dirtyParent.Invoke(DirtyType_SyncAndPersist)
}

func (x *MapMessage[K, T, V]) updateDirtyUpdated(k K, v V, dirtyType DirtyType) {
	switch dirtyType {
	case DirtyType_Sync:
		x.updateSyncUpdated(k, v)
	case DirtyType_Persist:
		x.updatePersistUpdated(k, v)
	case DirtyType_SyncAndPersist:
		x.updateSyncAndPersistUpdated(k, v)
	default:
		// nothing to do
	}
}

func (x *MapMessage[K, T, V]) updateSyncUpdated(k K, v V) {

	if _, syncOk := x.syncUpdated[k]; syncOk {
		return
	}
	x.syncUpdated[k] = v
	delete(x.syncDeleted, k)
	x.dirtyParent.Invoke(DirtyType_Sync)
}

func (x *MapMessage[K, T, V]) updatePersistUpdated(k K, v V) {
	if _, persistOk := x.persistUpdated[k]; persistOk {
		return
	}
	x.persistUpdated[k] = v
	delete(x.persistDeleted, k)
	x.dirtyParent.Invoke(DirtyType_Persist)
}

func (x *MapMessage[K, T, V]) updateSyncAndPersistUpdated(k K, v V) {
	_, syncOk := x.syncUpdated[k]
	_, persistOk := x.persistUpdated[k]
	if syncOk && persistOk {
		return
	}
	x.syncUpdated[k] = v
	delete(x.syncDeleted, k)
	x.persistUpdated[k] = v
	delete(x.persistDeleted, k)
	x.dirtyParent.Invoke(DirtyType_SyncAndPersist)
}

func (x *MapMessage[K, T, V]) updateDirtyDeleted(k K) {
	_, syncOk := x.syncUpdated[k]
	_, persistOk := x.persistDeleted[k]
	if syncOk && persistOk {
		return
	}
	delete(x.syncUpdated, k)
	x.syncDeleted[k] = struct{}{}
	delete(x.persistUpdated, k)
	x.persistDeleted[k] = struct{}{}
	x.dirtyParent.Invoke(DirtyType_SyncAndPersist)
}

func (x *MapMessage[K, T, V]) ClearDirty() {
	if !x.syncCleared && len(x.syncUpdated) == 0 && len(x.syncDeleted) == 0 {
		return
	}
	for _, v := range x.syncUpdated {
		if v != nil {
			v.ClearDirty()
		}
	}
	x.syncCleared = false
	clear(x.syncUpdated)
	clear(x.syncDeleted)
}

func (x *MapMessage[K, T, V]) ClearPersistDirty() {
	if !x.persistCleared && len(x.persistUpdated) == 0 && len(x.persistDeleted) == 0 {
		return
	}
	for _, v := range x.persistUpdated {
		if v != nil {
			v.ClearPersistDirty()
		}
	}
	x.persistCleared = false
	clear(x.persistUpdated)
	clear(x.persistDeleted)
}

func (x *MapMessage[K, T, V]) Marshal(b []byte) ([]byte, error) {
	var pos int
	var err error
	if b, err = wire.MarshalBool(b, wire.MapClearFieldNumber, true); err != nil {
		return b, err
	}
	for k, v := range x.data {
		b = wire.AppendTag(b, wire.MapEntryFieldNumber, wire.BytesType)
		b, pos = wire.AppendSpeculativeLength(b)
		b = wire.AppendTag(b, wire.MapEntryKeyFieldNumber, x.keyCodec.wireType)
		b = x.keyCodec.marshalFunc(b, k)
		if b, err = wire.MarshalMessage(b, wire.MapEntryValueFieldNumber, v); err != nil {
			return b, err
		}
		b = wire.FinishSpeculativeLength(b, pos)
	}
	return b, err
}

func (x *MapMessage[K, T, V]) MarshalDirty(b []byte) ([]byte, error) {
	var pos int
	var err error
	if x.syncCleared {
		if b, err = wire.MarshalBool(b, wire.MapClearFieldNumber, true); err != nil {
			return b, err
		}
	}
	if len(x.syncDeleted) > 0 {
		b = wire.AppendTag(b, wire.MapDeleteFieldNumber, wire.BytesType)
		b, pos = wire.AppendSpeculativeLength(b)
		for k := range x.syncDeleted {
			b = x.keyCodec.marshalFunc(b, k)
		}
		b = wire.FinishSpeculativeLength(b, pos)
	}
	for k, v := range x.syncUpdated {
		b = wire.AppendTag(b, wire.MapEntryFieldNumber, wire.BytesType)
		b, pos = wire.AppendSpeculativeLength(b)
		b = wire.AppendTag(b, wire.MapEntryKeyFieldNumber, x.keyCodec.wireType)
		b = x.keyCodec.marshalFunc(b, k)
		if b, err = wire.MarshalMessageDirty(b, wire.MapEntryValueFieldNumber, v); err != nil {
			return b, err
		}
		b = wire.FinishSpeculativeLength(b, pos)
	}
	return b, err
}

func (x *MapMessage[K, T, V]) Unmarshal(b []byte) error {
	var clear bool
	var deletes []byte
	var entries [][]byte
	for len(b) > 0 {
		num, wtyp, tagLen, err := wire.ConsumeTag(b)
		if err != nil {
			return err
		}
		var valLen int
		err = wire.ErrUnknown
		switch num {
		case wire.MapClearFieldNumber:
			clear, valLen, err = wire.UnmarshalBool(b[tagLen:], wtyp)
		case wire.MapDeleteFieldNumber:
			deletes, valLen, err = wire.UnmarshalBytes(b[tagLen:], wtyp)
		case wire.MapEntryFieldNumber:
			var entry []byte
			if entry, valLen, err = wire.UnmarshalBytes(b[tagLen:], wtyp); err != nil {
				break
			}
			entries = append(entries, entry)
		}
		if err == wire.ErrUnknown {
			if valLen, err = wire.ConsumeFieldValue(num, wtyp, b[tagLen:]); err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
		b = b[tagLen+valLen:]
	}
	if clear {
		x.Clear()
	}
	for b := deletes; len(b) > 0; {
		k, n, err := x.keyCodec.unmarshalFunc(b)
		if err != nil {
			return err
		}
		b = b[n:]
		x.Delete(k)
	}
	for _, b := range entries {
		var k K
		var v []byte
		for len(b) > 0 {
			num, wtyp, tagLen, err := wire.ConsumeTag(b)
			if err != nil {
				return err
			}
			var valLen int
			err = wire.ErrUnknown
			switch num {
			case wire.MapEntryKeyFieldNumber:
				if wtyp != x.keyCodec.wireType {
					break
				}
				k, valLen, err = x.keyCodec.unmarshalFunc(b)
			case wire.MapEntryValueFieldNumber:
				v, valLen, err = wire.UnmarshalBytes(b[tagLen:], wtyp)
			}
			if err == wire.ErrUnknown {
				if valLen, err = wire.ConsumeFieldValue(num, wtyp, b[tagLen:]); err != nil {
					return err
				}
			} else if err != nil {
				return err
			}
			b = b[tagLen+valLen:]
		}

		if c, ok := x.data[k]; !ok {
			c = x.valueType.New()
			if err := c.Unmarshal(v); err != nil {
				return err
			}
			x.data[k] = c
			x.valueType.SetDirtyParent(c, func(dirtyType DirtyType) {
				x.updateDirtyUpdated(k, c, dirtyType)
			})
		} else if err := c.Unmarshal(v); err != nil {
			return err
		}
	}
	return nil
}

func (x *MapMessage[K, T, V]) MarshalJSONIndent(b []byte, prefix string, indent string) ([]byte, error) {
	if len(x.data) == 0 {
		return append(b, "{}"...), nil
	}

	keys := slices.SortedFunc(maps.Keys(x.data), x.keyCodec.compareFunc)
	var err error
	b = append(b, "{\n"...)
	for i, k := range keys {
		b = append(b, prefix+indent...)
		b = append(b, '"')
		b = fmt.Append(b, k)
		b = append(b, '"')
		b = append(b, ": "...)
		b, err = MarshalJSONIndent(b, x.data[k], prefix+indent, indent)
		if err != nil {
			return nil, err
		}
		if i+1 < len(keys) {
			b = append(b, ',')
		}
		b = append(b, '\n')
	}
	b = append(b, prefix...)
	b = append(b, '}')
	return b, nil
}
