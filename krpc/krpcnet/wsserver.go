package krpcnet

import (
	"log/slog"
	"sync"
	"sync/atomic"

	"github.com/iakud/knoll/knet"
	"github.com/iakud/knoll/krpc/knetpb"
	"google.golang.org/protobuf/proto"
)

type wsServer struct {
	connId  atomic.Uint64
	server  *knet.WSServer
	handler ServerHandler
	backend bool
	locker  sync.RWMutex
	conns   map[uint64]*wsConn
}

func NewWSServer(addr string, handler ServerHandler, backend bool) Server {
	s := &wsServer{
		handler: handler,
		backend: backend,
		conns:   make(map[uint64]*wsConn),
	}
	s.server = knet.NewWSServer(addr, s)
	return s
}

func (s *wsServer) ListenAndServe() error {
	return s.server.ListenAndServe()
}

func (s *wsServer) Close() error {
	return s.server.Close()
}

func (s *wsServer) GetConn(id uint64) (Conn, bool) {
	s.locker.RLock()
	conn, ok := s.conns[id]
	s.locker.RUnlock()
	return conn, ok
}

func (s *wsServer) Connect(wsconn *knet.WSConn, connected bool) {
	if connected {
		connId := s.connId.Add(1)
		conn := newWSConn(connId, wsconn, s.backend)
		wsconn.Userdata = conn

		s.locker.Lock()
		s.conns[conn.id] = conn
		s.locker.Unlock()

		s.handler.Connect(conn, true)
	} else {
		if wsconn.Userdata == nil {
			return
		}
		conn, ok := wsconn.Userdata.(*wsConn)
		if !ok {
			return
		}
		s.locker.Lock()
		delete(s.conns, conn.id)
		s.locker.Unlock()

		s.handler.Connect(conn, false)
	}
}

func (s *wsServer) Receive(wsconn *knet.WSConn, data []byte) {
	if wsconn.Userdata == nil {
		return
	}
	conn, ok := wsconn.Userdata.(*wsConn)
	if !ok {
		wsconn.Close()
		return
	}

	m := NewMsg(s.backend)
	if _, err := m.Unmarshal(data); err != nil {
		conn.Close()
		return
	}

	/*
		if m.Header().FlagReply() && conn.rt != nil {
			if err := conn.rt.handleReply(m); err != nil {
				slog.Info("krpcnet: wsserver handle reply", "error", err)
			}
			return
		}
	*/

	if err := s.handleMessage(conn, m); err != nil {
		slog.Info("krpcnet: wsserver handle msg", "msgId", m.Header().MsgId(), "error", err)
	}
}

func (s *wsServer) handleMessage(conn *wsConn, m Msg) error {
	switch m.Header().MsgId() {
	case uint16(knetpb.Msg_HANDSHAKE):
		return s.handleHandshake(conn, m)
	case uint16(knetpb.Msg_PING):
		return s.handlePing(conn, m)
	case uint16(knetpb.Msg_PONG):
		return s.handlePong(conn, m)
	case uint16(knetpb.Msg_USER_ONLINE):
		return s.handleUserOnline(conn, m)
	case uint16(knetpb.Msg_KICK_OUT):
		return s.handleKickOut(conn, m)
	}
	return s.handler.Receive(conn, m)
}

func (s *wsServer) handleHandshake(conn *wsConn, m Msg) error {
	var msg knetpb.ClientHandshake
	if err := proto.Unmarshal(m.Payload(), &msg); err != nil {
		return err
	}
	conn.hash = msg.GetHash()
	return s.handler.Handshake(conn, &msg)
}

func (s *wsServer) handlePing(conn *wsConn, m Msg) error {
	return nil
}

func (s *wsServer) handlePong(conn *wsConn, m Msg) error {
	return nil
}

func (s *wsServer) handleUserOnline(conn *wsConn, m Msg) error {
	var req knetpb.UserOnlineRequest
	if err := proto.Unmarshal(m.Payload(), &req); err != nil {
		return err
	}
	reply, err := s.handler.UserOnline(conn, &req)
	if err != nil {
		return err
	}
	return Reply(conn, m, uint16(knetpb.Msg_USER_ONLINE), reply)
}

func (s *wsServer) handleKickOut(conn *wsConn, m Msg) error {
	var req knetpb.KickOutRequest
	if err := proto.Unmarshal(m.Payload(), &req); err != nil {
		return err
	}
	reply, err := s.handler.KickOut(conn, &req)
	if err != nil {
		return err
	}
	return Reply(conn, m, uint16(knetpb.Msg_KICK_OUT), reply)
}
