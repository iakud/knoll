package kdsjson

import (
	"encoding/base64"
	"strconv"
	"time"

	"github.com/iakud/knoll/kdsync/kdsjson/internal"
)

type Encoder struct {
	buf []byte

	indented     bool
	depth        int
	tokenType    internal.TokenType
	indentLength int
}

func (e *Encoder) String() string {
	return string(e.buf)
}

func (e *Encoder) WriteStartObject() {
	e.writeStart('{')
	e.tokenType = internal.TokenStartObject
}

func (e *Encoder) WriteEndObject() {
	e.writeEnd('}')
	e.tokenType = internal.TokenEndObject
}

func (e *Encoder) WriteStartArray() {
	e.writeStart('[')
	e.tokenType = internal.TokenStartArray
}

func (e *Encoder) WriteEndArray() {
	e.writeEnd(']')
	e.tokenType = internal.TokenEndArray
}

func (e *Encoder) writeStart(token byte) {
	if e.indented {
		e.writeStartIndented(token)
	} else {
		e.writeStartMinimized(token)
	}
}

func (e *Encoder) writeStartMinimized(token byte) {
	if e.tokenType != internal.TokenNone && e.tokenType != internal.TokenPropertyName && e.tokenType != internal.TokenStartObject && e.tokenType != internal.TokenStartArray {
		e.buf = append(e.buf, ',')
	}
	e.buf = append(e.buf, token)
	e.depth++
}

func (e *Encoder) writeStartIndented(token byte) {
	if e.tokenType != internal.TokenNone && e.tokenType != internal.TokenPropertyName && e.tokenType != internal.TokenStartObject && e.tokenType != internal.TokenStartArray {
		e.buf = append(e.buf, ',')
	}
	if e.tokenType != internal.TokenNone && e.tokenType != internal.TokenPropertyName {
		e.buf = append(e.buf, '\n')
		e.writeIndentation(e.depth)
	}
	e.buf = append(e.buf, token)
	e.depth++
}

func (e *Encoder) writeEnd(token byte) {
	if e.indented {
		e.writeEndIndented(token)
	} else {
		e.writeEndMinimized(token)
	}
}

func (e *Encoder) writeEndMinimized(token byte) {
	e.depth--
	e.buf = append(e.buf, token)
}

func (e *Encoder) writeEndIndented(token byte) {
	e.depth--
	if e.tokenType != internal.TokenStartObject && e.tokenType != internal.TokenStartArray {
		e.buf = append(e.buf, '\n')
		e.writeIndentation(e.depth)
	}
	e.buf = append(e.buf, token)
}

func (e *Encoder) WritePropertyName(name string) {
	if e.indented {
		e.writePropertyNameIndented(name)
	} else {
		e.writePropertyNameMinimized(name)
	}
	e.tokenType = internal.TokenPropertyName
}

func (e *Encoder) writePropertyNameMinimized(name string) {
	if e.tokenType != internal.TokenStartObject {
		e.buf = append(e.buf, ',')
	}
	e.writeEscapedString(name)
	e.buf = append(e.buf, ':')
}

func (e *Encoder) writePropertyNameIndented(name string) {
	if e.tokenType != internal.TokenStartObject {
		e.buf = append(e.buf, ',')
	}
	e.buf = append(e.buf, '\n')
	e.writeIndentation(e.depth)
	e.writeEscapedString(name)
	e.buf = append(e.buf, ':')
	e.buf = append(e.buf, ' ')
}

func (e *Encoder) WriteNull(name string) {
	e.WritePropertyName(name)
	e.WriteNullValue()
}

func (e *Encoder) WriteBool(name string, v bool) {
	e.WritePropertyName(name)
	e.WriteBoolValue(v)
}

func (e *Encoder) WriteInt32(name string, v int32) {
	e.WritePropertyName(name)
	e.WriteInt32Value(v)
}

func (e *Encoder) WriteUint32(name string, v uint32) {
	e.WritePropertyName(name)
	e.WriteUint32Value(v)
}

func (e *Encoder) WriteInt64(name string, v int64) {
	e.WritePropertyName(name)
	e.WriteInt64Value(v)
}

func (e *Encoder) WriteUint64(name string, v uint64) {
	e.WritePropertyName(name)
	e.WriteUint64Value(v)
}

func (e *Encoder) WriteFloat32(name string, v float32) {
	e.WritePropertyName(name)
	e.WriteFloat32Value(v)
}

func (e *Encoder) WriteFloat64(name string, v float64) {
	e.WritePropertyName(name)
	e.WriteFloat64Value(v)
}

func (e *Encoder) WriteBytes(name string, v []byte) {
	e.WritePropertyName(name)
	e.WriteBytesValue(v)
}

func (e *Encoder) WriteString(name string, v string) {
	e.WritePropertyName(name)
	e.WriteStringValue(v)
}

func (e *Encoder) WriteTimestamp(name string, v time.Time) {
	e.WritePropertyName(name)
	e.WriteTimestampValue(v)
}

func (e *Encoder) WriteDuration(name string, v time.Duration) {
	e.WritePropertyName(name)
	e.WriteDurationValue(v)
}

func (e *Encoder) WriteEmpty(name string, v struct{}) {
	e.WritePropertyName(name)
	e.WriteEmptyValue(v)
}

func (e *Encoder) Write(name string, v Marshaler) {
	e.WritePropertyName(name)
	e.WriteValue(v)
}

func (e *Encoder) WriteNullValue() {
	e.writeValueSeparator()
	e.buf = append(e.buf, "null"...)
	e.tokenType = internal.TokenNull
}

func (e *Encoder) WriteBoolValue(v bool) {
	e.writeValueSeparator()
	e.buf = strconv.AppendBool(e.buf, v)
	if v {
		e.tokenType = internal.TokenTrue
	} else {
		e.tokenType = internal.TokenFalse
	}
}

func (e *Encoder) WriteInt32Value(v int32) {
	e.writeValueSeparator()
	e.buf = strconv.AppendInt(e.buf, int64(v), 10)
	e.tokenType = internal.TokenNumber
}

func (e *Encoder) WriteUint32Value(v uint32) {
	e.writeValueSeparator()
	e.buf = strconv.AppendUint(e.buf, uint64(v), 10)
	e.tokenType = internal.TokenNumber
}

func (e *Encoder) WriteInt64Value(v int64) {
	e.writeValueSeparator()
	e.buf = append(e.buf, '"')
	e.buf = strconv.AppendInt(e.buf, v, 10)
	e.buf = append(e.buf, '"')
	e.tokenType = internal.TokenNumber
}

func (e *Encoder) WriteUint64Value(v uint64) {
	e.writeValueSeparator()
	e.buf = append(e.buf, '"')
	e.buf = strconv.AppendUint(e.buf, v, 10)
	e.buf = append(e.buf, '"')
	e.tokenType = internal.TokenNumber
}

func (e *Encoder) WriteFloat32Value(v float32) {
	e.writeValueSeparator()
	e.buf = strconv.AppendFloat(e.buf, float64(v), 'f', -1, 32)
	e.tokenType = internal.TokenNumber
}

func (e *Encoder) WriteFloat64Value(v float64) {
	e.writeValueSeparator()
	e.buf = strconv.AppendFloat(e.buf, v, 'f', -1, 64)
	e.tokenType = internal.TokenNumber
}

func (e *Encoder) WriteBytesValue(v []byte) {
	e.writeValueSeparator()
	e.buf = append(e.buf, '"')
	e.buf = base64.StdEncoding.AppendEncode(e.buf, v)
	e.buf = append(e.buf, '"')
	e.tokenType = internal.TokenString
}

func (e *Encoder) WriteStringValue(v string) {
	e.writeValueSeparator()
	e.writeEscapedString(v)
	e.tokenType = internal.TokenString
}

func (e *Encoder) WriteTimestampValue(v time.Time) {
	e.WriteStartObject()
	e.WriteInt64("Seconds", v.Unix())
	e.WriteInt32("Nanos", int32(v.Nanosecond()))
	e.WriteEndObject()
}

func (e *Encoder) WriteDurationValue(v time.Duration) {
	nanos := v.Nanoseconds()
	secs := nanos / 1e9
	nanos -= secs * 1e9
	e.WriteStartObject()
	e.WriteInt64("Seconds", secs)
	e.WriteInt32("Nanos", int32(nanos))
	e.WriteEndObject()
}

func (e *Encoder) WriteEmptyValue(v struct{}) {
	e.WriteStartObject()
	e.WriteEndObject()
}

func (e *Encoder) WriteValue(v Marshaler) {
	if v == nil {
		e.WriteNullValue()
		return
	}
	v.WriteJSON(e)
}

func (e *Encoder) writeEscapedString(v string) {
	e.buf = strconv.AppendQuote(e.buf, v)
}

func (e *Encoder) writeValueSeparator() {
	if e.indented {
		e.writeValueSeparatorIndented()
	} else {
		e.writeValueSeparatorMinimized()
	}
}

func (e *Encoder) writeValueSeparatorMinimized() {
	if e.tokenType != internal.TokenPropertyName && e.tokenType != internal.TokenStartArray {
		e.buf = append(e.buf, ',')
	}
}

func (e *Encoder) writeValueSeparatorIndented() {
	if e.tokenType != internal.TokenPropertyName && e.tokenType != internal.TokenStartArray {
		e.buf = append(e.buf, ',')
	}
	if e.tokenType != internal.TokenPropertyName {
		e.buf = append(e.buf, '\n')
		e.writeIndentation(e.depth)
	}
}

func (e *Encoder) writeIndentation(depth int) {
	for i := 0; i < depth*e.indentLength; i++ {
		e.buf = append(e.buf, ' ')
	}
}
