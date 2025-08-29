package krpc

import (
	"sync"

	"github.com/iakud/knoll/knet"
	"github.com/iakud/knoll/krpc/knetpb"
)

type wsClient struct {
	hash       uint64
	client     *knet.WSClient
	handler    Handler
	newMessage func() Message
	locker     sync.RWMutex
	conn       *wsConn
}

func NewWSClient(url string, hash uint64, handler Handler, newMessage func() Message) Client {
	c := &wsClient{
		hash:       hash,
		handler:    handler,
		newMessage: newMessage,
	}
	c.client = knet.NewWSClient(url, c)
	c.client.EnableRetry()
	return c
}

func (c *wsClient) DialAndServe() error {
	return c.client.DialAndServe()
}

func (c *wsClient) Close() error {
	return c.client.Close()
}

func (c *wsClient) GetConn() (Conn, bool) {
	c.locker.RLock()
	conn := c.conn
	c.locker.RUnlock()
	return conn, conn != nil
}

func (c *wsClient) Connect(wsconn *knet.WSConn, connected bool) {
	if connected {
		conn := newWSConn(0, wsconn, c.newMessage)
		wsconn.Userdata = conn

		c.locker.Lock()
		c.conn = conn
		c.locker.Unlock()

		if err := requestHandshake(conn, c.hash); err != nil {
			wsconn.Close()
			return
		}
	} else {
		if wsconn.Userdata == nil {
			return
		}
		conn, ok := wsconn.Userdata.(*wsConn)
		if !ok {
			return
		}
		c.locker.Lock()
		c.conn = nil
		c.locker.Unlock()

		c.handler.Connect(conn, false)
	}
}

func (c *wsClient) Receive(wsconn *knet.WSConn, data []byte) {
	if wsconn.Userdata == nil {
		return
	}
	conn, ok := wsconn.Userdata.(*wsConn)
	if !ok {
		wsconn.Close()
		return
	}

	m := c.newMessage()
	if _, err := m.Unmarshal(data); err != nil {
		wsconn.Close()
		return
	}

	if m.Header().MsgId() < uint16(knetpb.Msg_RESERVED_END) {
		if err := handleClientMsg(conn, m, c.handler); err != nil {
			wsconn.Close()
		}
		return
	}

	c.handler.Receive(conn, m)
}
