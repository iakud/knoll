package kdsync

import (
	"bytes"
	"cmp"
	"strings"
	"time"

	"github.com/iakud/knoll/kdsync/wire"
)

type FieldCodec[T any] struct {
	WireType      wire.Type
	CompareFunc   func(a, b T) int
	MarshalFunc   func(b []byte, v T) []byte
	UnmarshalFunc func(b []byte) (T, int, error)
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

var BoolCodec = FieldCodec[bool]{
	WireType:      wire.VarintType,
	CompareFunc:   boolCompare,
	MarshalFunc:   wire.AppendBool,
	UnmarshalFunc: wire.ConsumeBool,
}

var Int32Codec = FieldCodec[int32]{
	WireType:      wire.VarintType,
	CompareFunc:   cmp.Compare[int32],
	MarshalFunc:   wire.AppendInt32,
	UnmarshalFunc: wire.ConsumeInt32,
}

var Uint32Codec = FieldCodec[uint32]{
	WireType:      wire.VarintType,
	CompareFunc:   cmp.Compare[uint32],
	MarshalFunc:   wire.AppendUint32,
	UnmarshalFunc: wire.ConsumeUint32,
}

var Int64Codec = FieldCodec[int64]{
	WireType:      wire.VarintType,
	CompareFunc:   cmp.Compare[int64],
	MarshalFunc:   wire.AppendInt64,
	UnmarshalFunc: wire.ConsumeInt64,
}

var Uint64Codec = FieldCodec[uint64]{
	WireType:      wire.VarintType,
	CompareFunc:   cmp.Compare[uint64],
	MarshalFunc:   wire.AppendUint64,
	UnmarshalFunc: wire.ConsumeUint64,
}

var Float32Codec = FieldCodec[float32]{
	WireType:      wire.Fixed32Type,
	CompareFunc:   cmp.Compare[float32],
	MarshalFunc:   wire.AppendFloat,
	UnmarshalFunc: wire.ConsumeFloat,
}

var Float64Codec = FieldCodec[float64]{
	WireType:      wire.Fixed64Type,
	CompareFunc:   cmp.Compare[float64],
	MarshalFunc:   wire.AppendDouble,
	UnmarshalFunc: wire.ConsumeDouble,
}

var StringCodec = FieldCodec[string]{
	WireType:      wire.BytesType,
	CompareFunc:   strings.Compare,
	MarshalFunc:   wire.AppendString,
	UnmarshalFunc: wire.ConsumeString,
}

var BytesCodec = FieldCodec[[]byte]{
	WireType:      wire.BytesType,
	CompareFunc:   bytes.Compare,
	MarshalFunc:   wire.AppendBytes,
	UnmarshalFunc: wire.ConsumeBytes,
}

var TimestampCodec = FieldCodec[time.Time]{
	WireType:      wire.BytesType,
	CompareFunc:   timestampCompare,
	MarshalFunc:   wire.AppendTimestamp,
	UnmarshalFunc: wire.ConsumeTimestamp,
}

var DurationCodec = FieldCodec[time.Duration]{
	WireType:      wire.BytesType,
	CompareFunc:   cmp.Compare[time.Duration],
	MarshalFunc:   wire.AppendDuration,
	UnmarshalFunc: wire.ConsumeDuration,
}

var EmptyCodec = FieldCodec[struct{}]{
	WireType:      wire.BytesType,
	CompareFunc:   emptyCompare,
	MarshalFunc:   wire.AppendEmpty,
	UnmarshalFunc: wire.ConsumeEmpty,
}

func NewEnumCodec[T ~int32]() FieldCodec[T] {
	return FieldCodec[T]{
		WireType:      wire.VarintType,
		CompareFunc:   cmp.Compare[T],
		MarshalFunc:   wire.AppendEnum[T],
		UnmarshalFunc: wire.ConsumeEnum[T],
	}
}
