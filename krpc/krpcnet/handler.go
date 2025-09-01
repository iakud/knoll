package krpcnet

import (
	"github.com/iakud/knoll/krpc/knetpb"
)

type Server interface {
	ListenAndServe() error
	Close() error
	GetConn(id uint64) (Conn, bool)
}

type ServerHandler interface {
	Connect(conn Conn, connected bool)
	Receive(conn Conn, msg Msg)
	Handshake(conn Conn, msg *knetpb.ClientHandshake) error
	UserOnline(conn Conn, req *knetpb.UserOnlineRequest) (*knetpb.UserOnlineReply, error)
	KickOut(conn Conn, req *knetpb.KickOutRequest) (*knetpb.KickOutReply, error)
}

type Client interface {
	DialAndServe() error
	Close() error
	GetConn() (Conn, bool)
}

type ClientHandler interface {
	Connect(conn Conn, connected bool)
	Receive(conn Conn, msg Msg)
	Handshake(conn Conn, msg *knetpb.ServerHandshake) error
	UserOffline(conn Conn, msg *knetpb.UserOfflineNotify) error
	KickedOut(conn Conn, msg *knetpb.KickedOutNotify) error
}
