package wire

import (
	"errors"
	"math"
	"time"

	"google.golang.org/protobuf/encoding/protowire"
)

type Number = protowire.Number

const (
	MaxValidNumber Number = 1<<29 - 1
)

type Type = protowire.Type

const (
	VarintType  Type = protowire.VarintType
	Fixed32Type Type = protowire.Fixed32Type
	Fixed64Type Type = protowire.Fixed64Type
	BytesType   Type = protowire.BytesType
)

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
	MarshalMessageDirty(b []byte) ([]byte, error)
	UnmarshalMessage(b []byte) error
}

type List interface {
	MarshalList(b []byte) ([]byte, error)
	UnmarshalList(b []byte) error
}

type Map interface {
	MarshalMap(b []byte) ([]byte, error)
	UnmarshalMap(b []byte) error
}

func AppendMessage(b []byte, v Message) ([]byte, error) {
	var pos int
	var err error
	b, pos = AppendSpeculativeLength(b)
	b, err = v.MarshalMessage(b)
	if err != nil {
		return b, err
	}
	b = FinishSpeculativeLength(b, pos)
	return b, nil
}

func AppendMessageDirty(b []byte, v Message) ([]byte, error) {
	var pos int
	var err error
	b, pos = AppendSpeculativeLength(b)
	b, err = v.MarshalMessageDirty(b)
	if err != nil {
		return b, err
	}
	b = FinishSpeculativeLength(b, pos)
	return b, nil
}

func AppendList(b []byte, v List) ([]byte, error) {
	var pos int
	var err error
	b, pos = AppendSpeculativeLength(b)
	b, err = v.MarshalList(b)
	if err != nil {
		return b, err
	}
	b = FinishSpeculativeLength(b, pos)
	return b, nil
}

func AppendMap(b []byte, v Map) ([]byte, error) {
	var pos int
	var err error
	b, pos = AppendSpeculativeLength(b)
	b, err = v.MarshalMap(b)
	if err != nil {
		return b, err
	}
	b = FinishSpeculativeLength(b, pos)
	return b, nil
}

func AppendTimestamp(b []byte, v time.Time) []byte {
	var pos int
	b, pos = AppendSpeculativeLength(b)
	b = protowire.AppendTag(b, 1, protowire.VarintType)
	b = protowire.AppendVarint(b, uint64(v.Unix()))
	b = protowire.AppendTag(b, 2, protowire.VarintType)
	b = protowire.AppendVarint(b, uint64(int32(v.Nanosecond())))
	return FinishSpeculativeLength(b, pos)
}
func AppendDuration(b []byte, v time.Duration) []byte {
	nanos := v.Nanoseconds()
	secs := nanos / 1e9
	nanos -= secs * 1e9
	var pos int
	b, pos = AppendSpeculativeLength(b)
	b = protowire.AppendTag(b, 1, protowire.VarintType)
	b = protowire.AppendVarint(b, uint64(secs))
	b = protowire.AppendTag(b, 2, protowire.VarintType)
	b = protowire.AppendVarint(b, uint64(nanos))
	return FinishSpeculativeLength(b, pos)
}

func AppendEmpty(b []byte) []byte {
	return protowire.AppendBytes(b, nil)
}

// When encoding length-prefixed fields, we speculatively set aside some number of bytes
// for the length, encode the data, and then encode the length (shifting the data if necessary
// to make room).
const speculativeLength = 1

func AppendSpeculativeLength(b []byte) ([]byte, int) {
	pos := len(b)
	b = append(b, "\x00\x00\x00\x00"[:speculativeLength]...)
	return b, pos
}

func FinishSpeculativeLength(b []byte, pos int) []byte {
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

// decode
func ConsumeTag(b []byte) (Number, Type, int) {
	return protowire.ConsumeTag(b)
}

func ConsumeBytes(b []byte) ([]byte, int) {
	return protowire.ConsumeBytes(b)
}

func ConsumeFieldValue(num Number, wtyp Type, b []byte) int {
	return protowire.ConsumeFieldValue(num, wtyp, b)
}

func UnmarshalList(b []byte, wtyp Type, v List) (int, error) {
	if wtyp != BytesType {
		return 0, ErrUnknown
	}
	var n int
	b, n = protowire.ConsumeBytes(b)
	if n < 0 {
		return 0, ErrDecode
	}
	if err := v.UnmarshalList(b); err != nil {
		return 0, err
	}
	return n, nil
}

func UnmarshalMap(b []byte, wtyp Type, v Map) (int, error) {
	if wtyp != BytesType {
		return 0, ErrUnknown
	}
	var n int
	b, n = protowire.ConsumeBytes(b)
	if n < 0 {
		return 0, ErrDecode
	}
	if err := v.UnmarshalMap(b); err != nil {
		return 0, err
	}
	return n, nil
}

func UnmarshalMessage(b []byte, wtyp Type, v Message) (int, error) {
	if wtyp != BytesType {
		return 0, ErrUnknown
	}
	var n int
	b, n = protowire.ConsumeBytes(b)
	if n < 0 {
		return 0, ErrDecode
	}
	if err := v.UnmarshalMessage(b); err != nil {
		return 0, err
	}
	return n, nil
}

func UnmarshalBool(b []byte, wtyp Type, m *bool) (int, error) {
	if wtyp != VarintType {
		return 0, ErrUnknown
	}
	v, n := protowire.ConsumeVarint(b)
	if n < 0 {
		return 0, ErrDecode
	}
	*m = protowire.DecodeBool(v)
	return n, nil
}

func UnmarshalInt32(b []byte, wtyp Type, m *int32) (int, error) {
	if wtyp != VarintType {
		return 0, ErrUnknown
	}
	v, n := protowire.ConsumeVarint(b)
	if n < 0 {
		return 0, ErrDecode
	}
	*m = int32(v)
	return n, nil
}

func UnmarshalSint32(b []byte, wtyp Type, m *int32) (int, error) {
	if wtyp != VarintType {
		return 0, ErrUnknown
	}
	v, n := protowire.ConsumeVarint(b)
	if n < 0 {
		return 0, ErrDecode
	}
	*m = int32(protowire.DecodeZigZag(v & math.MaxUint32))
	return n, nil
}

func UnmarshalUint32(b []byte, wtyp Type, m *uint32) (int, error) {
	if wtyp != VarintType {
		return 0, ErrUnknown
	}
	v, n := protowire.ConsumeVarint(b)
	if n < 0 {
		return 0, ErrDecode
	}
	*m = uint32(v)
	return n, nil
}

func UnmarshalInt64(b []byte, wtyp Type, m *int64) (int, error) {
	if wtyp != VarintType {
		return 0, ErrUnknown
	}
	v, n := protowire.ConsumeVarint(b)
	if n < 0 {
		return 0, ErrDecode
	}
	*m = int64(v)
	return n, nil
}

func UnmarshalSint64(b []byte, wtyp Type, m *int64) (int, error) {
	if wtyp != VarintType {
		return 0, ErrUnknown
	}
	v, n := protowire.ConsumeVarint(b)
	if n < 0 {
		return 0, ErrDecode
	}
	*m = protowire.DecodeZigZag(v & math.MaxUint64)
	return n, nil
}

func UnmarshalUint64(b []byte, wtyp Type, m *uint64) (int, error) {
	if wtyp != VarintType {
		return 0, ErrUnknown
	}
	v, n := protowire.ConsumeVarint(b)
	if n < 0 {
		return 0, ErrDecode
	}
	*m = v
	return n, nil
}

func UnmarshalSfixed32(b []byte, wtyp Type, m *int32) (int, error) {
	if wtyp != Fixed32Type {
		return 0, ErrUnknown
	}
	v, n := protowire.ConsumeFixed32(b)
	if n < 0 {
		return 0, ErrDecode
	}
	*m = int32(v)
	return n, nil
}

func UnmarshalFixed32(b []byte, wtyp Type, m *uint32) (int, error) {
	if wtyp != Fixed32Type {
		return 0, ErrUnknown
	}
	v, n := protowire.ConsumeFixed32(b)
	if n < 0 {
		return 0, ErrDecode
	}
	*m = v
	return n, nil
}

func UnmarshalFloat(b []byte, wtyp Type, m *float32) (int, error) {
	if wtyp != Fixed32Type {
		return 0, ErrUnknown
	}
	v, n := protowire.ConsumeFixed32(b)
	if n < 0 {
		return 0, ErrDecode
	}
	*m = math.Float32frombits(v)
	return n, nil
}

func UnmarshalSfixed64(b []byte, wtyp Type, m *int64) (int, error) {
	if wtyp != Fixed64Type {
		return 0, ErrUnknown
	}
	v, n := protowire.ConsumeFixed64(b)
	if n < 0 {
		return 0, ErrDecode
	}
	*m = int64(v)
	return n, nil
}

func UnmarshalFixed64(b []byte, wtyp Type, m *uint64) (int, error) {
	if wtyp != Fixed64Type {
		return 0, ErrUnknown
	}
	v, n := protowire.ConsumeFixed64(b)
	if n < 0 {
		return 0, ErrDecode
	}
	*m = v
	return n, nil
}

func UnmarshalDouble(b []byte, wtyp Type, m *float64) (int, error) {
	if wtyp != Fixed64Type {
		return 0, ErrUnknown
	}
	v, n := protowire.ConsumeFixed64(b)
	if n < 0 {
		return 0, ErrDecode
	}
	*m = math.Float64frombits(v)
	return n, nil
}

func UnmarshalString(b []byte, wtyp Type, m *string) (int, error) {
	if wtyp != BytesType {
		return 0, ErrUnknown
	}
	v, n := protowire.ConsumeString(b)
	if n < 0 {
		return 0, ErrDecode
	}
	*m = v
	return n, nil
}

func UnmarshalBytes(b []byte, wtyp Type, m *[]byte) (int, error) {
	if wtyp != BytesType {
		return 0, ErrUnknown
	}
	v, n := protowire.ConsumeBytes(b)
	if n < 0 {
		return 0, ErrDecode
	}
	*m = v
	return n, nil
}

func UnmarshalTimestamp(b []byte, wtyp Type, m *time.Time) (int, error) {
	if wtyp != BytesType {
		return 0, ErrUnknown
	}
	v, n := protowire.ConsumeBytes(b)
	if n < 0 {
		return 0, ErrDecode
	}
	var sec, nsec int64
	for len(v) > 0 {
		num, wtyp, n := protowire.ConsumeTag(v)
		if n < 0 {
			return 0, ErrDecode
		}
		if num > protowire.MaxValidNumber {
			return 0, ErrDecode
		}
		v = v[n:]
		var err error = ErrUnknown
		switch num {
		case 1:
			n, err = UnmarshalInt64(v, wtyp, &sec)
		case 2:
			n, err = UnmarshalInt64(v, wtyp, &nsec)
		}
		if err == ErrUnknown {
			n = protowire.ConsumeFieldValue(num, wtyp, v)
			if n < 0 {
				return 0, ErrDecode
			}
		} else if err != nil {
			return 0, err
		}
		v = v[n:]
	}
	*m = time.Unix(sec, nsec)
	return n, nil
}

func UnmarshalDuration(b []byte, wtyp Type, m *time.Duration) (int, error) {
	if wtyp != BytesType {
		return 0, ErrUnknown
	}
	v, n := protowire.ConsumeBytes(b)
	if n < 0 {
		return 0, ErrDecode
	}
	var secs, nanos int64
	for len(v) > 0 {
		num, wtyp, n := protowire.ConsumeTag(v)
		if n < 0 {
			return 0, ErrDecode
		}
		if num > protowire.MaxValidNumber {
			return 0, ErrDecode
		}
		v = v[n:]
		var err error = ErrUnknown
		switch num {
		case 1:
			n, err = UnmarshalInt64(v, wtyp, &secs)
		case 2:
			n, err = UnmarshalInt64(v, wtyp, &nanos)
		}
		if err == ErrUnknown {
			n := protowire.ConsumeFieldValue(num, wtyp, v)
			if n < 0 {
				return 0, ErrDecode
			}
		} else if err != nil {
			return 0, err
		}
		v = v[n:]
	}
	d := time.Duration(secs) * time.Second
	overflow := d/time.Second != time.Duration(secs)
	d += time.Duration(nanos) * time.Nanosecond
	overflow = overflow || (secs < 0 && nanos < 0 && d > 0)
	overflow = overflow || (secs > 0 && nanos > 0 && d < 0)
	if overflow {
		switch {
		case secs < 0:
			*m = time.Duration(math.MinInt64)
		case secs > 0:
			*m = time.Duration(math.MaxInt64)
		}
	} else {
		*m = d
	}
	return n, nil
}

func UnmarshalEmpty(b []byte, wtyp Type) (int, error) {
	if wtyp != BytesType {
		return 0, ErrUnknown
	}
	_, n := protowire.ConsumeBytes(b)
	if n < 0 {
		return 0, ErrDecode
	}
	return n, nil
}

// errUnknown is used internally to indicate fields which should be added
// to the unknown field set of a message. It is never returned from an exported
// function.
var ErrUnknown = errors.New("BUG: internal error (unknown)")

var ErrDecode = errors.New("cannot parse invalid wire-format data")
