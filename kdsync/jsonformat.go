package kdsync

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"time"
)

const indent = "  "

type JSONMarshaler interface {
	MarshalJSONIndent(b []byte, prefix, indent string) ([]byte, error)
}

func MarshalJSONIndent(b []byte, v any, prefix, indent string) ([]byte, error) {
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
		return strconv.AppendQuote(b, t), nil
	case []byte:
		return strconv.AppendQuote(b, base64.StdEncoding.EncodeToString(t)), nil
	case time.Time:
		b = append(b, '{')
		b = append(b, '\n')
		b = append(b, prefix...)
		b = append(b, indent...)
		b = append(b, `"Seconds": `...)
		b = strconv.AppendQuote(b, strconv.FormatInt(t.Unix(), 10))
		b = append(b, ',', '\n')
		b = append(b, prefix...)
		b = append(b, indent...)
		b = append(b, `"Nanos": `...)
		b = strconv.AppendInt(b, int64(t.Nanosecond()), 10)
		b = append(b, '\n')
		b = append(b, prefix...)
		b = append(b, '}')
		return b, nil
	case time.Duration:
		nanos := t.Nanoseconds()
		secs := nanos / 1e9
		nanos -= secs * 1e9
		b = append(b, '{')
		b = append(b, '\n')
		b = append(b, prefix...)
		b = append(b, indent...)
		b = append(b, `"Seconds": `...)
		b = strconv.AppendQuote(b, strconv.FormatInt(secs, 10))
		b = append(b, ',', '\n')
		b = append(b, prefix...)
		b = append(b, indent...)
		b = append(b, `"Nanos": `...)
		b = strconv.AppendInt(b, nanos, 10)
		b = append(b, '\n')
		b = append(b, prefix...)
		b = append(b, '}')
		return b, nil
	case struct{}:
		return append(b, '{', '}'), nil
	case JSONMarshaler:
		return t.MarshalJSONIndent(b, prefix, indent)
	default:
		return fmt.Append(b, v), nil
	}
}
