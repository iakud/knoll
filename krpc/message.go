package krpc

import (
	"errors"
	"slices"
)

type Message interface {
	Header() *Header
	Size() int
	ConnId() uint64
	SetConnId(connId uint64)
	UserId() uint64
	SetUserId(userId uint64)
	Payload() []byte
	SetPayload(payload []byte)
	Marshal(buf []byte) (int, error)
	Unmarshal(buf []byte) (int, error)
}

type CMessage struct {
	header  Header
	payload []byte
}

func (m *CMessage) Header() *Header {
	return &m.header
}

func (m *CMessage) ConnId() uint64 {
	return 0
}

func (m *CMessage) SetConnId(connId uint64) {
}

func (m *CMessage) UserId() uint64 {
	return 0
}

func (m *CMessage) SetUserId(userId uint64) {
}

func (m *CMessage) Payload() []byte {
	return m.payload
}

func (m *CMessage) SetPayload(payload []byte) {
	m.payload = payload
}

func (m *CMessage) Size() int {
	return HeaderSize + len(m.payload)
}

func (m *CMessage) Marshal(buf []byte) (int, error) {
	if len(buf) < HeaderSize+len(m.payload) {
		return 0, errors.New("buffer too small")
	}
	n, err := m.header.Marshal(buf)
	if err != nil {
		return n, err
	}

	n += copy(buf[n:], m.payload)
	return n, nil
}

func (m *CMessage) Unmarshal(buf []byte) (int, error) {
	if len(buf) < HeaderSize {
		return 0, errors.New("buffer too small")
	}
	n, err := m.header.Unmarshal(buf)
	if err != nil {
		return n, err
	}

	m.payload = slices.Clone(buf[n:])
	n += len(m.payload)
	return n, nil
}
