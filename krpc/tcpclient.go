package krpc

import (
	"errors"
	"sync"

	"github.com/iakud/knoll/knet"
	"github.com/iakud/knoll/krpc/knetpb"
	"google.golang.org/protobuf/proto"
)

type tcpClient struct {
	hash    uint64
	client  *knet.TCPClient
	handler Handler
	locker  sync.RWMutex
	conn    *tcpConn
}

func NewTCPClient(addr string, hash uint64, handler Handler) Client {
	c := &tcpClient{handler: handler}
	c.client = knet.NewTCPClient(addr, c, knet.StdCodec)
	c.client.EnableRetry()
	return c
}

func (c *tcpClient) DialAndServe() error {
	return c.client.DialAndServe()
}

func (c *tcpClient) Close() error {
	return c.client.Close()
}

func (c *tcpClient) GetConn() (Conn, bool) {
	c.locker.RLock()
	conn := c.conn
	c.locker.RUnlock()
	return conn, conn != nil
}

func (c *tcpClient) Connect(tcpconn *knet.TCPConn, connected bool) {
	if connected {
		c.handshake(tcpconn)
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

func (c *tcpClient) Receive(tcpconn *knet.TCPConn, data []byte) {
	var msg InternalMsg
	if err := Unmarshal(data, &msg); err != nil {
		tcpconn.Close()
		return
	}

	if msg.MsgId() < uint16(knetpb.Msg_RESERVED_END) {
		if err := c.handleMsg(tcpconn, &msg); err != nil {
			tcpconn.Close()
		}
		return
	}

	conn, ok := tcpconn.Userdata.(*tcpConn)
	if !ok {
		tcpconn.Close()
		return
	}

	c.handler.Receive(conn, &msg)
}

func (c *tcpClient) handleMsg(tcpconn *knet.TCPConn, msg *InternalMsg) error {
	switch msg.MsgId() {
	case uint16(knetpb.Msg_HANDSHAKE):
		return c.handleHandshake(tcpconn, msg)
	default:
		return errors.New("unknow message")
	}
}

func (c *tcpClient) handshake(tcpconn *knet.TCPConn) {
	var req knetpb.HandshakeRequest
	req.SetHash(c.hash)
	payload, err := proto.Marshal(&req)
	if err != nil {
		tcpconn.Close()
		return
	}
	var msg InternalMsg
	msg.SetMsgId(uint16(knetpb.Msg_HANDSHAKE))
	msg.SetPayload(payload)
	data, err := Marshal(&msg)
	if err != nil {
		tcpconn.Close()
		return
	}
	tcpconn.Send(data)
}

func (c *tcpClient) handleHandshake(tcpconn *knet.TCPConn, msg *InternalMsg) error {
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
