package krpc

import (
	"errors"
	"sync"
	"sync/atomic"

	"github.com/iakud/knoll/knet"
	"github.com/iakud/knoll/krpc/knetpb"
	"google.golang.org/protobuf/proto"
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
		// FIXME: rouge timeout
		return
	}

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

func (s *wsServer) Receive(wsconn *knet.WSConn, data []byte) {
	m := s.newMessage()
	if _, err := m.Unmarshal(data); err != nil {
		wsconn.Close()
		return
	}

	if m.MsgId() < uint16(knetpb.Msg_RESERVED_END) {
		if err := s.handleMsg(wsconn, m); err != nil {
			wsconn.Close()
		}
		return
	}

	conn, ok := wsconn.Userdata.(*wsConn)
	if !ok {
		wsconn.Close()
		return
	}

	s.handler.Receive(conn, m)
}

func (s *wsServer) handleMsg(wsconn *knet.WSConn, m Message) error {
	switch m.MsgId() {
	case uint16(knetpb.Msg_HANDSHAKE):
		return s.handleHandshake(wsconn, m)
	default:
		return errors.New("unknow message")
	}
}

func (s *wsServer) handleHandshake(wsconn *knet.WSConn, m Message) error {
	var req knetpb.HandshakeRequest
	if err := proto.Unmarshal(m.Payload(), &req); err != nil {
		return err
	}

	if wsconn.Userdata != nil {
		return errors.New("already handshake")
	}

	if err := s.handshakeReply(wsconn); err != nil {
		return err
	}

	connId := s.connId.Add(1)
	conn := newWSConn(connId, wsconn, req.GetHash())
	wsconn.Userdata = conn

	s.locker.Lock()
	s.conns[conn.id] = conn
	s.locker.Unlock()

	s.handler.Connect(conn, true)
	return nil
}

func (s *wsServer) handshakeReply(wsconn *knet.WSConn) error {
	var reply knetpb.HandshakeReply
	payload, err := proto.Marshal(&reply)
	if err != nil {
		return err
	}

	m := s.newMessage()
	m.SetMsgId(uint16(knetpb.Msg_HANDSHAKE))
	m.SetPayload(payload)

	data := make([]byte, m.Size())
	if _, err := m.Marshal(data); err != nil {
		return err
	}
	return wsconn.Send(data)
}
