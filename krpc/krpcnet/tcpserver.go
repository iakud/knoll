package krpcnet

import (
	"log/slog"
	"sync"
	"sync/atomic"

	"github.com/iakud/knoll/knet"
	"github.com/iakud/knoll/krpc/knetpb"
	"google.golang.org/protobuf/proto"
)

type tcpServer struct {
	connId  atomic.Uint64
	server  *knet.TCPServer
	handler ServerHandler
	backend bool
	locker  sync.RWMutex
	conns   map[uint64]*tcpConn
}

func NewTCPServer(addr string, handler ServerHandler, backend bool) Server {
	s := &tcpServer{
		handler: handler,
		backend: backend,
		conns:   make(map[uint64]*tcpConn),
	}
	s.server = knet.NewTCPServer(addr, s, knet.StdCodec)
	return s
}

func (s *tcpServer) ListenAndServe() error {
	return s.server.ListenAndServe()
}

func (s *tcpServer) Close() error {
	return s.server.Close()
}

func (s *tcpServer) GetConn(id uint64) (Conn, bool) {
	s.locker.RLock()
	conn, ok := s.conns[id]
	s.locker.RUnlock()
	return conn, ok
}

func (s *tcpServer) Connect(tcpconn *knet.TCPConn, connected bool) {
	if connected {
		connId := s.connId.Add(1)
		conn := newTCPConn(connId, tcpconn, s.backend)
		tcpconn.Userdata = conn

		s.locker.Lock()
		s.conns[conn.id] = conn
		s.locker.Unlock()

		s.handler.Connect(conn, true)
	} else {
		if tcpconn.Userdata == nil {
			return
		}
		conn, ok := tcpconn.Userdata.(*tcpConn)
		if !ok {
			return
		}
		s.locker.Lock()
		delete(s.conns, conn.id)
		s.locker.Unlock()

		s.handler.Connect(conn, false)
	}
}

func (s *tcpServer) Receive(tcpconn *knet.TCPConn, data []byte) {
	if tcpconn.Userdata == nil {
		return
	}
	conn, ok := tcpconn.Userdata.(*tcpConn)
	if !ok {
		tcpconn.Close()
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
				slog.Info("krpcnet: tcpserver handle reply", "error", err)
			}
			return
		}
	*/

	if err := s.handleMessage(conn, m); err != nil {
		slog.Info("krpcnet: tcpserver handle msg", "msgId", m.Header().MsgId(), "error", err)
	}
}

func (s *tcpServer) handleMessage(conn *tcpConn, m Msg) error {
	switch m.Header().MsgId() {
	case uint16(knetpb.Msg_HANDSHAKE):
		return s.handleHandshake(conn, m)
	case uint16(knetpb.Msg_USER_ONLINE):
		return s.handleUserOnline(conn, m)
	case uint16(knetpb.Msg_KICK_OUT):
		return s.handleKickOut(conn, m)
	}
	return s.handler.Receive(conn, m)
}

func (s *tcpServer) handleHandshake(conn *tcpConn, m Msg) error {
	var msg knetpb.ClientHandshake
	if err := proto.Unmarshal(m.Payload(), &msg); err != nil {
		return err
	}
	conn.hash = msg.GetHash()
	return s.handler.Handshake(conn, &msg)
}

func (s *tcpServer) handleUserOnline(conn *tcpConn, m Msg) error {
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

func (s *tcpServer) handleKickOut(conn *tcpConn, m Msg) error {
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
