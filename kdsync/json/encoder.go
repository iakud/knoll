package json

import (
	"encoding/base64"
	"strconv"
	"time"
)

type tokenType byte

const (
	None tokenType = iota
	PropertyName
	String
	Number
	True
	False
	Null
	StartObject
	EndObject
	StartArray
	EndArray
	Comment
)

type Writer interface {
	WriteJSON(e *Encoder)
}

type Encoder struct {
	buf []byte

	indented      bool
	_currentDepth int
	_tokenType    tokenType
	_indentLength int
}

func (e *Encoder) String() string {
	return string(e.buf)
}

func (e *Encoder) WriteStartObject() {
	e.WriteStart('{')
	e._tokenType = StartObject
}

func (e *Encoder) WriteEndObject() {
	e.WriteEnd('}')
	e._tokenType = EndObject
}

func (e *Encoder) WriteStartArray() {
	e.WriteStart('[')
	e._tokenType = StartArray
}

func (e *Encoder) WriteEndArray() {
	e.WriteEnd(']')
	e._tokenType = EndArray
}

func (e *Encoder) WriteStart(token byte) {
	if e.indented {
		e.WriteStartIndented(token)
	} else {
		e.WriteStartMinimized(token)
	}
}

func (e *Encoder) WriteStartMinimized(token byte) {
	if e._tokenType != None && e._tokenType != PropertyName && e._tokenType != StartObject && e._tokenType != StartArray {
		e.buf = append(e.buf, ',')
	}
	e.buf = append(e.buf, token)
	e._currentDepth++
}

func (e *Encoder) WriteStartIndented(token byte) {
	if e._tokenType != None && e._tokenType != PropertyName && e._tokenType != StartObject && e._tokenType != StartArray {
		e.buf = append(e.buf, ',')
	}
	if e._tokenType != None && e._tokenType != PropertyName {
		e.buf = append(e.buf, '\n')
		e.WriteIndentation(e._currentDepth)
	}
	e.buf = append(e.buf, token)
	e._currentDepth++
}

func (e *Encoder) WriteEnd(token byte) {
	if e.indented {
		e.WriteEndIndented(token)
	} else {
		e.WriteEndMinimized(token)
	}
}

func (e *Encoder) WriteEndMinimized(token byte) {
	e._currentDepth--
	e.buf = append(e.buf, token)
}

func (e *Encoder) WriteEndIndented(token byte) {
	e._currentDepth--

	if e._tokenType != StartObject && e._tokenType != StartArray {
		e.buf = append(e.buf, '\n')
		e.WriteIndentation(e._currentDepth)
	}
	e.buf = append(e.buf, token)
}

func (e *Encoder) WritePropertyName(name string) {
	e.WriteStringValue(name)
	e.buf = append(e.buf, ':')
	e._tokenType = PropertyName
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

func (e *Encoder) WriteMessage(name string, v Writer) {
	e.WritePropertyName(name)
	e.WriteMessageValue(v)
}

func (e *Encoder) WriteNullValue() {
	e.buf = append(e.buf, "null"...)
	e._tokenType = Null
}

func (e *Encoder) WriteBoolValue(v bool) {
	e.buf = strconv.AppendBool(e.buf, v)
	if v {
		e._tokenType = True
	} else {
		e._tokenType = False
	}
}

func (e *Encoder) WriteInt32Value(v int32) {
	e.buf = strconv.AppendInt(e.buf, int64(v), 10)
	e._tokenType = Number
}

func (e *Encoder) WriteUint32Value(v uint32) {
	e.buf = strconv.AppendUint(e.buf, uint64(v), 10)
	e._tokenType = Number
}

func (e *Encoder) WriteInt64Value(v int64) {
	e.buf = strconv.AppendInt(e.buf, v, 10)
	e._tokenType = Number
}

func (e *Encoder) WriteUint64Value(v uint64) {
	e.buf = strconv.AppendUint(e.buf, v, 10)
	e._tokenType = Number
}

func (e *Encoder) WriteFloat32Value(v float32) {
	e.buf = strconv.AppendFloat(e.buf, float64(v), 'f', -1, 32)
	e._tokenType = Number
}

func (e *Encoder) WriteFloat64Value(v float64) {
	e.buf = strconv.AppendFloat(e.buf, v, 'f', -1, 64)
	e._tokenType = Number
}

func (e *Encoder) WriteBytesValue(v []byte) {
	e.buf = append(e.buf, '"')
	e.buf = base64.StdEncoding.AppendEncode(e.buf, v)
	e.buf = append(e.buf, '"')
	e._tokenType = String
}

func (e *Encoder) WriteStringValue(v string) {
	e.buf = strconv.AppendQuote(e.buf, v)
	e._tokenType = String
}

func (e *Encoder) WriteTimestampValue(v time.Time) {
	e.WriteStartObject()
	e.WriteInt64("Seconds", v.Unix())
	e.WriteInt64("Nanos", int64(v.Nanosecond()))
	e.WriteEndObject()
}

func (e *Encoder) WriteDurationValue(v time.Duration) {
	nanos := v.Nanoseconds()
	secs := nanos / 1e9
	nanos -= secs * 1e9
	e.WriteStartObject()
	e.WriteInt64("Seconds", secs)
	e.WriteInt64("Nanos", nanos)
	e.WriteEndObject()
}

func (e *Encoder) WriteEmptyValue(v struct{}) {
	e.WriteStartObject()
	e.WriteEndObject()
}

func (e *Encoder) WriteMessageValue(v Writer) {
	v.WriteJSON(e)
}

func (e *Encoder) WriteIndentation(depth int) {
	for i := 0; i < depth*e._indentLength; i++ {
		e.buf = append(e.buf, ' ')
	}
}
