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
	data  map[K]V
	clear bool

	updates map[K]V
	deletes map[K]struct{}

	dirty       bool
	dirtyParent DirtyFunc
	keyCodec    FieldCodec[K]
	valueCodec  FieldCodec[V]
}

func (x *MapField[K, V]) Init(dirtyParent DirtyFunc, keyCodec FieldCodec[K], valueCodec FieldCodec[V]) {
	x.data = make(map[K]V)
	x.updates = make(map[K]V)
	x.deletes = make(map[K]struct{})
	x.dirtyParent = dirtyParent
	x.keyCodec = keyCodec
	x.valueCodec = valueCodec
}

func (x *MapField[K, V]) Len() int {
	return len(x.data)
}

func (x *MapField[K, V]) Clear() {
	if len(x.data) == 0 && len(x.deletes) == 0 {
		return
	}
	clear(x.data)
	x.clear = true
	clear(x.updates)
	clear(x.deletes)
	x.markDirty()
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
	x.updates[k] = v
	delete(x.deletes, k)
	x.markDirty()
}

func (x *MapField[K, V]) Delete(k K) {
	if _, ok := x.data[k]; !ok {
		return
	}
	delete(x.data, k)
	delete(x.updates, k)
	x.deletes[k] = struct{}{}
	x.markDirty()
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

func (x *MapField[K, V]) markDirty() {
	if x.dirty {
		return
	}
	x.dirty = true
	x.dirtyParent.Invoke()
}

func (x *MapField[K, V]) ClearDirty() {
	if !x.dirty {
		return
	}
	x.clear = false
	clear(x.updates)
	clear(x.deletes)
	x.dirty = false
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
	if x.clear {
		if b, err = wire.MarshalBool(b, wire.MapClearFieldNumber, true); err != nil {
			return b, err
		}
	}
	if len(x.deletes) > 0 {
		b = wire.AppendTag(b, wire.MapDeleteFieldNumber, wire.BytesType)
		b, pos = wire.AppendSpeculativeLength(b)
		for k := range x.deletes {
			b = x.keyCodec.marshalFunc(b, k)
		}
		b = wire.FinishSpeculativeLength(b, pos)
	}
	for k, v := range x.updates {
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
	data  map[K]V
	clear bool

	updates map[K]V
	deletes map[K]struct{}

	dirty       bool
	dirtyParent DirtyFunc
	keyCodec    FieldCodec[K]
}

func (x *MapMessage[K, T, V]) Init(dirtyParent DirtyFunc, keyCodec FieldCodec[K]) {
	x.data = make(map[K]V)
	x.updates = make(map[K]V)
	x.deletes = make(map[K]struct{})
	x.dirtyParent = dirtyParent
	x.keyCodec = keyCodec
}

func (x *MapMessage[K, T, V]) Len() int {
	return len(x.data)
}

func (x *MapMessage[K, T, V]) Clear() {
	if len(x.data) == 0 && len(x.deletes) == 0 {
		return
	}
	for _, v := range x.data {
		if v != nil {
			v.SetDirtyParent(nil)
		}
	}
	clear(x.data)
	x.clear = true
	clear(x.updates)
	clear(x.deletes)
	x.markDirty()
}

func (x *MapMessage[K, T, V]) Get(k K) (V, bool) {
	v, ok := x.data[k]
	return v, ok
}

func (x *MapMessage[K, T, V]) Set(k K, v V) {
	if v != nil && v.GetDirtyParent() != nil {
		panic("the component should be removed from its original place first")
	}
	if e, ok := x.data[k]; ok {
		if e == v {
			return
		}
		if e != nil {
			e.SetDirtyParent(nil)
		}
	}
	if v != nil {
		v.SetDirtyParent(func() {
			if _, ok := x.updates[k]; ok {
				return
			}
			x.updates[k] = v
			x.markDirty()
		})
		v.MarkDirtyAll()
	}
	x.data[k] = v
	x.updates[k] = v
	delete(x.deletes, k)
	x.markDirty()
}

func (x *MapMessage[K, T, V]) Delete(k K) {
	if v, ok := x.data[k]; !ok {
		return
	} else if v != nil {
		v.SetDirtyParent(nil)
	}
	delete(x.data, k)
	delete(x.updates, k)
	x.deletes[k] = struct{}{}
	x.markDirty()
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

func (x *MapMessage[K, T, V]) markDirty() {
	if x.dirty {
		return
	}
	x.dirty = true
	x.dirtyParent.Invoke()
}

func (x *MapMessage[K, T, V]) ClearDirty() {
	if !x.dirty {
		return
	}
	for _, v := range x.updates {
		if v != nil {
			v.ClearDirty()
		}
	}
	x.clear = false
	clear(x.updates)
	clear(x.deletes)
	x.dirty = false
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
	if x.clear {
		if b, err = wire.MarshalBool(b, wire.MapClearFieldNumber, true); err != nil {
			return b, err
		}
	}
	if len(x.deletes) > 0 {
		b = wire.AppendTag(b, wire.MapDeleteFieldNumber, wire.BytesType)
		b, pos = wire.AppendSpeculativeLength(b)
		for k := range x.deletes {
			b = x.keyCodec.marshalFunc(b, k)
		}
		b = wire.FinishSpeculativeLength(b, pos)
	}
	for k, v := range x.updates {
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
			c = new(T)
			if err := c.Unmarshal(v); err != nil {
				return err
			}
			x.Set(k, c)
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
