package wire

import (
	"math"
	"time"

	"google.golang.org/protobuf/encoding/protowire"
)

type Type = protowire.Type

const (
	VarintType  Type = protowire.VarintType
	Fixed32Type Type = protowire.Fixed32Type
	Fixed64Type Type = protowire.Fixed64Type
	BytesType   Type = protowire.BytesType
)

type Number = protowire.Number

func AppendTag(b []byte, num Number, typ Type) []byte {
	return protowire.AppendTag(b, num, typ)
}

func AppendBool(b []byte, v bool) []byte {
	return protowire.AppendVarint(b, protowire.EncodeBool(v))
}

func AppendInt32(b []byte, v int32) []byte {
	return protowire.AppendVarint(b, uint64(v))
}

func AppendSint32(b []byte, v int32) []byte {
	return protowire.AppendVarint(b, protowire.EncodeZigZag(int64(v)))
}

func AppendUint32(b []byte, v uint32) []byte {
	return protowire.AppendVarint(b, uint64(v))
}

func AppendInt64(b []byte, v int64) []byte {
	return protowire.AppendVarint(b, uint64(v))
}

func AppendSint64(b []byte, v int64) []byte {
	return protowire.AppendVarint(b, protowire.EncodeZigZag(v))
}

func AppendUint64(b []byte, v uint64) []byte {
	return protowire.AppendVarint(b, v)
}

func AppendSfixed32(b []byte, v int32) []byte {
	return protowire.AppendFixed32(b, uint32(v))
}

func AppendFixed32(b []byte, v uint32) []byte {
	return protowire.AppendFixed32(b, v)
}

func AppendFloat(b []byte, v float32) []byte {
	return protowire.AppendFixed32(b, math.Float32bits(v))
}

func AppendSfixed64(b []byte, v int64) []byte {
	return protowire.AppendFixed64(b, uint64(v))
}

func AppendFixed64(b []byte, v uint64) []byte {
	return protowire.AppendFixed64(b, v)
}

func AppendDouble(b []byte, v float64) []byte {
	return protowire.AppendFixed64(b, math.Float64bits(v))
}

func AppendString(b []byte, v string) []byte {
	return protowire.AppendString(b, v)
}

func AppendBytes(b []byte, v []byte) []byte {
	return protowire.AppendBytes(b, v)
}

type Message interface {
	MarshalMessage(b []byte) ([]byte, error)
}

type List interface {
	MarshalList(b []byte) ([]byte, error)
}

type Map interface {
	MarshalMap(b []byte) ([]byte, error)
}

func AppendMessage(b []byte, v Message) ([]byte, error) {
	var pos int
	var err error
	b, pos = appendSpeculativeLength(b)
	b, err = v.MarshalMessage(b)
	if err != nil {
		return b, err
	}
	b = finishSpeculativeLength(b, pos)
	return b, nil
}

func AppendList(b []byte, v List) ([]byte, error) {
	var pos int
	var err error
	b, pos = appendSpeculativeLength(b)
	b, err = v.MarshalList(b)
	if err != nil {
		return b, err
	}
	b = finishSpeculativeLength(b, pos)
	return b, nil
}

func AppendMap(b []byte, v Map) ([]byte, error) {
	var pos int
	var err error
	b, pos = appendSpeculativeLength(b)
	b, err = v.MarshalMap(b)
	if err != nil {
		return b, err
	}
	b = finishSpeculativeLength(b, pos)
	return b, nil
}

func AppendTimestamp(b []byte, v time.Time) []byte {
	var pos int
	b, pos = appendSpeculativeLength(b)
	b = protowire.AppendTag(b, 1, protowire.VarintType)
	b = protowire.AppendVarint(b, uint64(v.Unix()))
	b = protowire.AppendTag(b, 2, protowire.VarintType)
	b = protowire.AppendVarint(b, uint64(v.Nanosecond()))
	return finishSpeculativeLength(b, pos)
}
func AppendDuration(b []byte, v time.Duration) []byte {
	nanos := v.Nanoseconds()
	secs := nanos / 1e9
	nanos -= secs * 1e9
	var pos int
	b, pos = appendSpeculativeLength(b)
	b = protowire.AppendTag(b, 1, protowire.VarintType)
	b = protowire.AppendVarint(b, uint64(nanos))
	b = protowire.AppendTag(b, 2, protowire.VarintType)
	b = protowire.AppendVarint(b, uint64(secs))
	return finishSpeculativeLength(b, pos)
}

func AppendEmpty(b []byte) []byte {
	return protowire.AppendBytes(b, nil)
}

// When encoding length-prefixed fields, we speculatively set aside some number of bytes
// for the length, encode the data, and then encode the length (shifting the data if necessary
// to make room).
const speculativeLength = 1

func appendSpeculativeLength(b []byte) ([]byte, int) {
	pos := len(b)
	b = append(b, "\x00\x00\x00\x00"[:speculativeLength]...)
	return b, pos
}

func finishSpeculativeLength(b []byte, pos int) []byte {
	mlen := len(b) - pos - speculativeLength
	msiz := protowire.SizeVarint(uint64(mlen))
	if msiz != speculativeLength {
		for i := 0; i < msiz-speculativeLength; i++ {
			b = append(b, 0)
		}
		copy(b[pos+msiz:], b[pos+speculativeLength:])
		b = b[:pos+msiz+mlen]
	}
	protowire.AppendVarint(b[:pos], uint64(mlen))
	return b
}
