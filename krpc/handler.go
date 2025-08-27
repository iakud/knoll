package krpc

import (
	"errors"

	"github.com/iakud/knoll/krpc/knetpb"
	"google.golang.org/protobuf/proto"
)

func handleServerMsg(conn Conn, m Message, handler Handler) error {
	switch m.Header().MsgId() {
	case uint16(knetpb.Msg_HANDSHAKE):
		return handleServerHandshake(conn, m, handler)
	case uint16(knetpb.Msg_USER_ONLINE):
		return handleServerUserOnline(conn, m, handler)
	default:
		return errors.New("unknow message")
	}
}

func handleServerHandshake(conn Conn, m Message, handler Handler) error {
	var req knetpb.HandshakeRequest
	if err := proto.Unmarshal(m.Payload(), &req); err != nil {
		return err
	}

	if err := replyServerHandshake(conn); err != nil {
		return err
	}
	conn.setHash(req.GetHash())
	handler.Handshake(conn, req.GetHash())
	return nil
}

func handleServerUserOnline(conn Conn, m Message, handler Handler) error {
	var req knetpb.UserOnlineRequest
	if err := proto.Unmarshal(m.Payload(), &req); err != nil {
		return err
	}

	if err := replyUserOnline(conn); err != nil {
		return err
	}
	handler.Handshake(conn, req.GetUid())
	return nil
}

func replyServerHandshake(conn Conn) error {
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

func handleClientMsg(conn Conn, m Message, handler Handler) error {
	switch m.Header().MsgId() {
	case uint16(knetpb.Msg_HANDSHAKE):
		return handleClientHandshake(conn, m, handler)
	default:
		return errors.New("unknow message")
	}
}

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

func handleClientHandshake(conn Conn, m Message, handler Handler) error {
	var reply knetpb.HandshakeReply
	if err := proto.Unmarshal(m.Payload(), &reply); err != nil {
		return err
	}

	conn.setHash(0)
	handler.Handshake(conn, 0)
	return nil
}
