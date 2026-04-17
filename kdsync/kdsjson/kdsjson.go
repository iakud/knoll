package kdsjson

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

type Marshaler interface {
	WriteJSON(e *Encoder) error
}

func Marshal(v Marshaler) (string, error) {
	var e Encoder
	err := v.WriteJSON(&e)
	return e.String(), err
}

func MarshalIndent(v Marshaler) (string, error) {
	e := Encoder{indented: true, indentLength: 2}
	err := v.WriteJSON(&e)
	return e.String(), err
}

func WriteBoolName(e *Encoder, name bool) error {
	if name {
		return e.WriteName("true")
	} else {
		return e.WriteName("false")
	}
}

func WriteInt32Name(e *Encoder, name int32) error {
	return e.WriteName(strconv.FormatInt(int64(name), 10))
}

func WriteUint32Name(e *Encoder, name uint32) error {
	return e.WriteName(strconv.FormatUint(uint64(name), 10))
}

func WriteInt64Name(e *Encoder, name int64) error {
	return e.WriteName(strconv.FormatInt(name, 10))
}

func WriteUint64Name(e *Encoder, name uint64) error {
	return e.WriteName(strconv.FormatUint(name, 10))
}

func WriteStringName(e *Encoder, name string) error {
	return e.WriteName(name)
}

func WriteEnum[T ~int32](e *Encoder, name string, v T) error {
	return e.WriteInt32(name, int32(v))
}

func WriteEnumValue[T ~int32](e *Encoder, v T) {
	e.WriteInt32Value(int32(v))
}

func WriteBoolValue(e *Encoder, v bool) error {
	e.WriteBoolValue(v)
	return nil
}

func WriteInt32Value(e *Encoder, v int32) error {
	e.WriteInt32Value(v)
	return nil
}

func WriteUint32Value(e *Encoder, v uint32) error {
	e.WriteUint32Value(v)
	return nil
}

func WriteInt64Value(e *Encoder, v int64) error {
	e.WriteInt64Value(v)
	return nil
}

func WriteUint64Value(e *Encoder, v uint64) error {
	e.WriteUint64Value(v)
	return nil
}

func WriteFloat32Value(e *Encoder, v float32) error {
	e.WriteFloat32Value(v)
	return nil
}

func WriteFloat64Value(e *Encoder, v float64) error {
	e.WriteFloat64Value(v)
	return nil
}

func WriteBytesValue(e *Encoder, v []byte) error {
	e.WriteBytesValue(v)
	return nil
}

func WriteStringValue(e *Encoder, v string) error {
	return e.WriteStringValue(v)
}

func WriteTimestampValue(e *Encoder, v time.Time) error {
	e.WriteTimestampValue(v)
	return nil
}

func WriteDurationValue(e *Encoder, v time.Duration) error {
	e.WriteDurationValue(v)
	return nil
}

func WriteEmptyValue(e *Encoder, v struct{}) error {
	e.WriteEmptyValue(v)
	return nil
}

type WriterFunc[T any] func(e *Encoder, v T) error

func NameWriter[T comparable]() WriterFunc[T] {
	switch reflect.TypeFor[T]() {
	case reflect.TypeFor[bool]():
		return (any)(WriteBoolName).(func(e *Encoder, v T) error)
	case reflect.TypeFor[int32]():
		return (any)(WriteInt32Name).(func(e *Encoder, v T) error)
	case reflect.TypeFor[uint32]():
		return (any)(WriteUint32Name).(func(e *Encoder, v T) error)
	case reflect.TypeFor[int64]():
		return (any)(WriteInt64Name).(func(e *Encoder, v T) error)
	case reflect.TypeFor[uint64]():
		return (any)(WriteUint64Name).(func(e *Encoder, v T) error)
	case reflect.TypeFor[string]():
		return (any)(WriteStringName).(func(e *Encoder, v T) error)
	default:
		return func(e *Encoder, v T) error { return WriteStringValue(e, fmt.Sprint(v)) }
	}
}

func ValueWriter[T any]() WriterFunc[T] {
	switch reflect.TypeFor[T]() {
	case reflect.TypeFor[bool]():
		return (any)(WriteBoolValue).(func(e *Encoder, v T) error)
	case reflect.TypeFor[int32]():
		return (any)(WriteInt32Value).(func(e *Encoder, v T) error)
	case reflect.TypeFor[uint32]():
		return (any)(WriteUint32Value).(func(e *Encoder, v T) error)
	case reflect.TypeFor[int64]():
		return (any)(WriteInt64Value).(func(e *Encoder, v T) error)
	case reflect.TypeFor[uint64]():
		return (any)(WriteUint64Value).(func(e *Encoder, v T) error)
	case reflect.TypeFor[float32]():
		return (any)(WriteFloat32Value).(func(e *Encoder, v T) error)
	case reflect.TypeFor[float64]():
		return (any)(WriteFloat64Value).(func(e *Encoder, v T) error)
	case reflect.TypeFor[string]():
		return (any)(WriteStringValue).(func(e *Encoder, v T) error)
	case reflect.TypeFor[[]byte]():
		return (any)(WriteBytesValue).(func(e *Encoder, v T) error)
	case reflect.TypeFor[time.Time]():
		return (any)(WriteTimestampValue).(func(e *Encoder, v T) error)
	case reflect.TypeFor[time.Duration]():
		return (any)(WriteDurationValue).(func(e *Encoder, v T) error)
	case reflect.TypeFor[struct{}]():
		return (any)(WriteEmptyValue).(func(e *Encoder, v T) error)
	default:
		return func(e *Encoder, v T) error { return WriteStringValue(e, fmt.Sprint(v)) }
	}
}
