package krpc

import (
	"errors"
	"slices"

	"github.com/iakud/knoll/krpc/knetpb"
	"google.golang.org/protobuf/proto"
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

func ReplyOK(conn Conn, req Message) error {
	m := conn.NewMessage()
	m.Header().setFlagReply()
	m.Header().setReqId(req.Header().ReqId())
	m.Header().SetMsgId(uint16(knetpb.Msg_OK))
	m.SetConnId(req.ConnId())
	return conn.Send(m)
}

func ReplyError(conn Conn, req Message, code int32, message string) error {
	var reply knetpb.Error
	reply.SetCode(code)
	reply.SetMessage(message)
	return Reply(conn, req, uint16(knetpb.Msg_ERROR), &reply)
}

func Reply(conn Conn, req Message, msgId uint16, message proto.Message) error {
	payload, err := proto.Marshal(message)
	if err != nil {
		return err
	}
	m := conn.NewMessage()
	m.Header().setFlagReply()
	m.Header().setReqId(req.Header().ReqId())
	m.Header().SetMsgId(msgId)
	m.SetConnId(req.ConnId())
	m.SetPayload(payload)
	return conn.Send(m)
}

func Send(conn Conn, msgId uint16, message proto.Message, connId uint64) error {
	payload, err := proto.Marshal(message)
	if err != nil {
		return err
	}
	m := conn.NewMessage()
	m.Header().SetMsgId(msgId)
	m.SetPayload(payload)
	m.SetConnId(connId)
	return conn.Send(m)
}
