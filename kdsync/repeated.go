package kdsync

import (
	"bytes"
	"encoding/json"
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
	MarshalJSON() ([]byte, error)
	String(indent string) string
}

type repeated[E any] struct {
	data []E

	dirty       bool
	dirtyParent DirtyFunc
}

// Field repeated check
var _ Repeated[bool] = (*FieldRepeated[bool])(nil)
var _ Repeated[int32] = (*FieldRepeated[int32])(nil)
var _ Repeated[uint32] = (*FieldRepeated[uint32])(nil)
var _ Repeated[int64] = (*FieldRepeated[int64])(nil)
var _ Repeated[uint64] = (*FieldRepeated[uint64])(nil)
var _ Repeated[float32] = (*FieldRepeated[float32])(nil)
var _ Repeated[float64] = (*FieldRepeated[float64])(nil)
var _ Repeated[string] = (*FieldRepeated[string])(nil)
var _ Repeated[time.Duration] = (*FieldRepeated[time.Duration])(nil)
var _ Repeated[struct{}] = (*FieldRepeated[struct{}])(nil)

// Bytes repeated check
var _ Repeated[[]byte] = (*BytesRepeated)(nil)

// Timestamp repeated check
var _ Repeated[time.Time] = (*TimestampRepeated)(nil)

type Int32Repeated FieldRepeated[int32]

// Field repeated
type FieldRepeated[T Field] struct {
	repeated[T]
	fieldCodec FieldCodec[T]
}

func NewFieldRepeated[T Field](dirtyParent DirtyFunc, fieldCodec FieldCodec[T]) FieldRepeated[T] {
	return FieldRepeated[T]{repeated: repeated[T]{dirtyParent: dirtyParent}, fieldCodec: fieldCodec}
}

func (x *FieldRepeated[T]) Len() int {
	return len(x.data)
}

func (x *FieldRepeated[T]) Clear() {
	if len(x.data) == 0 {
		return
	}
	clear(x.data)
	x.data = x.data[:0]
	x.markDirty()
}

func (x *FieldRepeated[T]) Get(i int) T {
	return x.data[i]
}

func (x *FieldRepeated[T]) Set(i int, v T) {
	if v == x.data[i] {
		return
	}
	x.data[i] = v
	x.markDirty()
}

func (x *FieldRepeated[T]) Append(v ...T) {
	if len(v) == 0 {
		return
	}
	x.data = append(x.data, v...)
	x.markDirty()
}

func (x *FieldRepeated[T]) Index(v T) int {
	return slices.Index(x.data, v)
}

func (x *FieldRepeated[T]) IndexFunc(f func(T) bool) int {
	for i := range x.data {
		if f(x.data[i]) {
			return i
		}
	}
	return -1
}

func (x *FieldRepeated[T]) Contains(v T) bool {
	return x.Index(v) >= 0
}

func (x *FieldRepeated[T]) ContainsFunc(f func(T) bool) bool {
	return x.IndexFunc(f) >= 0
}

func (x *FieldRepeated[T]) Insert(i int, v ...T) {
	_ = x.data[i:] // bounds check
	if len(v) == 0 {
		return
	}
	x.data = slices.Insert(x.data, i, v...)
	x.markDirty()
}

func (x *FieldRepeated[T]) Delete(i, j int) {
	_ = x.data[i:j:len(x.data)] // bounds check
	if i == j {
		return
	}
	x.data = slices.Delete(x.data, i, j)
	x.markDirty()
}

func (x *FieldRepeated[T]) DeleteFunc(del func(T) bool) {
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

func (x *FieldRepeated[T]) Replace(i, j int, v ...T) {
	_ = x.data[i:j] // bounds check
	if i == j && len(v) == 0 {
		return
	}
	x.data = slices.Replace(x.data, i, j, v...)
	x.markDirty()
}

func (x *FieldRepeated[T]) Reverse() {
	if len(x.data) < 2 {
		return
	}
	slices.Reverse(x.data)
	x.markDirty()
}

func (x *FieldRepeated[T]) All() iter.Seq2[int, T] {
	return slices.All(x.data)
}

func (x *FieldRepeated[T]) Backward() iter.Seq2[int, T] {
	return slices.Backward(x.data)
}

func (x *FieldRepeated[T]) Values() iter.Seq[T] {
	return slices.Values(x.data)
}

func (x *FieldRepeated[T]) markDirty() {
	if x.dirty {
		return
	}
	x.dirty = true
	x.dirtyParent.Invoke()
}

func (x *FieldRepeated[T]) ClearDirty() {
	x.dirty = false
}

func (x *FieldRepeated[T]) Marshal(b []byte) ([]byte, error) {
	if len(x.data) == 0 {
		return b, nil
	}
	for _, v := range x.data {
		b = x.fieldCodec.MarshalFunc(b, v)
	}
	return b, nil
}

func (x *FieldRepeated[T]) MarshalDirty(b []byte) ([]byte, error) {
	return x.Marshal(b)
}

func (x *FieldRepeated[T]) Unmarshal(b []byte) error {
	x.Clear()
	for len(b) > 0 {
		v, n, err := x.fieldCodec.UnmarshalFunc(b)
		if err != nil {
			return err
		}
		b = b[n:]
		x.data = append(x.data, v)
	}
	return nil
}

func (x *FieldRepeated[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.data)
}

func (x *FieldRepeated[E]) String(indent string) string {
	if len(x.data) == 0 {
		return "[]"
	}
	var b []byte
	b = append(b, "[\n"...)
	for i, v := range x.data {
		b = append(b, (indent + "  ")...)
		b = append(b, wire.Format(v)...)
		if i+1 < len(x.data) {
			b = append(b, ',')
		}
		b = append(b, '\n')
	}
	b = append(b, (indent + "]")...)
	return string(b)
}

// Bytes repeated
type BytesRepeated struct {
	repeated[[]byte]
}

func NewBytesRepeated(dirtyParent DirtyFunc) BytesRepeated {
	return BytesRepeated{repeated: repeated[[]byte]{dirtyParent: dirtyParent}}
}

func (x *BytesRepeated) Len() int {
	return len(x.data)
}

func (x *BytesRepeated) Clear() {
	if len(x.data) == 0 {
		return
	}
	clear(x.data)
	x.data = x.data[:0]
	x.markDirty()
}

func (x *BytesRepeated) Get(i int) []byte {
	return x.data[i]
}

func (x *BytesRepeated) Set(i int, v []byte) {
	if bytes.Equal(v, x.data[i]) {
		return
	}
	x.data[i] = bytes.Clone(v)
	x.markDirty()
}

func (x *BytesRepeated) Append(v ...[]byte) {
	if len(v) == 0 {
		return
	}
	for i := range v {
		v[i] = bytes.Clone(v[i])
	}
	x.data = append(x.data, v...)
	x.markDirty()
}

func (x *BytesRepeated) Index(v []byte) int {
	for i := range x.data {
		if bytes.Equal(v, x.data[i]) {
			return i
		}
	}
	return -1
}

func (x *BytesRepeated) IndexFunc(f func([]byte) bool) int {
	for i := range x.data {
		if f(x.data[i]) {
			return i
		}
	}
	return -1
}

func (x *BytesRepeated) Contains(v []byte) bool {
	return x.Index(v) >= 0
}

func (x *BytesRepeated) ContainsFunc(f func([]byte) bool) bool {
	return x.IndexFunc(f) >= 0
}

func (x *BytesRepeated) Insert(i int, v ...[]byte) {
	_ = x.data[i:] // bounds check
	if len(v) == 0 {
		return
	}
	for i := range v {
		v[i] = bytes.Clone(x.data[i])
	}
	x.data = slices.Insert(x.data, i, v...)
	x.markDirty()
}

func (x *BytesRepeated) Delete(i, j int) {
	_ = x.data[i:j:len(x.data)] // bounds check
	if i == j {
		return
	}
	x.data = slices.Delete(x.data, i, j)
	x.markDirty()
}

func (x *BytesRepeated) DeleteFunc(del func([]byte) bool) {
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

func (x *BytesRepeated) Replace(i, j int, v ...[]byte) {
	_ = x.data[i:j] // bounds check
	if i == j && len(v) == 0 {
		return
	}
	for i := range v {
		v[i] = bytes.Clone(x.data[i])
	}
	x.data = slices.Replace(x.data, i, j, v...)
	x.markDirty()
}

func (x *BytesRepeated) Reverse() {
	if len(x.data) < 2 {
		return
	}
	slices.Reverse(x.data)
	x.markDirty()
}

func (x *BytesRepeated) All() iter.Seq2[int, []byte] {
	return slices.All(x.data)
}

func (x *BytesRepeated) Backward() iter.Seq2[int, []byte] {
	return slices.Backward(x.data)
}

func (x *BytesRepeated) Values() iter.Seq[[]byte] {
	return slices.Values(x.data)
}

func (x *BytesRepeated) markDirty() {
	if x.dirty {
		return
	}
	x.dirty = true
	x.dirtyParent.Invoke()
}

func (x *BytesRepeated) ClearDirty() {
	x.dirty = false
}

func (x *BytesRepeated) Marshal(b []byte) ([]byte, error) {
	if len(x.data) == 0 {
		return b, nil
	}
	for _, v := range x.data {
		b = wire.AppendBytes(b, v)
	}
	return b, nil
}

func (x *BytesRepeated) MarshalDirty(b []byte) ([]byte, error) {
	return x.Marshal(b)
}

func (x *BytesRepeated) Unmarshal(b []byte) error {
	x.Clear()
	for len(b) > 0 {
		v, n, err := wire.ConsumeBytes(b)
		if err != nil {
			return err
		}
		b = b[n:]
		x.data = append(x.data, bytes.Clone(v))
	}
	return nil
}

func (x *BytesRepeated) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.data)
}

func (x *BytesRepeated) String(indent string) string {
	if len(x.data) == 0 {
		return "[]"
	}
	var b []byte
	b = append(b, "[\n"...)
	for i, v := range x.data {
		b = append(b, (indent + "  ")...)
		b = append(b, wire.Format(v)...)
		if i+1 < len(x.data) {
			b = append(b, ',')
		}
		b = append(b, '\n')
	}
	b = append(b, (indent + "]")...)
	return string(b)
}

// Timestamp repeated
type TimestampRepeated struct {
	repeated[time.Time]
}

func NewTimestampRepeated(dirtyParent DirtyFunc) TimestampRepeated {
	return TimestampRepeated{repeated: repeated[time.Time]{dirtyParent: dirtyParent}}
}

func (x *TimestampRepeated) Len() int {
	return len(x.data)
}

func (x *TimestampRepeated) Clear() {
	if len(x.data) == 0 {
		return
	}
	clear(x.data)
	x.data = x.data[:0]
	x.markDirty()
}

func (x *TimestampRepeated) Get(i int) time.Time {
	return x.data[i]
}

func (x *TimestampRepeated) Set(i int, v time.Time) {
	if v.Equal(x.data[i]) {
		return
	}
	x.data[i] = v
	x.markDirty()
}

func (x *TimestampRepeated) Append(v ...time.Time) {
	if len(v) == 0 {
		return
	}
	x.data = append(x.data, v...)
	x.markDirty()
}

func (x *TimestampRepeated) Index(v time.Time) int {
	for i := range x.data {
		if v.Equal(x.data[i]) {
			return i
		}
	}
	return -1
}

func (x *TimestampRepeated) IndexFunc(f func(time.Time) bool) int {
	for i := range x.data {
		if f(x.data[i]) {
			return i
		}
	}
	return -1
}

func (x *TimestampRepeated) Contains(v time.Time) bool {
	return x.Index(v) >= 0
}

func (x *TimestampRepeated) ContainsFunc(f func(time.Time) bool) bool {
	return x.IndexFunc(f) >= 0
}

func (x *TimestampRepeated) Insert(i int, v ...time.Time) {
	_ = x.data[i:] // bounds check
	if len(v) == 0 {
		return
	}
	x.data = slices.Insert(x.data, i, v...)
	x.markDirty()
}

func (x *TimestampRepeated) Delete(i, j int) {
	_ = x.data[i:j:len(x.data)] // bounds check
	if i == j {
		return
	}
	x.data = slices.Delete(x.data, i, j)
	x.markDirty()
}

func (x *TimestampRepeated) DeleteFunc(del func(time.Time) bool) {
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

func (x *TimestampRepeated) Replace(i, j int, v ...time.Time) {
	_ = x.data[i:j] // bounds check
	if i == j && len(v) == 0 {
		return
	}
	x.data = slices.Replace(x.data, i, j, v...)
	x.markDirty()
}

func (x *TimestampRepeated) Reverse() {
	if len(x.data) < 2 {
		return
	}
	slices.Reverse(x.data)
	x.markDirty()
}

func (x *TimestampRepeated) All() iter.Seq2[int, time.Time] {
	return slices.All(x.data)
}

func (x *TimestampRepeated) Backward() iter.Seq2[int, time.Time] {
	return slices.Backward(x.data)
}

func (x *TimestampRepeated) Values() iter.Seq[time.Time] {
	return slices.Values(x.data)
}

func (x *TimestampRepeated) markDirty() {
	if x.dirty {
		return
	}
	x.dirty = true
	x.dirtyParent.Invoke()
}

func (x *TimestampRepeated) ClearDirty() {
	x.dirty = false
}

func (x *TimestampRepeated) Marshal(b []byte) ([]byte, error) {
	if len(x.data) == 0 {
		return b, nil
	}
	for _, v := range x.data {
		b = wire.AppendTimestamp(b, v)
	}
	return b, nil
}

func (x *TimestampRepeated) MarshalDirty(b []byte) ([]byte, error) {
	return x.Marshal(b)
}

func (x *TimestampRepeated) Unmarshal(b []byte) error {
	x.Clear()
	for len(b) > 0 {
		v, n, err := wire.ConsumeTimestamp(b)
		if err != nil {
			return err
		}
		b = b[n:]
		x.data = append(x.data, v)
	}
	return nil
}

func (x *TimestampRepeated) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.data)
}

func (x *TimestampRepeated) String(indent string) string {
	if len(x.data) == 0 {
		return "[]"
	}
	var b []byte
	b = append(b, "[\n"...)
	for i, v := range x.data {
		b = append(b, (indent + "  ")...)
		b = append(b, wire.Format(v)...)
		if i+1 < len(x.data) {
			b = append(b, ',')
		}
		b = append(b, '\n')
	}
	b = append(b, (indent + "]")...)
	return string(b)
}

// Message repeated
type MessageRepeated[T any, E Message[T]] struct {
	repeated[E]
}

func NewMessageRepeated[T any, E Message[T]](dirtyParent DirtyFunc) MessageRepeated[T, E] {
	return MessageRepeated[T, E]{repeated: repeated[E]{dirtyParent: dirtyParent}}
}

func (x *MessageRepeated[T, E]) Len() int {
	return len(x.data)
}

func (x *MessageRepeated[T, E]) Clear() {
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

func (x *MessageRepeated[T, E]) Get(i int) E {
	return x.data[i]
}

func (x *MessageRepeated[T, E]) Set(i int, v E) {
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

func (x *MessageRepeated[T, E]) Append(v ...E) {
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

func (x *MessageRepeated[T, E]) Index(v E) int {
	for i := range x.data {
		if x.data[i] == v {
			return i
		}
	}
	return -1
}

func (x *MessageRepeated[T, E]) IndexFunc(f func(E) bool) int {
	for i := range x.data {
		if f(x.data[i]) {
			return i
		}
	}
	return -1
}

func (x *MessageRepeated[T, E]) Contains(v E) bool {
	return x.Index(v) >= 0
}

func (x *MessageRepeated[T, E]) ContainsFunc(f func(E) bool) bool {
	return x.IndexFunc(f) >= 0
}

func (x *MessageRepeated[T, E]) Insert(i int, v ...E) {
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

func (x *MessageRepeated[T, E]) Delete(i, j int) {
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

func (x *MessageRepeated[T, E]) DeleteFunc(del func(E) bool) {
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

func (x *MessageRepeated[T, E]) Replace(i, j int, v ...E) {
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

func (x *MessageRepeated[T, E]) Reverse() {
	if len(x.data) < 2 {
		return
	}
	slices.Reverse(x.data)
	x.markDirty()
}

func (x *MessageRepeated[T, E]) All() iter.Seq2[int, E] {
	return slices.All(x.data)
}

func (x *MessageRepeated[T, E]) Backward() iter.Seq2[int, E] {
	return slices.Backward(x.data)
}

func (x *MessageRepeated[T, E]) Values() iter.Seq[E] {
	return slices.Values(x.data)
}

func (x *MessageRepeated[T, E]) markDirty() {
	if x.dirty {
		return
	}
	x.dirty = true
	x.dirtyParent.Invoke()
}

func (x *MessageRepeated[T, E]) ClearDirty() {
	for _, v := range x.data {
		if v != nil {
			v.ClearDirty()
		}
	}
	x.dirty = false
}

func (x *MessageRepeated[T, E]) Marshal(b []byte) ([]byte, error) {
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

func (x *MessageRepeated[T, E]) MarshalDirty(b []byte) ([]byte, error) {
	return x.Marshal(b)
}

func (x *MessageRepeated[T, E]) Unmarshal(b []byte) error {
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

func (x *MessageRepeated[T, E]) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.data)
}

func (x *MessageRepeated[T, E]) String(indent string) string {
	if len(x.data) == 0 {
		return "[]"
	}
	var b []byte
	b = append(b, "[\n"...)
	for i, v := range x.data {
		b = append(b, (indent + "  " + v.String(indent+"  "))...)
		if i+1 < len(x.data) {
			b = append(b, ',')
		}
		b = append(b, '\n')
	}
	b = append(b, (indent + "]")...)
	return string(b)
}
