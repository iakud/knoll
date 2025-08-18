package krpc

import (
	"errors"
	"sync"
	"sync/atomic"

	"github.com/iakud/knoll/knet"
	"github.com/iakud/knoll/krpc/knetpb"
	"google.golang.org/protobuf/proto"
)

type tcpServer struct {
	connId  atomic.Uint64
	server  *knet.TCPServer
	handler Handler
	locker  sync.RWMutex
	conns   map[uint64]*tcpConn
}

func NewTCPServer(addr string, handler Handler) Server {
	s := &tcpServer{handler: handler}
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

func (s *tcpServer) Receive(tcpconn *knet.TCPConn, data []byte) {
	var msg InternalMsg
	if err := Unmarshal(data, &msg); err != nil {
		tcpconn.Close()
		return
	}

	if msg.MsgId() < uint16(knetpb.Msg_RESERVED_END) {
		if err := s.handleMsg(tcpconn, &msg); err != nil {
			tcpconn.Close()
		}
		return
	}

	conn, ok := tcpconn.Userdata.(*tcpConn)
	if !ok {
		tcpconn.Close()
		return
	}

	s.handler.Receive(conn, &msg)
}

func (s *tcpServer) handleMsg(tcpconn *knet.TCPConn, msg *InternalMsg) error {
	switch msg.MsgId() {
	case uint16(knetpb.Msg_HANDSHAKE):
		return s.handleHandshake(tcpconn, msg)
	default:
		return errors.New("unknow message")
	}
}

func (s *tcpServer) handleHandshake(tcpconn *knet.TCPConn, msg *InternalMsg) error {
	var req knetpb.HandshakeRequest
	if err := proto.Unmarshal(msg.Payload(), &req); err != nil {
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

func (s *tcpServer) handshakeReply(tcpconn *knet.TCPConn) error {
	var reply knetpb.HandshakeReply
	payload, err := proto.Marshal(&reply)
	if err != nil {
		return err
	}
	var msg InternalMsg
	msg.SetMsgId(uint16(knetpb.Msg_HANDSHAKE))
	msg.SetPayload(payload)

	data, err := Marshal(&msg)
	if err != nil {
		return err
	}
	return tcpconn.Send(data)
}
