package kdsjson

import (
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

func WriteBoolPropertyName(e *Encoder, name bool) error {
	if name {
		return e.WritePropertyName("true")
	} else {
		return e.WritePropertyName("false")
	}
}

func WriteInt32PropertyName(e *Encoder, name int32) error {
	return e.WritePropertyName(strconv.FormatInt(int64(name), 10))
}

func WriteUint32PropertyName(e *Encoder, name uint32) error {
	return e.WritePropertyName(strconv.FormatUint(uint64(name), 10))
}

func WriteInt64PropertyName(e *Encoder, name int64) error {
	return e.WritePropertyName(strconv.FormatInt(name, 10))
}

func WriteUint64PropertyName(e *Encoder, name uint64) error {
	return e.WritePropertyName(strconv.FormatUint(name, 10))
}

func WritePropertyName(e *Encoder, name string) error {
	return e.WritePropertyName(name)
}

func WriteEnum[T ~int32](e *Encoder, name string, v T) error {
	return e.WriteInt32(name, int32(v))
}

func WriteEnumValue[T ~int32](e *Encoder, v T) {
	e.WriteInt32Value(int32(v))
}

func WriteBoolValue(e *Encoder, v bool) {
	e.WriteBoolValue(v)
}

func WriteInt32Value(e *Encoder, v int32) {
	e.WriteInt32Value(v)
}

func WriteUint32Value(e *Encoder, v uint32) {
	e.WriteUint32Value(v)
}

func WriteInt64Value(e *Encoder, v int64) {
	e.WriteInt64Value(v)
}

func WriteUint64Value(e *Encoder, v uint64) {
	e.WriteUint64Value(v)
}

func WriteFloat32Value(e *Encoder, v float32) {
	e.WriteFloat32Value(v)
}

func WriteFloat64Value(e *Encoder, v float64) {
	e.WriteFloat64Value(v)
}

func WriteBytesValue(e *Encoder, v []byte) {
	e.WriteBytesValue(v)
}

func WriteStringValue(e *Encoder, v string) error {
	return e.WriteStringValue(v)
}

func WriteTimestampValue(e *Encoder, v time.Time) {
	e.WriteTimestampValue(v)
}

func WriteDurationValue(e *Encoder, v time.Duration) {
	e.WriteDurationValue(v)
}

func WriteEmptyValue(e *Encoder, v struct{}) {
	e.WriteEmptyValue(v)
}
