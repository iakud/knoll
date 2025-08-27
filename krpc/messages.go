package krpc

import (
	"github.com/iakud/knoll/krpc/knetpb"
	"google.golang.org/protobuf/proto"
)

func requestHandshake(conn Conn, hash uint64) error {
	var req knetpb.HandshakeRequest
	req.SetHash(hash)
	payload, err := proto.Marshal(&req)
	if err != nil {
		return err
	}
	m := conn.NewMessage()
	m.Header().SetMsgId(uint16(knetpb.Msg_HANDSHAKE))
	m.SetPayload(payload)
	return conn.Send(m)
}

func replyHandshake(conn Conn) error {
	var reply knetpb.HandshakeReply
	payload, err := proto.Marshal(&reply)
	if err != nil {
		return err
	}
	m := conn.NewMessage()
	m.Header().SetMsgId(uint16(knetpb.Msg_HANDSHAKE))
	m.SetPayload(payload)
	return conn.Send(m)
}

func replyUserOnline(conn Conn) error {
	var reply knetpb.UserOnlineReply
	payload, err := proto.Marshal(&reply)
	if err != nil {
		return err
	}
	m := conn.NewMessage()
	m.Header().SetMsgId(uint16(knetpb.Msg_USER_ONLINE))
	m.SetPayload(payload)
	return conn.Send(m)
}
