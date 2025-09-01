package krpc

import (
	"context"

	"github.com/iakud/knoll/krpc/knetpb"
	"google.golang.org/protobuf/proto"
)

func SendClientHandshake(conn Conn, hash uint64) error {
	var msg knetpb.ClientHandshake
	msg.SetHash(hash)
	payload, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	m := conn.NewMessage()
	m.Header().SetMsgId(uint16(knetpb.Msg_HANDSHAKE))
	m.SetPayload(payload)
	return conn.Send(m)
}

func SendServerHandshake(conn Conn, hash uint64) error {
	var msg knetpb.ServerHandshake
	payload, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	m := conn.NewMessage()
	m.Header().SetMsgId(uint16(knetpb.Msg_HANDSHAKE))
	m.SetPayload(payload)
	return conn.Send(m)
}

func RequestUserOnline(ctx context.Context, conn Conn, req *knetpb.UserOnlineRequest) (*knetpb.UserOnlineReply, error) {
	payload, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}
	m := conn.NewMessage()
	m.Header().SetMsgId(uint16(knetpb.Msg_USER_ONLINE))
	m.SetPayload(payload)
	m, err = conn.Request(ctx, m)
	if err != nil {
		return nil, err
	}
	var reply knetpb.UserOnlineReply
	if err := proto.Unmarshal(m.Payload(), &reply); err != nil {
		return nil, err
	}
	return &reply, nil
}

func SendUserOffline(conn Conn, userId uint64) error {
	var msg knetpb.UserOfflineNotify
	msg.SetUserId(userId)
	payload, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	m := conn.NewMessage()
	m.Header().SetMsgId(uint16(knetpb.Msg_USER_OFFLINE_NTF))
	m.SetPayload(payload)
	return conn.Send(m)
}

func RequestKickOut(ctx context.Context, conn Conn, req *knetpb.KickOutRequest) (*knetpb.KickOutReply, error) {
	payload, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}
	m := conn.NewMessage()
	m.Header().SetMsgId(uint16(knetpb.Msg_KICK_OUT))
	m.SetPayload(payload)
	m, err = conn.Request(ctx, m)
	if err != nil {
		return nil, err
	}
	var reply knetpb.KickOutReply
	if err := proto.Unmarshal(m.Payload(), &reply); err != nil {
		return nil, err
	}
	return &reply, nil
}

func SendKickedOut(conn Conn, message string) error {
	var msg knetpb.KickedOutNotify
	msg.SetMessage(message)
	payload, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	m := conn.NewMessage()
	m.Header().SetMsgId(uint16(knetpb.Msg_KICKED_OUT_NTF))
	m.SetPayload(payload)
	return conn.Send(m)
}
