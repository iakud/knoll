package krpc

import (
	"sync"
	"sync/atomic"

	"github.com/iakud/knoll/knet"
	"github.com/iakud/knoll/krpc/knetpb"
)

type wsServer struct {
	connId     atomic.Uint64
	server     *knet.WSServer
	handler    Handler
	newMessage func() Message
	locker     sync.RWMutex
	conns      map[uint64]*wsConn
}

func NewWSServer(addr string, handler Handler, newMessage func() Message) Server {
	s := &wsServer{
		handler:    handler,
		newMessage: newMessage,
		conns:      make(map[uint64]*wsConn),
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
		conn := newWSConn(connId, wsconn, s.newMessage)
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

	m := s.newMessage()
	if _, err := m.Unmarshal(data); err != nil {
		wsconn.Close()
		return
	}

	if m.Header().MsgId() < uint16(knetpb.Msg_RESERVED_END) {
		if err := handleServerMsg(conn, m, s.handler); err != nil {
			wsconn.Close()
		}
		return
	}

	s.handler.Receive(conn, m)
}
