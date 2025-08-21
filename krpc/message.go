package krpc

import (
	"encoding/binary"
	"errors"
	"slices"
)

type Message interface {
	FlagRequest() bool
	setFlagRequest()
	FlagReply() bool
	setFlagReply()
	MsgId() uint16
	SetMsgId(msgId uint16)
	ReqId() uint32
	setReqId(reqId uint32)
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
	Header
	payload []byte
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
	n, err := m.Header.Marshal(buf)
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
	n, err := m.Header.Unmarshal(buf)
	if err != nil {
		return n, err
	}

	m.payload = slices.Clone(buf[n:])
	n += len(m.payload)
	return n, nil
}

const (
	MessageConnIdSize = 8
	MessageUserIdSize = 8
)

type BMessage struct {
	Header
	connId  uint64
	userId  uint64
	payload []byte
}

func (m *BMessage) ConnId() uint64 {
	return m.connId
}

func (m *BMessage) SetConnId(connId uint64) {
	m.connId = connId
}

func (m *BMessage) UserId() uint64 {
	return m.userId
}

func (m *BMessage) SetUserId(userId uint64) {
	m.userId = userId
}

func (m *BMessage) Payload() []byte {
	return m.payload
}

func (m *BMessage) SetPayload(payload []byte) {
	m.payload = payload
}

func (m *BMessage) Size() int {
	return HeaderSize + MessageConnIdSize + MessageUserIdSize + len(m.payload)
}

func (m *BMessage) Marshal(buf []byte) (int, error) {
	if len(buf) < HeaderSize+MessageConnIdSize+MessageUserIdSize+len(m.payload) {
		return 0, errors.New("buffer too small")
	}
	n, err := m.Header.Marshal(buf)
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

func (m *BMessage) Unmarshal(buf []byte) (int, error) {
	if len(buf) < HeaderSize+MessageConnIdSize+MessageUserIdSize {
		return 0, errors.New("buffer too small")
	}
	n, err := m.Header.Unmarshal(buf)
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

func BuildCMessage(msgId uint16, payload []byte) Message {
	m := NewCMessage()
	m.SetMsgId(msgId)
	m.SetPayload(payload)
	return m
}

func BuildBMessage(msgId uint16, payload []byte) Message {
	m := NewBMessage()
	m.SetMsgId(msgId)
	m.SetPayload(payload)
	return m
}
