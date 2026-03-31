package kdsync

import (
	"iter"
	"slices"
	"time"

	"github.com/iakud/knoll/kdsync/wire"
)

type Repeated[T any] interface {
	Len() int
	Clear()
	Get(i int) T
	Set(i int, v T)
	Append(v ...T)
	Index(v T) int
	IndexFunc(f func(T) bool) int
	Contains(v T) bool
	ContainsFunc(f func(T) bool) bool
	Insert(i int, v ...T)
	Delete(i, j int)
	DeleteFunc(del func(T) bool)
	Replace(i, j int, v ...T)
	Reverse()
	All() iter.Seq2[int, T]
	Backward() iter.Seq2[int, T]
	Values() iter.Seq[T]

	ClearDirty()
	Marshal(b []byte) ([]byte, error)
	MarshalDirty(b []byte) ([]byte, error)
	Unmarshal(b []byte) error
	MarshalJSONIndent(b []byte, prefix string, indent string) ([]byte, error)
}

// Field repeated check
var _ Repeated[bool] = (*RepeatedField[bool])(nil)
var _ Repeated[int32] = (*RepeatedField[int32])(nil)
var _ Repeated[uint32] = (*RepeatedField[uint32])(nil)
var _ Repeated[int64] = (*RepeatedField[int64])(nil)
var _ Repeated[uint64] = (*RepeatedField[uint64])(nil)
var _ Repeated[float32] = (*RepeatedField[float32])(nil)
var _ Repeated[float64] = (*RepeatedField[float64])(nil)
var _ Repeated[string] = (*RepeatedField[string])(nil)
var _ Repeated[[]byte] = (*RepeatedField[[]byte])(nil)
var _ Repeated[time.Time] = (*RepeatedField[time.Time])(nil)
var _ Repeated[time.Duration] = (*RepeatedField[time.Duration])(nil)
var _ Repeated[struct{}] = (*RepeatedField[struct{}])(nil)

// Field repeated
type RepeatedField[E any] struct {
	data []E

	dirty       bool
	dirtyParent DirtyFunc
	fieldCodec  FieldCodec[E]
}

func (x *RepeatedField[E]) Init(dirtyParent DirtyFunc, fieldCodec FieldCodec[E]) {
	x.dirtyParent = dirtyParent
	x.fieldCodec = fieldCodec
}

func (x *RepeatedField[E]) Len() int {
	return len(x.data)
}

func (x *RepeatedField[E]) Clear() {
	if len(x.data) == 0 {
		return
	}
	clear(x.data)
	x.data = x.data[:0]
	x.markDirty()
}

func (x *RepeatedField[E]) Get(i int) E {
	return x.data[i]
}

func (x *RepeatedField[E]) Set(i int, v E) {
	if x.fieldCodec.compareFunc(v, x.data[i]) == 0 {
		return
	}
	x.data[i] = v
	x.markDirty()
}

func (x *RepeatedField[E]) Append(v ...E) {
	if len(v) == 0 {
		return
	}
	x.data = append(x.data, v...)
	x.markDirty()
}

func (x *RepeatedField[E]) Index(v E) int {
	for i := range x.data {
		if x.fieldCodec.compareFunc(v, x.data[i]) == 0 {
			return i
		}
	}
	return -1
}

func (x *RepeatedField[E]) IndexFunc(f func(E) bool) int {
	for i := range x.data {
		if f(x.data[i]) {
			return i
		}
	}
	return -1
}

func (x *RepeatedField[E]) Contains(v E) bool {
	return x.Index(v) >= 0
}

func (x *RepeatedField[E]) ContainsFunc(f func(E) bool) bool {
	return x.IndexFunc(f) >= 0
}

func (x *RepeatedField[E]) Insert(i int, v ...E) {
	_ = x.data[i:] // bounds check
	if len(v) == 0 {
		return
	}
	x.data = slices.Insert(x.data, i, v...)
	x.markDirty()
}

func (x *RepeatedField[E]) Delete(i, j int) {
	_ = x.data[i:j:len(x.data)] // bounds check
	if i == j {
		return
	}
	x.data = slices.Delete(x.data, i, j)
	x.markDirty()
}

func (x *RepeatedField[E]) DeleteFunc(del func(E) bool) {
	i := x.IndexFunc(del)
	if i == -1 {
		return
	}
	for j := i + 1; j < len(x.data); j++ {
		v := x.data[j]
		if del(v) {
			continue
		}
		x.data[i] = v
		i++
	}
	clear(x.data[i:])
	x.data = x.data[:i]
	x.markDirty()
}

func (x *RepeatedField[E]) Replace(i, j int, v ...E) {
	_ = x.data[i:j] // bounds check
	if i == j && len(v) == 0 {
		return
	}
	x.data = slices.Replace(x.data, i, j, v...)
	x.markDirty()
}

func (x *RepeatedField[E]) Reverse() {
	if len(x.data) < 2 {
		return
	}
	slices.Reverse(x.data)
	x.markDirty()
}

func (x *RepeatedField[E]) All() iter.Seq2[int, E] {
	return slices.All(x.data)
}

func (x *RepeatedField[E]) Backward() iter.Seq2[int, E] {
	return slices.Backward(x.data)
}

func (x *RepeatedField[E]) Values() iter.Seq[E] {
	return slices.Values(x.data)
}

func (x *RepeatedField[E]) markDirty() {
	if x.dirty {
		return
	}
	x.dirty = true
	x.dirtyParent.Invoke()
}

func (x *RepeatedField[E]) ClearDirty() {
	x.dirty = false
}

func (x *RepeatedField[E]) Marshal(b []byte) ([]byte, error) {
	if len(x.data) == 0 {
		return b, nil
	}
	for _, v := range x.data {
		b = x.fieldCodec.marshalFunc(b, v)
	}
	return b, nil
}

func (x *RepeatedField[E]) MarshalDirty(b []byte) ([]byte, error) {
	return x.Marshal(b)
}

func (x *RepeatedField[E]) Unmarshal(b []byte) error {
	x.Clear()
	for len(b) > 0 {
		v, n, err := x.fieldCodec.unmarshalFunc(b)
		if err != nil {
			return err
		}
		b = b[n:]
		x.data = append(x.data, v)
	}
	return nil
}

func (x *RepeatedField[E]) MarshalJSONIndent(b []byte, prefix string, indent string) ([]byte, error) {
	if len(x.data) == 0 {
		return append(b, "[]"...), nil
	}
	var err error
	b = append(b, "[\n"...)
	for i, v := range x.data {
		b = append(b, prefix+indent...)
		b, err = MarshalJSONIndent(b, v, prefix+indent, indent)
		if err != nil {
			return nil, err
		}
		if i+1 < len(x.data) {
			b = append(b, ',')
		}
		b = append(b, '\n')
	}
	b = append(b, prefix...)
	b = append(b, ']')
	return b, nil
}

// Message repeated
type RepeatedMessage[T any, E Message[T]] struct {
	data []E

	dirty       bool
	dirtyParent DirtyFunc
}

func (x *RepeatedMessage[T, E]) Init(dirtyParent DirtyFunc) {
	x.dirtyParent = dirtyParent
}

func (x *RepeatedMessage[T, E]) Len() int {
	return len(x.data)
}

func (x *RepeatedMessage[T, E]) Clear() {
	if len(x.data) == 0 {
		return
	}
	for _, v := range x.data {
		if v != nil {
			v.SetDirtyParent(nil)
		}
	}
	clear(x.data)
	x.data = x.data[:0]
	x.markDirty()
}

func (x *RepeatedMessage[T, E]) Get(i int) E {
	return x.data[i]
}

func (x *RepeatedMessage[T, E]) Set(i int, v E) {
	if v != nil && v.GetDirtyParent() != nil {
		panic("the component should be removed from its original place first")
	}
	if v == x.data[i] {
		return
	}
	if v != nil {
		v.SetDirtyParent(nil)
	}
	x.data[i] = v
	if v != nil {
		v.SetDirtyParent(x.markDirty)
		v.MarkDirtyAll()
	}
	x.markDirty()
}

func (x *RepeatedMessage[T, E]) Append(v ...E) {
	for i := range v {
		if v[i] != nil && v[i].GetDirtyParent() != nil {
			panic("the component should be removed from its original place first")
		}
	}
	if len(v) == 0 {
		return
	}
	x.data = append(x.data, v...)
	for i := range v {
		if v[i] != nil {
			v[i].SetDirtyParent(x.markDirty)
			v[i].MarkDirtyAll()
		}
	}
	x.markDirty()
}

func (x *RepeatedMessage[T, E]) Index(v E) int {
	for i := range x.data {
		if x.data[i] == v {
			return i
		}
	}
	return -1
}

func (x *RepeatedMessage[T, E]) IndexFunc(f func(E) bool) int {
	for i := range x.data {
		if f(x.data[i]) {
			return i
		}
	}
	return -1
}

func (x *RepeatedMessage[T, E]) Contains(v E) bool {
	return x.Index(v) >= 0
}

func (x *RepeatedMessage[T, E]) ContainsFunc(f func(E) bool) bool {
	return x.IndexFunc(f) >= 0
}

func (x *RepeatedMessage[T, E]) Insert(i int, v ...E) {
	_ = x.data[i:] // bounds check
	if len(v) == 0 {
		return
	}
	for j := range v {
		if v[j] != nil && v[j].GetDirtyParent() != nil {
			panic("the component should be removed from its original place first")
		}
	}
	x.data = slices.Insert(x.data, i, v...)
	for j := range v {
		if v[j] != nil {
			v[j].SetDirtyParent(x.markDirty)
			v[j].MarkDirtyAll()
		}
	}
	x.markDirty()
}

func (x *RepeatedMessage[T, E]) Delete(i, j int) {
	if i == j {
		return
	}
	r := x.data[i:j:len(x.data)]
	for k := range r {
		if r[k] != nil {
			r[k].SetDirtyParent(nil)
		}
	}
	x.data = slices.Delete(x.data, i, j)
	x.markDirty()
}

func (x *RepeatedMessage[T, E]) DeleteFunc(del func(E) bool) {
	i := x.IndexFunc(del)
	if i == -1 {
		return
	}
	if x.data[i] != nil {
		x.data[i].SetDirtyParent(nil)
	}
	for j := i + 1; j < len(x.data); j++ {
		v := x.data[j]
		if del(v) {
			if v != nil {
				v.SetDirtyParent(nil)
			}
			continue
		}
		x.data[i] = v
		i++
	}
	clear(x.data[i:])
	x.data = x.data[:i]
	x.markDirty()
}

func (x *RepeatedMessage[T, E]) Replace(i, j int, v ...E) {
	if i == j && len(v) == 0 {
		return
	}
	for k := range v {
		if v[k] != nil && v[k].GetDirtyParent() != nil {
			panic("the component should be removed from its original place first")
		}
	}
	r := x.data[i:j:len(x.data)]
	for k := range r {
		if r[k] != nil && r[k].GetDirtyParent() != nil {
			r[k].SetDirtyParent(nil)
		}
	}
	x.data = slices.Replace(x.data, i, j, v...)
	for k := range v {
		if v[k] != nil {
			v[k].SetDirtyParent(x.markDirty)
			v[k].MarkDirtyAll()
		}
	}
	x.markDirty()
}

func (x *RepeatedMessage[T, E]) Reverse() {
	if len(x.data) < 2 {
		return
	}
	slices.Reverse(x.data)
	x.markDirty()
}

func (x *RepeatedMessage[T, E]) All() iter.Seq2[int, E] {
	return slices.All(x.data)
}

func (x *RepeatedMessage[T, E]) Backward() iter.Seq2[int, E] {
	return slices.Backward(x.data)
}

func (x *RepeatedMessage[T, E]) Values() iter.Seq[E] {
	return slices.Values(x.data)
}

func (x *RepeatedMessage[T, E]) markDirty() {
	if x.dirty {
		return
	}
	x.dirty = true
	x.dirtyParent.Invoke()
}

func (x *RepeatedMessage[T, E]) ClearDirty() {
	for _, v := range x.data {
		if v != nil {
			v.ClearDirty()
		}
	}
	x.dirty = false
}

func (x *RepeatedMessage[T, E]) Marshal(b []byte) ([]byte, error) {
	if len(x.data) == 0 {
		return b, nil
	}
	for _, v := range x.data {
		var err error
		if b, err = wire.AppendMessage(b, v); err != nil {
			return b, err
		}
	}
	return b, nil
}

func (x *RepeatedMessage[T, E]) MarshalDirty(b []byte) ([]byte, error) {
	return x.Marshal(b)
}

func (x *RepeatedMessage[T, E]) Unmarshal(b []byte) error {
	x.Clear()
	for len(b) > 0 {
		var v E = new(T)
		n, err := wire.ConsumeMessage(b, v)
		if err != nil {
			return err
		}
		b = b[n:]
		x.data = append(x.data, v)
		v.SetDirtyParent(x.markDirty)
		v.MarkDirtyAll()
	}
	return nil
}

func (x *RepeatedMessage[T, E]) MarshalJSONIndent(b []byte, prefix string, indent string) ([]byte, error) {
	if len(x.data) == 0 {
		return append(b, "[]"...), nil
	}
	var err error
	b = append(b, "[\n"...)
	for i, v := range x.data {
		b = append(b, prefix+indent...)
		b, err = MarshalJSONIndent(b, v, prefix+indent, indent)
		if err != nil {
			return nil, err
		}
		if i+1 < len(x.data) {
			b = append(b, ',')
		}
		b = append(b, '\n')
	}
	b = append(b, prefix...)
	b = append(b, ']')
	return b, nil
}
