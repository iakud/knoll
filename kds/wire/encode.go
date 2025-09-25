package wire

import (
	"math"
	"time"

	"google.golang.org/protobuf/encoding/protowire"
)

func MarshalMessage(b []byte, num Number, m Marshaler) ([]byte, error) {
	b = protowire.AppendTag(b, num, protowire.BytesType)
	return AppendMessage(b, m)
}

func MarshalMessageDirty(b []byte, num Number, m Marshaler) ([]byte, error) {
	b = protowire.AppendTag(b, num, protowire.BytesType)
	return AppendMessageDirty(b, m)
}

func MarshalMessageFunc(b []byte, num Number, f MarshalFunc) ([]byte, error) {
	b = protowire.AppendTag(b, num, protowire.BytesType)
	return AppendMessageFunc(b, f)
}

func MarshalTimestamp(b []byte, num Number, v time.Time) ([]byte, error) {
	b = protowire.AppendTag(b, num, protowire.BytesType)
	return AppendTimestamp(b, v), nil
}
func MarshalDuration(b []byte, num Number, v time.Duration) ([]byte, error) {
	b = protowire.AppendTag(b, num, protowire.BytesType)
	return AppendDuration(b, v), nil
}

func MarshalEmpty(b []byte, num Number, v struct{}) ([]byte, error) {
	b = protowire.AppendTag(b, num, protowire.BytesType)
	return AppendEmpty(b, v), nil
}

func MarshalBool(b []byte, num Number, v bool) ([]byte, error) {
	b = protowire.AppendTag(b, num, protowire.VarintType)
	b = protowire.AppendVarint(b, protowire.EncodeBool(v))
	return b, nil
}

func MarshalInt32(b []byte, num Number, v int32) ([]byte, error) {
	b = protowire.AppendTag(b, num, protowire.VarintType)
	b = protowire.AppendVarint(b, uint64(v))
	return b, nil
}

func MarshalSint32(b []byte, num Number, v int32) ([]byte, error) {
	b = protowire.AppendTag(b, num, protowire.VarintType)
	b = protowire.AppendVarint(b, protowire.EncodeZigZag(int64(v)))
	return b, nil
}

func MarshalUint32(b []byte, num Number, v uint32) ([]byte, error) {
	b = protowire.AppendTag(b, num, protowire.VarintType)
	b = protowire.AppendVarint(b, uint64(v))
	return b, nil
}

func MarshalInt64(b []byte, num Number, v int64) ([]byte, error) {
	b = protowire.AppendTag(b, num, protowire.VarintType)
	b = protowire.AppendVarint(b, uint64(v))
	return b, nil
}

func MarshalSint64(b []byte, num Number, v int64) ([]byte, error) {
	b = protowire.AppendTag(b, num, protowire.VarintType)
	b = protowire.AppendVarint(b, protowire.EncodeZigZag(v))
	return b, nil
}

func MarshalUint64(b []byte, num Number, v uint64) ([]byte, error) {
	b = protowire.AppendTag(b, num, protowire.VarintType)
	b = protowire.AppendVarint(b, v)
	return b, nil
}

func MarshalSfixed32(b []byte, num Number, v int32) ([]byte, error) {
	b = protowire.AppendTag(b, num, protowire.Fixed32Type)
	b = protowire.AppendFixed32(b, uint32(v))
	return b, nil
}

func MarshalFixed32(b []byte, num Number, v uint32) ([]byte, error) {
	b = protowire.AppendTag(b, num, protowire.Fixed32Type)
	b = protowire.AppendFixed32(b, v)
	return b, nil
}

func MarshalFloat(b []byte, num Number, v float32) ([]byte, error) {
	b = protowire.AppendTag(b, num, protowire.Fixed32Type)
	b = protowire.AppendFixed32(b, math.Float32bits(v))
	return b, nil
}

func MarshalSfixed64(b []byte, num Number, v int64) ([]byte, error) {
	b = protowire.AppendTag(b, num, protowire.Fixed64Type)
	b = protowire.AppendFixed64(b, uint64(v))
	return b, nil
}

func MarshalFixed64(b []byte, num Number, v uint64) ([]byte, error) {
	b = protowire.AppendTag(b, num, protowire.Fixed64Type)
	b = protowire.AppendFixed64(b, v)
	return b, nil
}

func MarshalDouble(b []byte, num Number, v float64) ([]byte, error) {
	b = protowire.AppendTag(b, num, protowire.Fixed64Type)
	b = protowire.AppendFixed64(b, math.Float64bits(v))
	return b, nil
}

func MarshalString(b []byte, num Number, v string) ([]byte, error) {
	b = protowire.AppendTag(b, num, protowire.BytesType)
	b = protowire.AppendString(b, v)
	return b, nil
}

func MarshalBytes(b []byte, num Number, v []byte) ([]byte, error) {
	b = protowire.AppendTag(b, num, protowire.BytesType)
	b = protowire.AppendBytes(b, v)
	return b, nil
}
