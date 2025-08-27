package krpc

import (
	"sync"

	"github.com/iakud/knoll/knet"
	"github.com/iakud/knoll/krpc/knetpb"
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
		conn := newTCPConn(0, tcpconn, c.newMessage)
		tcpconn.Userdata = conn

		c.locker.Lock()
		c.conn = conn
		c.locker.Unlock()
		if err := requestHandshake(conn, c.hash); err != nil {
			tcpconn.Close()
			return
		}
	} else {
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
}

func (c *tcpClient) Receive(tcpconn *knet.TCPConn, data []byte) {
	if tcpconn.Userdata == nil {
		return
	}
	conn, ok := tcpconn.Userdata.(*tcpConn)
	if !ok {
		tcpconn.Close()
		return
	}

	m := c.newMessage()
	if _, err := m.Unmarshal(data); err != nil {
		tcpconn.Close()
		return
	}

	if m.Header().MsgId() < uint16(knetpb.Msg_RESERVED_END) {
		if err := handleClientMsg(conn, m, c.handler); err != nil {
			tcpconn.Close()
		}
		return
	}

	c.handler.Receive(conn, m)
}
