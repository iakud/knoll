package krpc

import (
	"sync"
	"sync/atomic"

	"github.com/iakud/knoll/knet"
	"github.com/iakud/knoll/krpc/knetpb"
)

type tcpServer struct {
	connId     atomic.Uint64
	server     *knet.TCPServer
	handler    Handler
	newMessage func() Message
	locker     sync.RWMutex
	conns      map[uint64]*tcpConn
}

func NewTCPServer(addr string, handler Handler, newMessage func() Message) Server {
	s := &tcpServer{
		handler:    handler,
		newMessage: newMessage,
		conns:      make(map[uint64]*tcpConn),
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
		conn := newTCPConn(connId, tcpconn, s.newMessage)
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

	m := s.newMessage()
	if _, err := m.Unmarshal(data); err != nil {
		tcpconn.Close()
		return
	}

	if m.Header().MsgId() < uint16(knetpb.Msg_RESERVED_END) {
		if err := handleServerMsg(conn, m, s.handler); err != nil {
			tcpconn.Close()
		}
		return
	}

	s.handler.Receive(conn, m)
}
