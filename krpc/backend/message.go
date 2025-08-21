package backend

import (
	"encoding/binary"
	"errors"
	"slices"

	"github.com/iakud/knoll/krpc"
)

const (
	MessageConnIdSize = 8
	MessageUserIdSize = 8
)

var _ krpc.Message = (*Message)(nil)

type Message struct {
	header  krpc.Header
	connId  uint64
	userId  uint64
	payload []byte
}

func (m *Message) Header() *krpc.Header {
	return &m.header
}

func (m *Message) ConnId() uint64 {
	return m.connId
}

func (m *Message) SetConnId(connId uint64) {
	m.connId = connId
}

func (m *Message) UserId() uint64 {
	return m.userId
}

func (m *Message) SetUserId(userId uint64) {
	m.userId = userId
}

func (m *Message) Payload() []byte {
	return m.payload
}

func (m *Message) SetPayload(payload []byte) {
	m.payload = payload
}

func (m *Message) Size() int {
	return krpc.HeaderSize + MessageConnIdSize + MessageUserIdSize + len(m.payload)
}

func (m *Message) Marshal(buf []byte) (int, error) {
	if len(buf) < krpc.HeaderSize+MessageConnIdSize+MessageUserIdSize+len(m.payload) {
		return 0, errors.New("buffer too small")
	}
	n, err := m.header.Marshal(buf)
	if err != nil {
		return n, err
	}

	binary.BigEndian.PutUint64(buf[n:], m.connId)
	n += MessageConnIdSize
	binary.BigEndian.PutUint64(buf[n:], m.userId)
	n += MessageUserIdSize

	n += copy(buf[n:], m.payload)
	return n, nil
}

func (m *Message) Unmarshal(buf []byte) (int, error) {
	if len(buf) < krpc.HeaderSize+MessageConnIdSize+MessageUserIdSize {
		return 0, errors.New("buffer too small")
	}
	n, err := m.header.Unmarshal(buf)
	if err != nil {
		return n, err
	}

	m.connId = binary.BigEndian.Uint64(buf[n:])
	n += MessageConnIdSize
	m.userId = binary.BigEndian.Uint64(buf[n:])
	n += MessageUserIdSize

	m.payload = slices.Clone(buf[n:])
	n += len(m.payload)
	return n, nil
}
