package wire

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"time"
)

func Format(v any) string {
	switch t := v.(type) {
	case bool:
		return strconv.FormatBool(t)
	case int32:
		return strconv.FormatInt(int64(t), 10)
	case uint32:
		return strconv.FormatUint(uint64(t), 10)
	case int64:
		return string(strconv.AppendQuote(nil, strconv.FormatInt(t, 10)))
	case uint64:
		return string(strconv.AppendQuote(nil, strconv.FormatUint(t, 10)))
	case float32:
		return strconv.FormatFloat(float64(t), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(t, 'f', -1, 64)
	case string:
		return string(strconv.AppendQuote(nil, t))
	case []byte:
		return base64.StdEncoding.EncodeToString(t)
	case time.Time:
		return "{Seconds: " + strconv.FormatInt(t.Unix(), 10) + ", Nanos: " + strconv.Itoa(t.Nanosecond()) + "}"
	case time.Duration:
		nanos := t.Nanoseconds()
		secs := nanos / 1e9
		nanos -= secs * 1e9
		return "{Seconds: " + strconv.FormatInt(secs, 10) + ", Nanos: " + strconv.FormatInt(nanos, 10) + "}"
	case struct{}:
		return "{}"
	default:
		return fmt.Sprint(v)
	}
}

func FormatBool(v bool) string {
	return strconv.FormatBool(v)
}

func FormatInt32(v int32) string {
	return strconv.FormatInt(int64(v), 10)
}

func FormatSint32(v int32) string {
	return strconv.FormatInt(int64(v), 10)
}

func FormatUint32(v uint32) string {
	return strconv.FormatUint(uint64(v), 10)
}

func FormatInt64(v int64) string {
	return strconv.FormatInt(v, 10)
}

func FormatSint64(v int64) string {
	return strconv.FormatInt(v, 10)
}

func FormatUint64(v uint64) string {
	return strconv.FormatUint(v, 10)
}

func FormatSfixed32(v int32) string {
	return strconv.FormatInt(int64(v), 10)
}

func FormatFixed32(v uint32) string {
	return strconv.FormatUint(uint64(v), 10)
}

func FormatFloat(v float32) string {
	return strconv.FormatFloat(float64(v), 'f', -1, 32)
}

func FormatSfixed64(v int64) string {
	return strconv.FormatInt(v, 10)
}

func FormatFixed64(v uint64) string {
	return strconv.FormatUint(v, 10)
}

func FormatDouble(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func FormatString(v string) string {
	return v
}

func FormatBytes(v []byte) string {
	return base64.StdEncoding.EncodeToString(v)
}

func FormatTimestamp(v time.Time) string {
	return "{Seconds: " + strconv.FormatInt(v.Unix(), 10) + ", Nanos: " + strconv.Itoa(v.Nanosecond()) + "}"
}

func FormatDuration(v time.Duration) string {
	nanos := v.Nanoseconds()
	secs := nanos / 1e9
	nanos -= secs * 1e9
	return "{Seconds: " + strconv.FormatInt(secs, 10) + ", Nanos: " + strconv.FormatInt(nanos, 10) + "}"
}

func FormatEmpty(_ struct{}) string {
	return "{}"
}
