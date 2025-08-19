package krpc

import (
	"errors"
	"sync"

	"github.com/iakud/knoll/knet"
	"github.com/iakud/knoll/krpc/knetpb"
	"google.golang.org/protobuf/proto"
)

type wsClient struct {
	hash    uint64
	client  *knet.WSClient
	handler Handler
	locker  sync.RWMutex
	conn    *wsConn
}

func NewWSClient(url string, hash uint64, handler Handler) Client {
	c := &wsClient{handler: handler}
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
		c.handshake(wsconn)
		return
	}

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

func (c *wsClient) Receive(wsconn *knet.WSConn, data []byte) {
	var msg ClientMessage
	if err := Unmarshal(data, &msg); err != nil {
		wsconn.Close()
		return
	}

	if msg.MsgId() < uint16(knetpb.Msg_RESERVED_END) {
		if err := c.handleMsg(wsconn, &msg); err != nil {
			wsconn.Close()
		}
		return
	}

	conn, ok := wsconn.Userdata.(*wsConn)
	if !ok {
		wsconn.Close()
		return
	}

	c.handler.Receive(conn, &msg)
}

func (c *wsClient) handleMsg(wsconn *knet.WSConn, msg *ClientMessage) error {
	switch msg.MsgId() {
	case uint16(knetpb.Msg_HANDSHAKE):
		return c.handleHandshake(wsconn, msg)
	default:
		return errors.New("unknow message")
	}
}

func (c *wsClient) handshake(wsconn *knet.WSConn) {
	var req knetpb.HandshakeRequest
	req.SetHash(c.hash)
	payload, err := proto.Marshal(&req)
	if err != nil {
		wsconn.Close()
		return
	}
	var msg ClientMessage
	msg.SetMsgId(uint16(knetpb.Msg_HANDSHAKE))
	msg.SetPayload(payload)
	data, err := Marshal(&msg)
	if err != nil {
		wsconn.Close()
		return
	}
	wsconn.Send(data)
}

func (c *wsClient) handleHandshake(wsconn *knet.WSConn, msg *ClientMessage) error {
	if wsconn.Userdata != nil {
		return errors.New("already handshake")
	}
	conn := newWSConn(0, wsconn, 0)
	wsconn.Userdata = conn

	c.locker.Lock()
	c.conn = conn
	c.locker.Unlock()

	c.handler.Connect(conn, true)
	return nil
}
