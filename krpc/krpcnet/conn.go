package krpcnet

import (
	"context"
	"net"

	"github.com/iakud/knoll/krpc/knetpb"
	"google.golang.org/protobuf/proto"
)

type Conn interface {
	Id() int64
	Hash() uint64
	Shutdown()
	Close() error
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	NewMsg() Msg
	Send(msg Msg) error
	Request(ctx context.Context, m Msg) (Msg, error)
	Reply(reqId uint32, reply Msg) error

	SetUserdata(userdata any)
	GetUserdata() any
}

func ReplyOK(conn Conn, req Msg) error {
	m := conn.NewMsg()
	m.Header().SetFlagReply()
	m.Header().SetReqId(req.Header().ReqId())
	m.Header().SetMsgId(uint16(knetpb.Msg_OK))
	m.SetConnId(req.ConnId())
	return conn.Send(m)
}

func ReplyError(conn Conn, req Msg, code int32, message string) error {
	var reply knetpb.Error
	reply.SetCode(code)
	reply.SetMessage(message)
	return Reply(conn, req, uint16(knetpb.Msg_ERROR), &reply)
}

func Reply(conn Conn, req Msg, msgId uint16, message proto.Message) error {
	payload, err := proto.Marshal(message)
	if err != nil {
		return err
	}
	m := conn.NewMsg()
	m.Header().SetFlagReply()
	m.Header().SetReqId(req.Header().ReqId())
	m.Header().SetMsgId(msgId)
	m.SetConnId(req.ConnId())
	m.SetPayload(payload)
	return conn.Send(m)
}

func Send(conn Conn, msgId uint16, message proto.Message, connId int64) error {
	payload, err := proto.Marshal(message)
	if err != nil {
		return err
	}
	m := conn.NewMsg()
	m.Header().SetMsgId(msgId)
	m.SetPayload(payload)
	m.SetConnId(connId)
	return conn.Send(m)
}

func SendClientHandshake(conn Conn, hash uint64) error {
	var msg knetpb.ClientHandshake
	msg.SetHash(hash)
	payload, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	m := conn.NewMsg()
	m.Header().SetMsgId(uint16(knetpb.Msg_HANDSHAKE))
	m.SetPayload(payload)
	return conn.Send(m)
}

func SendServerHandshake(conn Conn, hash uint64) error {
	var msg knetpb.ServerHandshake
	msg.SetHash(hash)
	payload, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	m := conn.NewMsg()
	m.Header().SetMsgId(uint16(knetpb.Msg_HANDSHAKE))
	m.SetPayload(payload)
	return conn.Send(m)
}

func SendPing(conn Conn) error {
	var msg knetpb.Ping
	payload, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	m := conn.NewMsg()
	m.Header().SetMsgId(uint16(knetpb.Msg_PING))
	m.SetPayload(payload)
	return conn.Send(m)
}

func SendPong(conn Conn) error {
	var msg knetpb.Pong
	payload, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	m := conn.NewMsg()
	m.Header().SetMsgId(uint16(knetpb.Msg_PONG))
	m.SetPayload(payload)
	return conn.Send(m)
}

func RequestUserOnline(ctx context.Context, conn Conn, req *knetpb.UserOnlineRequest) (*knetpb.UserOnlineReply, error) {
	payload, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}
	m := conn.NewMsg()
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

func SendUserOffline(conn Conn, userId int64, userConnId int64) error {
	var msg knetpb.UserOfflineNotify
	msg.SetUserId(userId)
	msg.SetUserConnId(userConnId)
	payload, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}
	m := conn.NewMsg()
	m.Header().SetMsgId(uint16(knetpb.Msg_USER_OFFLINE_NTF))
	m.SetPayload(payload)
	return conn.Send(m)
}

func RequestKickOut(ctx context.Context, conn Conn, req *knetpb.KickOutRequest) (*knetpb.KickOutReply, error) {
	payload, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}
	m := conn.NewMsg()
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
	m := conn.NewMsg()
	m.Header().SetMsgId(uint16(knetpb.Msg_KICKED_OUT_NTF))
	m.SetPayload(payload)
	return conn.Send(m)
}
