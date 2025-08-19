package krpc

import (
	"encoding/binary"
	"errors"
	"slices"

	"google.golang.org/protobuf/proto"
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
	Marshal(buf []byte) (int, error)
	Unmarshal(buf []byte) (int, error)
}

func Marshal(m Message) ([]byte, error) {
	buf := make([]byte, m.Size())
	if _, err := m.Marshal(buf); err != nil {
		return nil, err
	}
	return buf, nil
}

func Unmarshal(buf []byte, m Message) error {
	if _, err := m.Unmarshal(buf); err != nil {
		return err
	}
	return nil
}

type ClientMessage struct {
	Header
	payload []byte
}

func (m *ClientMessage) ConnId() uint64 {
	return 0
}

func (m *ClientMessage) SetConnId(connId uint64) {
}

func (m *ClientMessage) UserId() uint64 {
	return 0
}

func (m *ClientMessage) SetUserId(userId uint64) {
}

func (m *ClientMessage) Payload() []byte {
	return m.payload
}

func (m *ClientMessage) SetPayload(payload []byte) {
	m.payload = payload
}

func (m *ClientMessage) Size() int {
	return HeaderSize + len(m.payload)
}

func (m *ClientMessage) Marshal(buf []byte) (int, error) {
	n, err := m.Header.Marshal(buf)
	if err != nil {
		return n, err
	}

	n += copy(buf[n:], m.payload)
	return n, nil
}

func (m *ClientMessage) Unmarshal(buf []byte) (int, error) {
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

type BackendMessage struct {
	Header
	connId  uint64
	userId  uint64
	payload []byte
}

func (m *BackendMessage) ConnId() uint64 {
	return m.connId
}

func (m *BackendMessage) SetConnId(connId uint64) {
	m.connId = connId
}

func (m *BackendMessage) UserId() uint64 {
	return m.userId
}

func (m *BackendMessage) SetUserId(userId uint64) {
	m.userId = userId
}

func (m *BackendMessage) Payload() []byte {
	return m.payload
}

func (m *BackendMessage) SetPayload(payload []byte) {
	m.payload = payload
}

func (m *BackendMessage) Size() int {
	return HeaderSize + MessageConnIdSize + MessageUserIdSize + len(m.payload)
}

func (m *BackendMessage) Marshal(buf []byte) (int, error) {
	if cap(buf) < HeaderSize+MessageConnIdSize+MessageUserIdSize+len(m.payload) {
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

func (m *BackendMessage) Unmarshal(buf []byte) (int, error) {
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

func BuildClientMessage(msgId uint16, payload proto.Message) (*ClientMessage, error) {
	data, err := proto.Marshal(payload)
	if err != nil {
		return nil, err
	}
	m := new(ClientMessage)
	m.SetMsgId(msgId)
	m.SetPayload(data)
	return m, nil
}

func BuildBackendMessage(msgId uint16, payload proto.Message, connId uint64, userId uint64) (*BackendMessage, error) {
	data, err := proto.Marshal(payload)
	if err != nil {
		return nil, err
	}
	m := new(BackendMessage)
	m.SetMsgId(msgId)
	m.SetConnId(connId)
	m.SetUserId(userId)
	m.SetPayload(data)
	return m, nil
}
