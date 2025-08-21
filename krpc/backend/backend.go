package backend

import (
	"sync"

	"github.com/iakud/knoll/krpc"
	"github.com/iakud/knoll/krpc/knetpb"
	"google.golang.org/protobuf/proto"
)

var pool = sync.Pool{New: func() any { return &Message{} }}

func New() krpc.Message {
	return pool.Get().(krpc.Message)
}

func ReplyOK(conn krpc.Conn, reqId uint32, connId uint64) error {
	m := New()
	m.Header().SetMsgId(uint16(knetpb.Msg_OK))
	m.SetConnId(connId)
	return conn.Reply(reqId, m)
}

func ReplyError(conn krpc.Conn, reqId uint32, code int32, msg string, connId uint64) error {
	var reply knetpb.Error
	reply.SetCode(code)
	reply.SetMsg(msg)
	return Reply(conn, reqId, uint16(knetpb.Msg_ERROR), &reply, connId)
}

func Reply(conn krpc.Conn, reqId uint32, msgId uint16, message proto.Message, connId uint64) error {
	payload, err := proto.Marshal(message)
	if err != nil {
		return err
	}
	m := New()
	m.Header().SetMsgId(msgId)
	m.SetPayload(payload)
	m.SetConnId(connId)
	return conn.Reply(reqId, m)
}

func Send(conn krpc.Conn, msgId uint16, message proto.Message, connId uint64) error {
	payload, err := proto.Marshal(message)
	if err != nil {
		return err
	}
	m := New()
	m.Header().SetMsgId(msgId)
	m.SetPayload(payload)
	m.SetConnId(connId)
	return conn.Send(m)
}
