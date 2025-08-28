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

	if err := replyHandshake(conn); err != nil {
		return err
	}
	conn.setHash(req.GetHash())
	handler.Connect(conn, true)
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
	handler.UserOnline(conn, req.GetUserConnId(), req.GetUserId())
	return nil
}

func handleClientMsg(conn Conn, m Message, handler Handler) error {
	switch m.Header().MsgId() {
	case uint16(knetpb.Msg_HANDSHAKE):
		return handleClientHandshake(conn, m, handler)
	case uint16(knetpb.Msg_USER_ONLINE):
		return handleClientUserOnline(conn, m, handler)
	default:
		return errors.New("unknow message")
	}
}

func handleClientHandshake(conn Conn, m Message, handler Handler) error {
	var reply knetpb.HandshakeReply
	if err := proto.Unmarshal(m.Payload(), &reply); err != nil {
		return err
	}

	conn.setHash(0)
	handler.Connect(conn, true)
	handler.Handshake(conn, 0)
	return nil
}

func handleClientUserOnline(conn Conn, m Message, handler Handler) error {
	var reply knetpb.UserOnlineReply
	if err := proto.Unmarshal(m.Payload(), &reply); err != nil {
		return err
	}
	return nil
}
