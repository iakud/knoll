package kdsync

import (
	"time"

	"github.com/iakud/knoll/kdsync/wire"
)

type FieldCodec[T Field] struct {
	MarshalFunc   func(b []byte, v T) []byte
	UnmarshalFunc func(b []byte) (T, int, error)
}

var BoolCodec = FieldCodec[bool]{
	MarshalFunc:   wire.AppendBool,
	UnmarshalFunc: wire.ConsumeBool,
}

var Int32Codec = FieldCodec[int32]{
	MarshalFunc:   wire.AppendInt32,
	UnmarshalFunc: wire.ConsumeInt32,
}

var Uint32Codec = FieldCodec[uint32]{
	MarshalFunc:   wire.AppendUint32,
	UnmarshalFunc: wire.ConsumeUint32,
}

var Int64Codec = FieldCodec[int64]{
	MarshalFunc:   wire.AppendInt64,
	UnmarshalFunc: wire.ConsumeInt64,
}

var Uint64Codec = FieldCodec[uint64]{
	MarshalFunc:   wire.AppendUint64,
	UnmarshalFunc: wire.ConsumeUint64,
}

var Float32Codec = FieldCodec[float32]{
	MarshalFunc:   wire.AppendFloat,
	UnmarshalFunc: wire.ConsumeFloat,
}

var Float64Codec = FieldCodec[float64]{
	MarshalFunc:   wire.AppendDouble,
	UnmarshalFunc: wire.ConsumeDouble,
}

var StringCodec = FieldCodec[string]{
	MarshalFunc:   wire.AppendString,
	UnmarshalFunc: wire.ConsumeString,
}

var DurationCodec = FieldCodec[time.Duration]{
	MarshalFunc:   wire.AppendDuration,
	UnmarshalFunc: wire.ConsumeDuration,
}

var EmptyCodec = FieldCodec[struct{}]{
	MarshalFunc:   wire.AppendEmpty,
	UnmarshalFunc: wire.ConsumeEmpty,
}
