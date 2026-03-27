package wire

import (
	"time"
)

func UnmarshalMessage(b []byte, wtyp Type, m Unmarshaler) (int, error) {
	if wtyp != BytesType {
		return 0, ErrUnknown
	}
	return ConsumeMessage(b, m)
}

func UnmarshalBool(b []byte, wtyp Type) (bool, int, error) {
	if wtyp != VarintType {
		return false, 0, ErrUnknown
	}
	return ConsumeBool(b)
}

func UnmarshalInt32(b []byte, wtyp Type) (int32, int, error) {
	if wtyp != VarintType {
		return 0, 0, ErrUnknown
	}
	return ConsumeInt32(b)
}

func UnmarshalSint32(b []byte, wtyp Type) (int32, int, error) {
	if wtyp != VarintType {
		return 0, 0, ErrUnknown
	}
	return ConsumeSint32(b)
}

func UnmarshalUint32(b []byte, wtyp Type) (uint32, int, error) {
	if wtyp != VarintType {
		return 0, 0, ErrUnknown
	}
	return ConsumeUint32(b)
}

func UnmarshalInt64(b []byte, wtyp Type) (int64, int, error) {
	if wtyp != VarintType {
		return 0, 0, ErrUnknown
	}
	return ConsumeInt64(b)
}

func UnmarshalSint64(b []byte, wtyp Type) (int64, int, error) {
	if wtyp != VarintType {
		return 0, 0, ErrUnknown
	}
	return ConsumeSint64(b)
}

func UnmarshalUint64(b []byte, wtyp Type) (uint64, int, error) {
	if wtyp != VarintType {
		return 0, 0, ErrUnknown
	}
	return ConsumeUint64(b)
}

func UnmarshalSfixed32(b []byte, wtyp Type) (int32, int, error) {
	if wtyp != Fixed32Type {
		return 0, 0, ErrUnknown
	}
	return ConsumeSfixed32(b)
}

func UnmarshalFixed32(b []byte, wtyp Type) (uint32, int, error) {
	if wtyp != Fixed32Type {
		return 0, 0, ErrUnknown
	}
	return ConsumeFixed32(b)
}

func UnmarshalFloat(b []byte, wtyp Type) (float32, int, error) {
	if wtyp != Fixed32Type {
		return 0, 0, ErrUnknown
	}
	return ConsumeFloat(b)
}

func UnmarshalSfixed64(b []byte, wtyp Type) (int64, int, error) {
	if wtyp != Fixed64Type {
		return 0, 0, ErrUnknown
	}
	return ConsumeSfixed64(b)
}

func UnmarshalFixed64(b []byte, wtyp Type) (uint64, int, error) {
	if wtyp != Fixed64Type {
		return 0, 0, ErrUnknown
	}
	return ConsumeFixed64(b)
}

func UnmarshalDouble(b []byte, wtyp Type) (float64, int, error) {
	if wtyp != Fixed64Type {
		return 0, 0, ErrUnknown
	}
	return ConsumeDouble(b)
}

func UnmarshalString(b []byte, wtyp Type) (string, int, error) {
	if wtyp != BytesType {
		return "", 0, ErrUnknown
	}
	return ConsumeString(b)
}

func UnmarshalBytes(b []byte, wtyp Type) ([]byte, int, error) {
	if wtyp != BytesType {
		return nil, 0, ErrUnknown
	}
	return ConsumeBytes(b)
}

func UnmarshalTimestamp(b []byte, wtyp Type) (time.Time, int, error) {
	if wtyp != BytesType {
		return time.Time{}, 0, ErrUnknown
	}
	return ConsumeTimestamp(b)
}

func UnmarshalDuration(b []byte, wtyp Type) (time.Duration, int, error) {
	if wtyp != BytesType {
		return 0, 0, ErrUnknown
	}
	return ConsumeDuration(b)
}

func UnmarshalEmpty(b []byte, wtyp Type) (struct{}, int, error) {
	if wtyp != BytesType {
		return struct{}{}, 0, ErrUnknown
	}
	return ConsumeEmpty(b)
}
