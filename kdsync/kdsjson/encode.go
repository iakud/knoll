package kdsjson

import (
	"encoding/base64"
	"errors"
	"math"
	"math/bits"
	"strconv"
	"time"
	"unicode/utf8"
)

type tokenKind byte

const (
	tokenNone tokenKind = iota
	tokenName
	tokenScalar
	tokenStartObject
	tokenEndObject
	tokenStartArray
	tokenEndArray
)

type Encoder struct {
	out []byte

	indented     bool
	depth        int
	tokenType    tokenKind
	indentLength int
}

func (e *Encoder) String() string {
	return string(e.out)
}

func (e *Encoder) WriteStartObject() {
	e.writeStart('{')
	e.tokenType = tokenStartObject
}

func (e *Encoder) WriteEndObject() {
	e.writeEnd('}')
	e.tokenType = tokenEndObject
}

func (e *Encoder) WriteStartArray() {
	e.writeStart('[')
	e.tokenType = tokenStartArray
}

func (e *Encoder) WriteEndArray() {
	e.writeEnd(']')
	e.tokenType = tokenEndArray
}

func (e *Encoder) writeStart(token byte) {
	if e.indented {
		e.writeStartIndented(token)
	} else {
		e.writeStartMinimized(token)
	}
}

func (e *Encoder) writeStartMinimized(token byte) {
	if e.tokenType != tokenNone && e.tokenType != tokenName && e.tokenType != tokenStartObject && e.tokenType != tokenStartArray {
		e.out = append(e.out, ',')
	}
	e.out = append(e.out, token)
	e.depth++
}

func (e *Encoder) writeStartIndented(token byte) {
	if e.tokenType != tokenNone && e.tokenType != tokenName && e.tokenType != tokenStartObject && e.tokenType != tokenStartArray {
		e.out = append(e.out, ',')
	}
	if e.tokenType != tokenNone && e.tokenType != tokenName {
		e.out = append(e.out, '\n')
		e.writeIndentation(e.depth)
	}
	e.out = append(e.out, token)
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
	e.out = append(e.out, token)
}

func (e *Encoder) writeEndIndented(token byte) {
	e.depth--
	if e.tokenType != tokenStartObject && e.tokenType != tokenStartArray {
		e.out = append(e.out, '\n')
		e.writeIndentation(e.depth)
	}
	e.out = append(e.out, token)
}

func (e *Encoder) WriteName(name string) error {
	var err error
	if e.indented {
		err = e.writeNameIndented(name)
	} else {
		err = e.writeNameMinimized(name)
	}
	e.tokenType = tokenName
	return err
}

func (e *Encoder) writeNameMinimized(name string) error {
	if e.tokenType != tokenStartObject {
		e.out = append(e.out, ',')
	}
	err := e.writeEscapedString(name)
	e.out = append(e.out, ':')
	return err
}

func (e *Encoder) writeNameIndented(name string) error {
	if e.tokenType != tokenStartObject {
		e.out = append(e.out, ',')
	}
	e.out = append(e.out, '\n')
	e.writeIndentation(e.depth)
	err := e.writeEscapedString(name)
	e.out = append(e.out, ':')
	e.out = append(e.out, ' ')
	return err
}

func (e *Encoder) WriteNull(name string) error {
	if err := e.WriteName(name); err != nil {
		return err
	}
	e.WriteNullValue()
	return nil
}

func (e *Encoder) WriteBool(name string, v bool) error {
	if err := e.WriteName(name); err != nil {
		return err
	}
	e.WriteBoolValue(v)
	return nil
}

func (e *Encoder) WriteInt32(name string, v int32) error {
	if err := e.WriteName(name); err != nil {
		return err
	}
	e.WriteInt32Value(v)
	return nil
}

func (e *Encoder) WriteUint32(name string, v uint32) error {
	if err := e.WriteName(name); err != nil {
		return err
	}
	e.WriteUint32Value(v)
	return nil
}

func (e *Encoder) WriteInt64(name string, v int64) error {
	if err := e.WriteName(name); err != nil {
		return err
	}
	e.WriteInt64Value(v)
	return nil
}

func (e *Encoder) WriteUint64(name string, v uint64) error {
	if err := e.WriteName(name); err != nil {
		return err
	}
	e.WriteUint64Value(v)
	return nil
}

func (e *Encoder) WriteFloat32(name string, v float32) error {
	if err := e.WriteName(name); err != nil {
		return err
	}
	e.WriteFloat32Value(v)
	return nil
}

func (e *Encoder) WriteFloat64(name string, v float64) error {
	if err := e.WriteName(name); err != nil {
		return err
	}
	e.WriteFloat64Value(v)
	return nil
}

func (e *Encoder) WriteBytes(name string, v []byte) error {
	if err := e.WriteName(name); err != nil {
		return err
	}
	e.WriteBytesValue(v)
	return nil
}

func (e *Encoder) WriteString(name string, v string) error {
	if err := e.WriteName(name); err != nil {
		return err
	}
	if err := e.WriteStringValue(v); err != nil {
		return InvalidUTF8(name)
	}
	return nil
}

func (e *Encoder) WriteTimestamp(name string, v time.Time) error {
	if err := e.WriteName(name); err != nil {
		return err
	}
	e.WriteTimestampValue(v)
	return nil
}

func (e *Encoder) WriteDuration(name string, v time.Duration) error {
	if err := e.WriteName(name); err != nil {
		return err
	}
	e.WriteDurationValue(v)
	return nil
}

func (e *Encoder) WriteEmpty(name string, v struct{}) error {
	if err := e.WriteName(name); err != nil {
		return err
	}
	e.WriteEmptyValue(v)
	return nil
}

func (e *Encoder) Write(name string, v Marshaler) error {
	if err := e.WriteName(name); err != nil {
		return err
	}
	return e.WriteValue(v)
}

func (e *Encoder) WriteNullValue() {
	e.writeValueSeparator()
	e.out = append(e.out, "null"...)
	e.tokenType = tokenScalar
}

func (e *Encoder) WriteBoolValue(v bool) {
	e.writeValueSeparator()
	if v {
		e.out = append(e.out, "true"...)
	} else {
		e.out = append(e.out, "false"...)
	}
	e.tokenType = tokenScalar
}

func (e *Encoder) WriteInt32Value(v int32) {
	e.writeValueSeparator()
	e.out = strconv.AppendInt(e.out, int64(v), 10)
	e.tokenType = tokenScalar
}

func (e *Encoder) WriteUint32Value(v uint32) {
	e.writeValueSeparator()
	e.out = strconv.AppendUint(e.out, uint64(v), 10)
	e.tokenType = tokenScalar
}

func (e *Encoder) WriteInt64Value(v int64) {
	e.writeValueSeparator()
	e.out = append(e.out, '"')
	e.out = strconv.AppendInt(e.out, v, 10)
	e.out = append(e.out, '"')
	e.tokenType = tokenScalar
}

func (e *Encoder) WriteUint64Value(v uint64) {
	e.writeValueSeparator()
	e.out = append(e.out, '"')
	e.out = strconv.AppendUint(e.out, v, 10)
	e.out = append(e.out, '"')
	e.tokenType = tokenScalar
}

func (e *Encoder) WriteFloat32Value(v float32) {
	e.writeValueSeparator()
	e.out = appendFloat(e.out, float64(v), 32)
	e.tokenType = tokenScalar
}

func (e *Encoder) WriteFloat64Value(v float64) {
	e.writeValueSeparator()
	e.out = appendFloat(e.out, v, 64)
	e.tokenType = tokenScalar
}

func appendFloat(out []byte, n float64, bitSize int) []byte {
	switch {
	case math.IsNaN(n):
		return append(out, `"NaN"`...)
	case math.IsInf(n, +1):
		return append(out, `"Infinity"`...)
	case math.IsInf(n, -1):
		return append(out, `"-Infinity"`...)
	}
	fmt := byte('f')
	if abs := math.Abs(n); abs != 0 {
		if bitSize == 64 && (abs < 1e-6 || abs >= 1e21) ||
			bitSize == 32 && (float32(abs) < 1e-6 || float32(abs) >= 1e21) {
			fmt = 'e'
		}
	}
	out = strconv.AppendFloat(out, n, fmt, -1, bitSize)
	if fmt == 'e' {
		n := len(out)
		if n >= 4 && out[n-4] == 'e' && out[n-3] == '-' && out[n-2] == '0' {
			out[n-2] = out[n-1]
			out = out[:n-1]
		}
	}
	return out
}

func (e *Encoder) WriteBytesValue(v []byte) {
	e.writeValueSeparator()
	e.out = append(e.out, '"')
	e.out = base64.StdEncoding.AppendEncode(e.out, v)
	e.out = append(e.out, '"')
	e.tokenType = tokenScalar
}

func (e *Encoder) WriteStringValue(v string) error {
	e.writeValueSeparator()
	err := e.writeEscapedString(v)
	e.tokenType = tokenScalar
	return err
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

func (e *Encoder) WriteValue(v Marshaler) error {
	if v == nil {
		e.WriteNullValue()
		return nil
	}
	return v.WriteJSON(e)
}

func (e *Encoder) writeEscapedString(v string) error {
	var err error
	if e.out, err = appendString(e.out, v); err != nil {
		return err
	}
	return nil
}

// Sentinel error used for indicating invalid UTF-8.
var errInvalidUTF8 = errors.New("invalid UTF-8")

func appendString(out []byte, in string) ([]byte, error) {
	out = append(out, '"')
	i := indexNeedEscapeInString(in)
	in, out = in[i:], append(out, in[:i]...)
	for len(in) > 0 {
		switch r, n := utf8.DecodeRuneInString(in); {
		case r == utf8.RuneError && n == 1:
			return out, errInvalidUTF8
		case r < ' ' || r == '"' || r == '\\':
			out = append(out, '\\')
			switch r {
			case '"', '\\':
				out = append(out, byte(r))
			case '\b':
				out = append(out, 'b')
			case '\f':
				out = append(out, 'f')
			case '\n':
				out = append(out, 'n')
			case '\r':
				out = append(out, 'r')
			case '\t':
				out = append(out, 't')
			default:
				out = append(out, 'u')
				out = append(out, "0000"[1+(bits.Len32(uint32(r))-1)/4:]...)
				out = strconv.AppendUint(out, uint64(r), 16)
			}
			in = in[n:]
		default:
			i := indexNeedEscapeInString(in[n:])
			in, out = in[n+i:], append(out, in[:n+i]...)
		}
	}
	out = append(out, '"')
	return out, nil
}

func indexNeedEscapeInString(s string) int {
	for i, r := range s {
		if r < ' ' || r == '\\' || r == '"' || r == utf8.RuneError {
			return i
		}
	}
	return len(s)
}

func (e *Encoder) writeValueSeparator() {
	if e.indented {
		e.writeValueSeparatorIndented()
	} else {
		e.writeValueSeparatorMinimized()
	}
}

func (e *Encoder) writeValueSeparatorMinimized() {
	if e.tokenType != tokenNone && e.tokenType != tokenName && e.tokenType != tokenStartArray {
		e.out = append(e.out, ',')
	}
}

func (e *Encoder) writeValueSeparatorIndented() {
	if e.tokenType != tokenNone && e.tokenType != tokenName && e.tokenType != tokenStartArray {
		e.out = append(e.out, ',')
	}
	if e.tokenType != tokenNone && e.tokenType != tokenName {
		e.out = append(e.out, '\n')
		e.writeIndentation(e.depth)
	}
}

func (e *Encoder) writeIndentation(depth int) {
	for i := 0; i < depth*e.indentLength; i++ {
		e.out = append(e.out, ' ')
	}
}
