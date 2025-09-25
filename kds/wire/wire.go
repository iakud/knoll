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

const (
	MapClearFieldNumber  Number = 1
	MapDeleteFieldNumber Number = 2
	MapEntryFieldNumber  Number = 3
)

const (
	MapEntryKeyFieldNumber   Number = 1
	MapEntryValueFieldNumber Number = 2
)

type Type = protowire.Type

const (
	VarintType  Type = protowire.VarintType
	Fixed32Type Type = protowire.Fixed32Type
	Fixed64Type Type = protowire.Fixed64Type
	BytesType   Type = protowire.BytesType
)

type MarshalFunc func(b []byte) ([]byte, error)
type UnmarshalFunc func(b []byte) error

type Marshaler interface {
	Marshal(b []byte) ([]byte, error)
	MarshalDirty(b []byte) ([]byte, error)
}

type Unmarshaler interface {
	Unmarshal(b []byte) error
}

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

func AppendMessage(b []byte, m Marshaler) ([]byte, error) {
	var pos int
	var err error
	b, pos = AppendSpeculativeLength(b)
	b, err = m.Marshal(b)
	if err != nil {
		return b, err
	}
	b = FinishSpeculativeLength(b, pos)
	return b, nil
}

func AppendMessageDirty(b []byte, m Marshaler) ([]byte, error) {
	var pos int
	var err error
	b, pos = AppendSpeculativeLength(b)
	b, err = m.MarshalDirty(b)
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
	b = protowire.AppendVarint(b, uint64(v.Nanosecond()))
	b = FinishSpeculativeLength(b, pos)
	return b
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
	b = FinishSpeculativeLength(b, pos)
	return b
}

func AppendEmpty(b []byte, _ struct{}) []byte {
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

func ConsumeTag(b []byte) (Number, Type, int, error) {
	num, wtyp, n := protowire.ConsumeTag(b)
	if n < 0 || num > protowire.MaxValidNumber {
		return num, wtyp, n, ErrDecode
	}
	return num, wtyp, n, nil
}

func ConsumeFieldValue(num Number, wtyp Type, b []byte) (int, error) {
	n := protowire.ConsumeFieldValue(num, wtyp, b)
	if n < 0 {
		return n, ErrDecode
	}
	return n, nil
}

func ConsumeBool(b []byte) (bool, int, error) {
	v, n := protowire.ConsumeVarint(b)
	if n < 0 {
		return false, 0, ErrDecode
	}
	return protowire.DecodeBool(v), n, nil
}

func ConsumeInt32(b []byte) (int32, int, error) {
	v, n := protowire.ConsumeVarint(b)
	if n < 0 {
		return 0, 0, ErrDecode
	}
	return int32(v), n, nil
}

func ConsumeSint32(b []byte) (int32, int, error) {
	v, n := protowire.ConsumeVarint(b)
	if n < 0 {
		return 0, 0, ErrDecode
	}
	return int32(protowire.DecodeZigZag(v & math.MaxUint32)), n, nil
}

func ConsumeUint32(b []byte) (uint32, int, error) {
	v, n := protowire.ConsumeVarint(b)
	if n < 0 {
		return 0, 0, ErrDecode
	}
	return uint32(v), n, nil
}

func ConsumeInt64(b []byte) (int64, int, error) {
	v, n := protowire.ConsumeVarint(b)
	if n < 0 {
		return 0, 0, ErrDecode
	}
	return int64(v), n, nil
}

func ConsumeSint64(b []byte) (int64, int, error) {
	v, n := protowire.ConsumeVarint(b)
	if n < 0 {
		return 0, 0, ErrDecode
	}
	return int64(protowire.DecodeZigZag(v & math.MaxUint32)), n, nil
}

func ConsumeUint64(b []byte) (uint64, int, error) {
	v, n := protowire.ConsumeVarint(b)
	if n < 0 {
		return 0, 0, ErrDecode
	}
	return v, n, nil
}

func ConsumeSfixed32(b []byte) (int32, int, error) {
	v, n := protowire.ConsumeFixed32(b)
	if n < 0 {
		return 0, 0, ErrDecode
	}
	return int32(v), n, nil
}

func ConsumeFixed32(b []byte) (uint32, int, error) {
	v, n := protowire.ConsumeFixed32(b)
	if n < 0 {
		return 0, 0, ErrDecode
	}
	return v, n, nil
}

func ConsumeFloat(b []byte) (float32, int, error) {
	v, n := protowire.ConsumeFixed32(b)
	if n < 0 {
		return 0, 0, ErrDecode
	}
	return math.Float32frombits(v), n, nil
}

func ConsumeSfixed64(b []byte) (int64, int, error) {
	v, n := protowire.ConsumeFixed64(b)
	if n < 0 {
		return 0, 0, ErrDecode
	}
	return int64(v), n, nil
}

func ConsumeFixed64(b []byte) (uint64, int, error) {
	v, n := protowire.ConsumeFixed64(b)
	if n < 0 {
		return 0, 0, ErrDecode
	}
	return v, n, nil
}

func ConsumeDouble(b []byte) (float64, int, error) {
	v, n := protowire.ConsumeFixed64(b)
	if n < 0 {
		return 0, 0, ErrDecode
	}
	return math.Float64frombits(v), n, nil
}

func ConsumeString(b []byte) (string, int, error) {
	v, n := protowire.ConsumeString(b)
	if n < 0 {
		return v, 0, ErrDecode
	}
	return v, n, nil
}

func ConsumeBytes(b []byte) ([]byte, int, error) {
	v, n := protowire.ConsumeBytes(b)
	if n < 0 {
		return v, 0, ErrDecode
	}
	return v, n, nil
}

func ConsumeTimestamp(b []byte) (time.Time, int, error) {
	var sec, nsec int64
	v, n := protowire.ConsumeBytes(b)
	if n < 0 {
		return time.Time{}, 0, ErrDecode
	}
	for len(v) > 0 {
		num, wtyp, n := protowire.ConsumeTag(v)
		if n < 0 {
			return time.Time{}, 0, ErrDecode
		}
		if num > protowire.MaxValidNumber {
			return time.Time{}, 0, ErrDecode
		}
		v = v[n:]
		err := ErrUnknown
		switch num {
		case 1:
			if wtyp != protowire.VarintType {
				break
			}
			sec, n, err = ConsumeInt64(v)
		case 2:
			if wtyp != protowire.VarintType {
				break
			}
			nsec, n, err = ConsumeInt64(v)
		}
		if err == ErrUnknown {
			n = protowire.ConsumeFieldValue(num, wtyp, v)
			if n < 0 {
				return time.Time{}, 0, ErrDecode
			}
		} else if err != nil {
			return time.Time{}, 0, err
		}
		v = v[n:]
	}
	return time.Unix(sec, nsec), n, nil
}

func ConsumeDuration(b []byte) (time.Duration, int, error) {
	v, n := protowire.ConsumeBytes(b)
	if n < 0 {
		return time.Duration(0), 0, ErrDecode
	}
	var secs, nanos int64
	for len(v) > 0 {
		num, wtyp, n := protowire.ConsumeTag(v)
		if n < 0 {
			return 0, 0, ErrDecode
		}
		if num > protowire.MaxValidNumber {
			return 0, 0, ErrDecode
		}
		v = v[n:]
		err := ErrUnknown
		switch num {
		case 1:
			if wtyp != protowire.VarintType {
				break
			}
			secs, n, err = ConsumeInt64(v)
		case 2:
			if wtyp != protowire.VarintType {
				break
			}
			nanos, n, err = ConsumeInt64(v)
		}
		if err == ErrUnknown {
			n := protowire.ConsumeFieldValue(num, wtyp, v)
			if n < 0 {
				return 0, 0, ErrDecode
			}
		} else if err != nil {
			return 0, 0, err
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
			return time.Duration(math.MinInt64), n, nil
		case secs > 0:
			return time.Duration(math.MaxInt64), n, nil
		}
	}
	return d, n, nil
}

func ConsumeEmpty(b []byte) (struct{}, int, error) {
	_, n := protowire.ConsumeBytes(b)
	if n < 0 {
		return struct{}{}, 0, ErrDecode
	}
	return struct{}{}, n, nil
}

func ConsumeMessage(b []byte, m Unmarshaler) (int, error) {
	b, n := protowire.ConsumeBytes(b)
	if n < 0 {
		return 0, ErrDecode
	}
	if err := m.Unmarshal(b); err != nil {
		return 0, err
	}
	return n, nil
}

// errUnknown is used internally to indicate fields which should be added
// to the unknown field set of a message. It is never returned from an exported
// function.
var ErrUnknown = errors.New("BUG: internal error (unknown)")

var ErrDecode = errors.New("cannot parse invalid wire-format data")
