package kdsync

import (
	"bytes"
	"cmp"
	"strings"
	"time"

	"github.com/iakud/knoll/kdsync/kdsjson"
	"github.com/iakud/knoll/kdsync/wire"
)

type fieldCodec[T any] struct {
	wireType      wire.Type
	compareFunc   func(a, b T) int
	marshalFunc   func(b []byte, v T) []byte
	unmarshalFunc func(b []byte) (T, int, error)
	writeJSONFunc func(e *kdsjson.Encoder, v T)
}

func boolCompare(a, b bool) int {
	if a {
		return 1
	} else if b {
		return -1
	}
	return 0
}

func timestampCompare(a, b time.Time) int {
	return a.Compare(b)
}

func emptyCompare(a, b struct{}) int {
	return 0
}

// Key Codec

var BoolKeyCodec = fieldCodec[bool]{
	wireType:      wire.VarintType,
	compareFunc:   boolCompare,
	marshalFunc:   wire.AppendBool,
	unmarshalFunc: wire.ConsumeBool,
	writeJSONFunc: kdsjson.WriteBoolPropertyName,
}

var Int32KeyCodec = fieldCodec[int32]{
	wireType:      wire.VarintType,
	compareFunc:   cmp.Compare[int32],
	marshalFunc:   wire.AppendInt32,
	unmarshalFunc: wire.ConsumeInt32,
	writeJSONFunc: kdsjson.WriteInt32PropertyName,
}

var Uint32KeyCodec = fieldCodec[uint32]{
	wireType:      wire.VarintType,
	compareFunc:   cmp.Compare[uint32],
	marshalFunc:   wire.AppendUint32,
	unmarshalFunc: wire.ConsumeUint32,
	writeJSONFunc: kdsjson.WriteUint32PropertyName,
}

var Int64KeyCodec = fieldCodec[int64]{
	wireType:      wire.VarintType,
	compareFunc:   cmp.Compare[int64],
	marshalFunc:   wire.AppendInt64,
	unmarshalFunc: wire.ConsumeInt64,
	writeJSONFunc: kdsjson.WriteInt64PropertyName,
}

var Uint64KeyCodec = fieldCodec[uint64]{
	wireType:      wire.VarintType,
	compareFunc:   cmp.Compare[uint64],
	marshalFunc:   wire.AppendUint64,
	unmarshalFunc: wire.ConsumeUint64,
	writeJSONFunc: kdsjson.WriteUint64PropertyName,
}

var StringKeyCodec = fieldCodec[string]{
	wireType:      wire.BytesType,
	compareFunc:   strings.Compare,
	marshalFunc:   wire.AppendString,
	unmarshalFunc: wire.ConsumeString,
	writeJSONFunc: kdsjson.WritePropertyName,
}

// Value Codec

var BoolValueCodec = fieldCodec[bool]{
	wireType:      wire.VarintType,
	compareFunc:   boolCompare,
	marshalFunc:   wire.AppendBool,
	unmarshalFunc: wire.ConsumeBool,
	writeJSONFunc: kdsjson.WriteBoolValue,
}

var Int32ValueCodec = fieldCodec[int32]{
	wireType:      wire.VarintType,
	compareFunc:   cmp.Compare[int32],
	marshalFunc:   wire.AppendInt32,
	unmarshalFunc: wire.ConsumeInt32,
	writeJSONFunc: kdsjson.WriteInt32Value,
}

var Uint32ValueCodec = fieldCodec[uint32]{
	wireType:      wire.VarintType,
	compareFunc:   cmp.Compare[uint32],
	marshalFunc:   wire.AppendUint32,
	unmarshalFunc: wire.ConsumeUint32,
	writeJSONFunc: kdsjson.WriteUint32Value,
}

var Int64ValueCodec = fieldCodec[int64]{
	wireType:      wire.VarintType,
	compareFunc:   cmp.Compare[int64],
	marshalFunc:   wire.AppendInt64,
	unmarshalFunc: wire.ConsumeInt64,
	writeJSONFunc: kdsjson.WriteInt64Value,
}

var Uint64ValueCodec = fieldCodec[uint64]{
	wireType:      wire.VarintType,
	compareFunc:   cmp.Compare[uint64],
	marshalFunc:   wire.AppendUint64,
	unmarshalFunc: wire.ConsumeUint64,
	writeJSONFunc: kdsjson.WriteUint64Value,
}

var Float32ValueCodec = fieldCodec[float32]{
	wireType:      wire.Fixed32Type,
	compareFunc:   cmp.Compare[float32],
	marshalFunc:   wire.AppendFloat,
	unmarshalFunc: wire.ConsumeFloat,
	writeJSONFunc: kdsjson.WriteFloat32Value,
}

var Float64ValueCodec = fieldCodec[float64]{
	wireType:      wire.Fixed64Type,
	compareFunc:   cmp.Compare[float64],
	marshalFunc:   wire.AppendDouble,
	unmarshalFunc: wire.ConsumeDouble,
	writeJSONFunc: kdsjson.WriteFloat64Value,
}

var StringValueCodec = fieldCodec[string]{
	wireType:      wire.BytesType,
	compareFunc:   strings.Compare,
	marshalFunc:   wire.AppendString,
	unmarshalFunc: wire.ConsumeString,
	writeJSONFunc: kdsjson.WriteStringValue,
}

var BytesValueCodec = fieldCodec[[]byte]{
	wireType:      wire.BytesType,
	compareFunc:   bytes.Compare,
	marshalFunc:   wire.AppendBytes,
	unmarshalFunc: wire.ConsumeBytes,
	writeJSONFunc: kdsjson.WriteBytesValue,
}

var TimestampValueCodec = fieldCodec[time.Time]{
	wireType:      wire.BytesType,
	compareFunc:   timestampCompare,
	marshalFunc:   wire.AppendTimestamp,
	unmarshalFunc: wire.ConsumeTimestamp,
	writeJSONFunc: kdsjson.WriteTimestampValue,
}

var DurationValueCodec = fieldCodec[time.Duration]{
	wireType:      wire.BytesType,
	compareFunc:   cmp.Compare[time.Duration],
	marshalFunc:   wire.AppendDuration,
	unmarshalFunc: wire.ConsumeDuration,
	writeJSONFunc: kdsjson.WriteDurationValue,
}

var EmptyValueCodec = fieldCodec[struct{}]{
	wireType:      wire.BytesType,
	compareFunc:   emptyCompare,
	marshalFunc:   wire.AppendEmpty,
	unmarshalFunc: wire.ConsumeEmpty,
	writeJSONFunc: kdsjson.WriteEmptyValue,
}

func EnumValueCodec[T ~int32]() fieldCodec[T] {
	return fieldCodec[T]{
		wireType:      wire.VarintType,
		compareFunc:   cmp.Compare[T],
		marshalFunc:   wire.AppendEnum[T],
		unmarshalFunc: wire.ConsumeEnum[T],
		writeJSONFunc: kdsjson.WriteEnumValue[T],
	}
}
