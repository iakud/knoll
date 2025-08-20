package krpc

import (
	"errors"
	"sync"
	"sync/atomic"

	"github.com/iakud/knoll/knet"
	"github.com/iakud/knoll/krpc/knetpb"
	"google.golang.org/protobuf/proto"
)

type tcpServer[T any, M MessagePtr[T]] struct {
	connId  atomic.Uint64
	server  *knet.TCPServer
	handler Handler
	locker  sync.RWMutex
	conns   map[uint64]*tcpConn
}

func NewTCPServer[T any, M MessagePtr[T]](addr string, handler Handler) Server {
	s := &tcpServer[T, M]{
		handler: handler,
		conns:   make(map[uint64]*tcpConn),
	}
	s.server = knet.NewTCPServer(addr, s, knet.StdCodec)
	return s
}

func (s *tcpServer[T, M]) ListenAndServe() error {
	return s.server.ListenAndServe()
}

func (s *tcpServer[T, M]) Close() error {
	return s.server.Close()
}

func (s *tcpServer[T, M]) GetConn(id uint64) (Conn, bool) {
	s.locker.RLock()
	conn, ok := s.conns[id]
	s.locker.RUnlock()
	return conn, ok
}

func (s *tcpServer[T, M]) Connect(tcpconn *knet.TCPConn, connected bool) {
	if connected {
		// FIXME: rouge timeout
		return
	}

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

func (s *tcpServer[T, M]) Receive(tcpconn *knet.TCPConn, data []byte) {
	var m M = new(T)
	if _, err := m.Unmarshal(data); err != nil {
		tcpconn.Close()
		return
	}

	if m.MsgId() < uint16(knetpb.Msg_RESERVED_END) {
		if err := s.handleMsg(tcpconn, m); err != nil {
			tcpconn.Close()
		}
		return
	}

	conn, ok := tcpconn.Userdata.(*tcpConn)
	if !ok {
		tcpconn.Close()
		return
	}

	s.handler.Receive(conn, m)
}

func (s *tcpServer[T, M]) handleMsg(tcpconn *knet.TCPConn, m M) error {
	switch m.MsgId() {
	case uint16(knetpb.Msg_HANDSHAKE):
		return s.handleHandshake(tcpconn, m)
	default:
		return errors.New("unknow message")
	}
}

func (s *tcpServer[T, M]) handleHandshake(tcpconn *knet.TCPConn, m M) error {
	var req knetpb.HandshakeRequest
	if err := proto.Unmarshal(m.Payload(), &req); err != nil {
		return err
	}

	if tcpconn.Userdata != nil {
		return errors.New("already handshake")
	}

	if err := s.handshakeReply(tcpconn); err != nil {
		return err
	}

	connId := s.connId.Add(1)
	conn := newTCPConn(connId, tcpconn, req.GetHash())
	tcpconn.Userdata = conn

	s.locker.Lock()
	s.conns[conn.id] = conn
	s.locker.Unlock()

	s.handler.Connect(conn, true)
	return nil
}

func (s *tcpServer[T, M]) handshakeReply(tcpconn *knet.TCPConn) error {
	var reply knetpb.HandshakeReply
	payload, err := proto.Marshal(&reply)
	if err != nil {
		return err
	}
	var m M = new(T)
	m.SetMsgId(uint16(knetpb.Msg_HANDSHAKE))
	m.SetPayload(payload)

	data := make([]byte, m.Size())
	if _, err := m.Marshal(data); err != nil {
		return err
	}
	return tcpconn.Send(data)
}
