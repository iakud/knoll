package krpc

import (
	"errors"
	"sync"

	"github.com/iakud/knoll/knet"
	"github.com/iakud/knoll/krpc/knetpb"
	"google.golang.org/protobuf/proto"
)

type tcpClient[T any, M MessagePtr[T]] struct {
	hash    uint64
	client  *knet.TCPClient
	handler Handler
	locker  sync.RWMutex
	conn    *tcpConn
}

func NewTCPClient[T any, M MessagePtr[T]](addr string, hash uint64, handler Handler) Client {
	c := &tcpClient[T, M]{handler: handler}
	c.client = knet.NewTCPClient(addr, c, knet.StdCodec)
	c.client.EnableRetry()
	return c
}

func (c *tcpClient[T, M]) DialAndServe() error {
	return c.client.DialAndServe()
}

func (c *tcpClient[T, M]) Close() error {
	return c.client.Close()
}

func (c *tcpClient[T, M]) GetConn() (Conn, bool) {
	c.locker.RLock()
	conn := c.conn
	c.locker.RUnlock()
	return conn, conn != nil
}

func (c *tcpClient[T, M]) Connect(tcpconn *knet.TCPConn, connected bool) {
	if connected {
		if err := c.handshake(tcpconn); err != nil {
			tcpconn.Close()
		}
		return
	}

	if tcpconn.Userdata == nil {
		return
	}
	conn, ok := tcpconn.Userdata.(*tcpConn)
	if !ok {
		return
	}
	c.locker.Lock()
	c.conn = nil
	c.locker.Unlock()

	c.handler.Connect(conn, false)
}

func (c *tcpClient[T, M]) Receive(tcpconn *knet.TCPConn, data []byte) {
	var m M = new(T)

	if _, err := m.Unmarshal(data); err != nil {
		tcpconn.Close()
		return
	}

	if m.MsgId() < uint16(knetpb.Msg_RESERVED_END) {
		if err := c.handleMsg(tcpconn, m); err != nil {
			tcpconn.Close()
		}
		return
	}

	conn, ok := tcpconn.Userdata.(*tcpConn)
	if !ok {
		tcpconn.Close()
		return
	}

	c.handler.Receive(conn, m)
}

func (c *tcpClient[T, M]) handleMsg(tcpconn *knet.TCPConn, m M) error {
	switch m.MsgId() {
	case uint16(knetpb.Msg_HANDSHAKE):
		return c.handleHandshake(tcpconn, m)
	default:
		return errors.New("unknow message")
	}
}

func (c *tcpClient[T, M]) handshake(tcpconn *knet.TCPConn) error {
	var req knetpb.HandshakeRequest
	req.SetHash(c.hash)
	payload, err := proto.Marshal(&req)
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

func (c *tcpClient[T, M]) handleHandshake(tcpconn *knet.TCPConn, m M) error {
	if tcpconn.Userdata != nil {
		return errors.New("already handshake")
	}
	conn := newTCPConn(0, tcpconn, 0)
	tcpconn.Userdata = conn

	c.locker.Lock()
	c.conn = conn
	c.locker.Unlock()

	c.handler.Connect(conn, true)
	return nil
}
