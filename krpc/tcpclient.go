package krpc

import (
	"errors"
	"sync"

	"github.com/iakud/knoll/knet"
	"github.com/iakud/knoll/krpc/knetpb"
	"google.golang.org/protobuf/proto"
)

type tcpClient struct {
	hash       uint64
	client     *knet.TCPClient
	handler    Handler
	newMessage func() Message
	locker     sync.RWMutex
	conn       *tcpConn
}

func NewTCPClient(addr string, hash uint64, handler Handler, newMessage func() Message) Client {
	c := &tcpClient{
		handler:    handler,
		newMessage: newMessage,
	}
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

func (c *tcpClient) Receive(tcpconn *knet.TCPConn, data []byte) {
	m := c.newMessage()
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

func (c *tcpClient) handleMsg(tcpconn *knet.TCPConn, m Message) error {
	switch m.MsgId() {
	case uint16(knetpb.Msg_HANDSHAKE):
		return c.handleHandshake(tcpconn, m)
	default:
		return errors.New("unknow message")
	}
}

func (c *tcpClient) handshake(tcpconn *knet.TCPConn) error {
	var req knetpb.HandshakeRequest
	req.SetHash(c.hash)
	payload, err := proto.Marshal(&req)
	if err != nil {
		return err
	}
	m := c.newMessage()
	m.SetMsgId(uint16(knetpb.Msg_HANDSHAKE))
	m.SetPayload(payload)

	data := make([]byte, m.Size())
	if _, err := m.Marshal(data); err != nil {
		return err
	}
	return tcpconn.Send(data)
}

func (c *tcpClient) handleHandshake(tcpconn *knet.TCPConn, m Message) error {
	var reply knetpb.HandshakeReply
	if err := proto.Unmarshal(m.Payload(), &reply); err != nil {
		return err
	}

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
