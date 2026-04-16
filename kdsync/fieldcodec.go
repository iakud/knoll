package kdsync

import (
	"bytes"
	"cmp"
	"strings"
	"time"

	"github.com/iakud/knoll/kdsync/kdsjson"
	"github.com/iakud/knoll/kdsync/wire"
)

// Field Codec
type FieldCodec[T any] interface {
	WireType() wire.Type
	Compare(a, b T) int
	Marshal(b []byte, v T) []byte
	Unmarshal(b []byte) (T, int, error)
	WriteJson(e *kdsjson.Encoder, v T) error
}

// Check Codecs
var _ FieldCodec[bool] = (*BoolKeyCodec)(nil)
var _ FieldCodec[int32] = (*Int32KeyCodec)(nil)
var _ FieldCodec[uint32] = (*Uint32KeyCodec)(nil)
var _ FieldCodec[int64] = (*Int64KeyCodec)(nil)
var _ FieldCodec[uint64] = (*Uint64KeyCodec)(nil)
var _ FieldCodec[string] = (*StringKeyCodec)(nil)
var _ FieldCodec[bool] = (*BoolValueCodec)(nil)
var _ FieldCodec[int32] = (*Int32ValueCodec)(nil)
var _ FieldCodec[uint32] = (*Uint32ValueCodec)(nil)
var _ FieldCodec[int64] = (*Int64ValueCodec)(nil)
var _ FieldCodec[uint64] = (*Uint64ValueCodec)(nil)
var _ FieldCodec[float32] = (*Float32ValueCodec)(nil)
var _ FieldCodec[float64] = (*Float64ValueCodec)(nil)
var _ FieldCodec[string] = (*StringValueCodec)(nil)
var _ FieldCodec[[]byte] = (*BytesValueCodec)(nil)
var _ FieldCodec[time.Time] = (*TimestampValueCodec)(nil)
var _ FieldCodec[time.Duration] = (*DurationValueCodec)(nil)
var _ FieldCodec[struct{}] = (*EmptyValueCodec)(nil)

// Key Codecs

type BoolKeyCodec struct{}

func (c *BoolKeyCodec) WireType() wire.Type { return wire.VarintType }
func (c *BoolKeyCodec) Compare(a, b bool) int {
	if a {
		return 1
	} else if b {
		return -1
	}
	return 0
}
func (c *BoolKeyCodec) Marshal(b []byte, v bool) []byte       { return wire.AppendBool(b, v) }
func (c *BoolKeyCodec) Unmarshal(b []byte) (bool, int, error) { return wire.ConsumeBool(b) }
func (c *BoolKeyCodec) WriteJson(e *kdsjson.Encoder, v bool) error {
	return kdsjson.WriteBoolPropertyName(e, v)
}

type Int32KeyCodec struct{}

func (c *Int32KeyCodec) WireType() wire.Type                    { return wire.VarintType }
func (c *Int32KeyCodec) Compare(a, b int32) int                 { return cmp.Compare(a, b) }
func (c *Int32KeyCodec) Marshal(b []byte, v int32) []byte       { return wire.AppendInt32(b, v) }
func (c *Int32KeyCodec) Unmarshal(b []byte) (int32, int, error) { return wire.ConsumeInt32(b) }
func (c *Int32KeyCodec) WriteJson(e *kdsjson.Encoder, v int32) error {
	return kdsjson.WriteInt32PropertyName(e, v)
}

type Uint32KeyCodec struct{}

func (c *Uint32KeyCodec) WireType() wire.Type                     { return wire.VarintType }
func (c *Uint32KeyCodec) Compare(a, b uint32) int                 { return cmp.Compare(a, b) }
func (c *Uint32KeyCodec) Marshal(b []byte, v uint32) []byte       { return wire.AppendUint32(b, v) }
func (c *Uint32KeyCodec) Unmarshal(b []byte) (uint32, int, error) { return wire.ConsumeUint32(b) }
func (c *Uint32KeyCodec) WriteJson(e *kdsjson.Encoder, v uint32) error {
	return kdsjson.WriteUint32PropertyName(e, v)
}

type Int64KeyCodec struct{}

func (c *Int64KeyCodec) WireType() wire.Type                    { return wire.VarintType }
func (c *Int64KeyCodec) Compare(a, b int64) int                 { return cmp.Compare(a, b) }
func (c *Int64KeyCodec) Marshal(b []byte, v int64) []byte       { return wire.AppendInt64(b, v) }
func (c *Int64KeyCodec) Unmarshal(b []byte) (int64, int, error) { return wire.ConsumeInt64(b) }
func (c *Int64KeyCodec) WriteJson(e *kdsjson.Encoder, v int64) error {
	return kdsjson.WriteInt64PropertyName(e, v)
}

type Uint64KeyCodec struct{}

func (c *Uint64KeyCodec) WireType() wire.Type                     { return wire.VarintType }
func (c *Uint64KeyCodec) Compare(a, b uint64) int                 { return cmp.Compare(a, b) }
func (c *Uint64KeyCodec) Marshal(b []byte, v uint64) []byte       { return wire.AppendUint64(b, v) }
func (c *Uint64KeyCodec) Unmarshal(b []byte) (uint64, int, error) { return wire.ConsumeUint64(b) }
func (c *Uint64KeyCodec) WriteJson(e *kdsjson.Encoder, v uint64) error {
	return kdsjson.WriteUint64PropertyName(e, v)
}

type StringKeyCodec struct{}

func (c *StringKeyCodec) WireType() wire.Type                     { return wire.BytesType }
func (c *StringKeyCodec) Compare(a, b string) int                 { return strings.Compare(a, b) }
func (c *StringKeyCodec) Marshal(b []byte, v string) []byte       { return wire.AppendString(b, v) }
func (c *StringKeyCodec) Unmarshal(b []byte) (string, int, error) { return wire.ConsumeString(b) }
func (c *StringKeyCodec) WriteJson(e *kdsjson.Encoder, v string) error {
	return kdsjson.WritePropertyName(e, v)
}

// Value Codecs

type BoolValueCodec struct{}

func (c *BoolValueCodec) WireType() wire.Type { return wire.VarintType }
func (c *BoolValueCodec) Compare(a, b bool) int {
	if a {
		return 1
	} else if b {
		return -1
	}
	return 0
}
func (c *BoolValueCodec) Marshal(b []byte, v bool) []byte       { return wire.AppendBool(b, v) }
func (c *BoolValueCodec) Unmarshal(b []byte) (bool, int, error) { return wire.ConsumeBool(b) }
func (c *BoolValueCodec) WriteJson(e *kdsjson.Encoder, v bool) error {
	kdsjson.WriteBoolValue(e, v)
	return nil
}

type Int32ValueCodec struct{}

func (c *Int32ValueCodec) WireType() wire.Type                    { return wire.VarintType }
func (c *Int32ValueCodec) Compare(a, b int32) int                 { return cmp.Compare(a, b) }
func (c *Int32ValueCodec) Marshal(b []byte, v int32) []byte       { return wire.AppendInt32(b, v) }
func (c *Int32ValueCodec) Unmarshal(b []byte) (int32, int, error) { return wire.ConsumeInt32(b) }
func (c *Int32ValueCodec) WriteJson(e *kdsjson.Encoder, v int32) error {
	kdsjson.WriteInt32Value(e, v)
	return nil
}

type Uint32ValueCodec struct{}

func (c *Uint32ValueCodec) WireType() wire.Type                     { return wire.VarintType }
func (c *Uint32ValueCodec) Compare(a, b uint32) int                 { return cmp.Compare(a, b) }
func (c *Uint32ValueCodec) Marshal(b []byte, v uint32) []byte       { return wire.AppendUint32(b, v) }
func (c *Uint32ValueCodec) Unmarshal(b []byte) (uint32, int, error) { return wire.ConsumeUint32(b) }
func (c *Uint32ValueCodec) WriteJson(e *kdsjson.Encoder, v uint32) error {
	kdsjson.WriteUint32Value(e, v)
	return nil
}

type Int64ValueCodec struct{}

func (c *Int64ValueCodec) WireType() wire.Type                    { return wire.VarintType }
func (c *Int64ValueCodec) Compare(a, b int64) int                 { return cmp.Compare(a, b) }
func (c *Int64ValueCodec) Marshal(b []byte, v int64) []byte       { return wire.AppendInt64(b, v) }
func (c *Int64ValueCodec) Unmarshal(b []byte) (int64, int, error) { return wire.ConsumeInt64(b) }
func (c *Int64ValueCodec) WriteJson(e *kdsjson.Encoder, v int64) error {
	kdsjson.WriteInt64Value(e, v)
	return nil
}

type Uint64ValueCodec struct{}

func (c *Uint64ValueCodec) WireType() wire.Type                     { return wire.VarintType }
func (c *Uint64ValueCodec) Compare(a, b uint64) int                 { return cmp.Compare(a, b) }
func (c *Uint64ValueCodec) Marshal(b []byte, v uint64) []byte       { return wire.AppendUint64(b, v) }
func (c *Uint64ValueCodec) Unmarshal(b []byte) (uint64, int, error) { return wire.ConsumeUint64(b) }
func (c *Uint64ValueCodec) WriteJson(e *kdsjson.Encoder, v uint64) error {
	kdsjson.WriteUint64Value(e, v)
	return nil
}

type Float32ValueCodec struct{}

func (c *Float32ValueCodec) WireType() wire.Type                      { return wire.Fixed32Type }
func (c *Float32ValueCodec) Compare(a, b float32) int                 { return cmp.Compare(a, b) }
func (c *Float32ValueCodec) Marshal(b []byte, v float32) []byte       { return wire.AppendFloat(b, v) }
func (c *Float32ValueCodec) Unmarshal(b []byte) (float32, int, error) { return wire.ConsumeFloat(b) }
func (c *Float32ValueCodec) WriteJson(e *kdsjson.Encoder, v float32) error {
	kdsjson.WriteFloat32Value(e, v)
	return nil
}

type Float64ValueCodec struct{}

func (c *Float64ValueCodec) WireType() wire.Type                      { return wire.Fixed64Type }
func (c *Float64ValueCodec) Compare(a, b float64) int                 { return cmp.Compare(a, b) }
func (c *Float64ValueCodec) Marshal(b []byte, v float64) []byte       { return wire.AppendDouble(b, v) }
func (c *Float64ValueCodec) Unmarshal(b []byte) (float64, int, error) { return wire.ConsumeDouble(b) }
func (c *Float64ValueCodec) WriteJson(e *kdsjson.Encoder, v float64) error {
	kdsjson.WriteFloat64Value(e, v)
	return nil
}

type StringValueCodec struct{}

func (c *StringValueCodec) WireType() wire.Type                     { return wire.BytesType }
func (c *StringValueCodec) Compare(a, b string) int                 { return strings.Compare(a, b) }
func (c *StringValueCodec) Marshal(b []byte, v string) []byte       { return wire.AppendString(b, v) }
func (c *StringValueCodec) Unmarshal(b []byte) (string, int, error) { return wire.ConsumeString(b) }
func (c *StringValueCodec) WriteJson(e *kdsjson.Encoder, v string) error {
	return kdsjson.WriteStringValue(e, v)
}

type BytesValueCodec struct{}

func (c *BytesValueCodec) WireType() wire.Type                     { return wire.BytesType }
func (c *BytesValueCodec) Compare(a, b []byte) int                 { return bytes.Compare(a, b) }
func (c *BytesValueCodec) Marshal(b []byte, v []byte) []byte       { return wire.AppendBytes(b, v) }
func (c *BytesValueCodec) Unmarshal(b []byte) ([]byte, int, error) { return wire.ConsumeBytes(b) }
func (c *BytesValueCodec) WriteJson(e *kdsjson.Encoder, v []byte) error {
	kdsjson.WriteBytesValue(e, v)
	return nil
}

type TimestampValueCodec struct{}

func (c *TimestampValueCodec) WireType() wire.Type        { return wire.BytesType }
func (c *TimestampValueCodec) Compare(a, b time.Time) int { return a.Compare(b) }
func (c *TimestampValueCodec) Marshal(b []byte, v time.Time) []byte {
	return wire.AppendTimestamp(b, v)
}
func (c *TimestampValueCodec) Unmarshal(b []byte) (time.Time, int, error) {
	return wire.ConsumeTimestamp(b)
}
func (c *TimestampValueCodec) WriteJson(e *kdsjson.Encoder, v time.Time) error {
	kdsjson.WriteTimestampValue(e, v)
	return nil
}

type DurationValueCodec struct{}

func (c *DurationValueCodec) WireType() wire.Type            { return wire.BytesType }
func (c *DurationValueCodec) Compare(a, b time.Duration) int { return cmp.Compare(a, b) }
func (c *DurationValueCodec) Marshal(b []byte, v time.Duration) []byte {
	return wire.AppendDuration(b, v)
}
func (c *DurationValueCodec) Unmarshal(b []byte) (time.Duration, int, error) {
	return wire.ConsumeDuration(b)
}
func (c *DurationValueCodec) WriteJson(e *kdsjson.Encoder, v time.Duration) error {
	kdsjson.WriteDurationValue(e, v)
	return nil
}

type EmptyValueCodec struct{}

func (c *EmptyValueCodec) WireType() wire.Type                       { return wire.BytesType }
func (c *EmptyValueCodec) Compare(a, b struct{}) int                 { return 0 }
func (c *EmptyValueCodec) Marshal(b []byte, v struct{}) []byte       { return wire.AppendEmpty(b, v) }
func (c *EmptyValueCodec) Unmarshal(b []byte) (struct{}, int, error) { return wire.ConsumeEmpty(b) }
func (c *EmptyValueCodec) WriteJson(e *kdsjson.Encoder, v struct{}) error {
	kdsjson.WriteEmptyValue(e, v)
	return nil
}

type EnumValueCodec[T ~int32] struct{}

func (c *EnumValueCodec[T]) WireType() wire.Type                { return wire.VarintType }
func (c *EnumValueCodec[T]) Compare(a, b T) int                 { return cmp.Compare(a, b) }
func (c *EnumValueCodec[T]) Marshal(b []byte, v T) []byte       { return wire.AppendEnum(b, v) }
func (c *EnumValueCodec[T]) Unmarshal(b []byte) (T, int, error) { return wire.ConsumeEnum[T](b) }
func (c *EnumValueCodec[T]) WriteJson(e *kdsjson.Encoder, v T) error {
	kdsjson.WriteEnumValue(e, v)
	return nil
}
