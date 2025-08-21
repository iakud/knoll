package krpc

import (
	"errors"
	"sync"

	"github.com/iakud/knoll/knet"
	"github.com/iakud/knoll/krpc/knetpb"
	"google.golang.org/protobuf/proto"
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
		if err := c.handshake(wsconn); err != nil {
			wsconn.Close()
		}
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
	m := c.newMessage()
	if _, err := m.Unmarshal(data); err != nil {
		wsconn.Close()
		return
	}

	if m.Header().MsgId() < uint16(knetpb.Msg_RESERVED_END) {
		if err := c.handleMsg(wsconn, m); err != nil {
			wsconn.Close()
		}
		return
	}

	conn, ok := wsconn.Userdata.(*wsConn)
	if !ok {
		wsconn.Close()
		return
	}

	c.handler.Receive(conn, m)
}

func (c *wsClient) handleMsg(wsconn *knet.WSConn, m Message) error {
	switch m.Header().MsgId() {
	case uint16(knetpb.Msg_HANDSHAKE):
		return c.handleHandshake(wsconn, m)
	default:
		return errors.New("unknow message")
	}
}

func (c *wsClient) handshake(wsconn *knet.WSConn) error {
	var req knetpb.HandshakeRequest
	req.SetHash(c.hash)
	payload, err := proto.Marshal(&req)
	if err != nil {
		return err
	}
	m := c.newMessage()
	m.Header().SetMsgId(uint16(knetpb.Msg_HANDSHAKE))
	m.SetPayload(payload)

	data := make([]byte, m.Size())
	if _, err := m.Marshal(data); err != nil {
		return err
	}
	return wsconn.Send(data)
}

func (c *wsClient) handleHandshake(wsconn *knet.WSConn, m Message) error {
	var reply knetpb.HandshakeReply
	if err := proto.Unmarshal(m.Payload(), &reply); err != nil {
		return err
	}

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
