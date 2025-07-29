package krpc

import (
	"encoding/binary"
	"errors"
	"slices"
)

const (
	InternalMsgSrcSize = 8
	InternalMsgDstSize = 8
)

type InternalMsg struct {
	Header
	src     uint64
	dst     uint64
	payload []byte
}

func (m *InternalMsg) Src() uint64 {
	return m.src
}

func (m *InternalMsg) SetSrc(src uint64) {
	m.src = src
}

func (m *InternalMsg) Dst() uint64 {
	return m.dst
}

func (m *InternalMsg) SetDst(dst uint64) {
	m.dst = dst
}

func (m *InternalMsg) Payload() []byte {
	return m.payload
}

func (m *InternalMsg) SetPayload(payload []byte) {
	m.payload = payload
}

func (m *InternalMsg) Size() int {
	return HeaderSize + InternalMsgSrcSize + InternalMsgDstSize + len(m.payload)
}

func (m *InternalMsg) Marshal(buf []byte) (int, error) {
	if cap(buf) < HeaderSize+InternalMsgSrcSize+InternalMsgDstSize+len(m.payload) {
		return 0, errors.New("buffer too small")
	}
	n, err := m.Header.Marshal(buf)
	if err != nil {
		return n, err
	}

	binary.BigEndian.PutUint64(buf[n:], m.src)
	n += InternalMsgSrcSize
	binary.BigEndian.PutUint64(buf[n:], m.dst)
	n += InternalMsgDstSize

	n += copy(buf[n:], m.payload)
	return n, nil
}

func (m *InternalMsg) Unmarshal(buf []byte) (int, error) {
	if len(buf) < HeaderSize+InternalMsgSrcSize+InternalMsgDstSize {
		return 0, errors.New("buffer too small")
	}
	n, err := m.Header.Unmarshal(buf)
	if err != nil {
		return n, err
	}

	m.src = binary.BigEndian.Uint64(buf[n:])
	n += InternalMsgSrcSize
	m.dst = binary.BigEndian.Uint64(buf[n:])
	n += InternalMsgDstSize

	m.payload = slices.Clone(buf[n:])
	n += len(m.payload)
	return n, nil
}
