package kdsync

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"time"
)

func AppendJson(b []byte, v any) ([]byte, error) {
	switch t := v.(type) {
	case bool:
		return strconv.AppendBool(b, t), nil
	case int32:
		return strconv.AppendInt(b, int64(t), 10), nil
	case uint32:
		return strconv.AppendUint(b, uint64(t), 10), nil
	case int64:
		return strconv.AppendQuote(b, strconv.FormatInt(t, 10)), nil
	case uint64:
		return strconv.AppendQuote(b, strconv.FormatUint(t, 10)), nil
	case float32:
		return strconv.AppendFloat(b, float64(t), 'f', -1, 32), nil
	case float64:
		return strconv.AppendFloat(b, t, 'f', -1, 64), nil
	case string:
		return strconv.AppendQuote(nil, t), nil
	case []byte:
		return base64.StdEncoding.AppendEncode(b, t), nil
	case time.Time:
		return append(b, "{Seconds: "+strconv.FormatInt(t.Unix(), 10)+", Nanos: "+strconv.Itoa(t.Nanosecond())+"}"...), nil
	case time.Duration:
		nanos := t.Nanoseconds()
		secs := nanos / 1e9
		nanos -= secs * 1e9
		return append(b, "{Seconds: "+strconv.FormatInt(secs, 10)+", Nanos: "+strconv.FormatInt(nanos, 10)+"}"...), nil
	case struct{}:
		return append(b, "{}"...), nil
	default:
		return fmt.Append(b, v), nil
	}
}
