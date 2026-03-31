package kdsync

import (
	"bytes"
	"cmp"
	"strings"
	"time"

	"github.com/iakud/knoll/kdsync/wire"
)

type FieldCodec[T any] struct {
	wireType      wire.Type
	compareFunc   func(a, b T) int
	marshalFunc   func(b []byte, v T) []byte
	unmarshalFunc func(b []byte) (T, int, error)
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

func BoolCodec() FieldCodec[bool] {
	return FieldCodec[bool]{
		wireType:      wire.VarintType,
		compareFunc:   boolCompare,
		marshalFunc:   wire.AppendBool,
		unmarshalFunc: wire.ConsumeBool,
	}
}

func Int32Codec() FieldCodec[int32] {
	return FieldCodec[int32]{
		wireType:      wire.VarintType,
		compareFunc:   cmp.Compare[int32],
		marshalFunc:   wire.AppendInt32,
		unmarshalFunc: wire.ConsumeInt32,
	}
}

func Uint32Codec() FieldCodec[uint32] {
	return FieldCodec[uint32]{
		wireType:      wire.VarintType,
		compareFunc:   cmp.Compare[uint32],
		marshalFunc:   wire.AppendUint32,
		unmarshalFunc: wire.ConsumeUint32,
	}
}

func Int64Codec() FieldCodec[int64] {
	return FieldCodec[int64]{
		wireType:      wire.VarintType,
		compareFunc:   cmp.Compare[int64],
		marshalFunc:   wire.AppendInt64,
		unmarshalFunc: wire.ConsumeInt64,
	}
}

func Uint64Codec() FieldCodec[uint64] {
	return FieldCodec[uint64]{
		wireType:      wire.VarintType,
		compareFunc:   cmp.Compare[uint64],
		marshalFunc:   wire.AppendUint64,
		unmarshalFunc: wire.ConsumeUint64,
	}
}

func Float32Codec() FieldCodec[float32] {
	return FieldCodec[float32]{
		wireType:      wire.Fixed32Type,
		compareFunc:   cmp.Compare[float32],
		marshalFunc:   wire.AppendFloat,
		unmarshalFunc: wire.ConsumeFloat,
	}
}

func Float64Codec() FieldCodec[float64] {
	return FieldCodec[float64]{
		wireType:      wire.Fixed64Type,
		compareFunc:   cmp.Compare[float64],
		marshalFunc:   wire.AppendDouble,
		unmarshalFunc: wire.ConsumeDouble,
	}
}

func StringCodec() FieldCodec[string] {
	return FieldCodec[string]{
		wireType:      wire.BytesType,
		compareFunc:   strings.Compare,
		marshalFunc:   wire.AppendString,
		unmarshalFunc: wire.ConsumeString,
	}
}

func BytesCodec() FieldCodec[[]byte] {
	return FieldCodec[[]byte]{
		wireType:      wire.BytesType,
		compareFunc:   bytes.Compare,
		marshalFunc:   wire.AppendBytes,
		unmarshalFunc: wire.ConsumeBytes,
	}
}

func TimestampCodec() FieldCodec[time.Time] {
	return FieldCodec[time.Time]{
		wireType:      wire.BytesType,
		compareFunc:   timestampCompare,
		marshalFunc:   wire.AppendTimestamp,
		unmarshalFunc: wire.ConsumeTimestamp,
	}
}

func DurationCodec() FieldCodec[time.Duration] {
	return FieldCodec[time.Duration]{
		wireType:      wire.BytesType,
		compareFunc:   cmp.Compare[time.Duration],
		marshalFunc:   wire.AppendDuration,
		unmarshalFunc: wire.ConsumeDuration,
	}
}

func EmptyCodec() FieldCodec[struct{}] {
	return FieldCodec[struct{}]{
		wireType:      wire.BytesType,
		compareFunc:   emptyCompare,
		marshalFunc:   wire.AppendEmpty,
		unmarshalFunc: wire.ConsumeEmpty,
	}
}

func EnumCodec[T ~int32]() FieldCodec[T] {
	return FieldCodec[T]{
		wireType:      wire.VarintType,
		compareFunc:   cmp.Compare[T],
		marshalFunc:   wire.AppendEnum[T],
		unmarshalFunc: wire.ConsumeEnum[T],
	}
}
