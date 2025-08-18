package knet

import (
	"bufio"
	"encoding/binary"
	"io"
)

type Codec interface {
	Read(r io.Reader) ([]byte, error)
	Write(w io.Writer, b []byte) error
}

type CodecReader interface {
	Read(r io.Reader) ([]byte, error)
}

type CodecWriter interface {
	Write(w io.Writer, b []byte) error
}

type defaultCodec struct{}

func (*defaultCodec) Read(r io.Reader) ([]byte, error) {
	rBuf := bufio.NewReader(r)
	if _, err := rBuf.Peek(1); err != nil {
		return nil, err
	}
	b := make([]byte, rBuf.Buffered())
	if _, err := rBuf.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}

func (*defaultCodec) Write(w io.Writer, b []byte) error {
	if _, err := w.Write(b); err != nil {
		return err
	}
	return nil
}

var DefaultCodec *defaultCodec = &defaultCodec{}

type stdCodec struct{}

func (c *stdCodec) Read(r io.Reader) ([]byte, error) {
	h := make([]byte, 2)
	if _, err := io.ReadFull(r, h); err != nil {
		return nil, err
	}
	n := binary.BigEndian.Uint16(h)
	b := make([]byte, n)
	if _, err := io.ReadFull(r, b); err != nil {
		return nil, err
	}
	return b, nil
}

func (c *stdCodec) Write(w io.Writer, b []byte) error {
	h := make([]byte, 2)
	binary.BigEndian.PutUint16(h, uint16(len(b)))
	if _, err := w.Write(h); err != nil {
		return err
	}
	if _, err := w.Write(b); err != nil {
		return err
	}
	return nil
}

var StdCodec *stdCodec = &stdCodec{}
