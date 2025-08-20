package krpc

import (
	"errors"
	"sync"

	"github.com/iakud/knoll/knet"
	"github.com/iakud/knoll/krpc/knetpb"
	"google.golang.org/protobuf/proto"
)

type wsClient[T any, M MessagePtr[T]] struct {
	hash    uint64
	client  *knet.WSClient
	handler Handler
	locker  sync.RWMutex
	conn    *wsConn
}

func NewWSClient[T any, M MessagePtr[T]](url string, hash uint64, handler Handler) Client {
	c := &wsClient[T, M]{handler: handler}
	c.client = knet.NewWSClient(url, c)
	c.client.EnableRetry()
	return c
}

func (c *wsClient[T, M]) DialAndServe() error {
	return c.client.DialAndServe()
}

func (c *wsClient[T, M]) Close() error {
	return c.client.Close()
}

func (c *wsClient[T, M]) GetConn() (Conn, bool) {
	c.locker.RLock()
	conn := c.conn
	c.locker.RUnlock()
	return conn, conn != nil
}

func (c *wsClient[T, M]) Connect(wsconn *knet.WSConn, connected bool) {
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

func (c *wsClient[T, M]) Receive(wsconn *knet.WSConn, data []byte) {
	var m M = new(T)
	if _, err := m.Unmarshal(data); err != nil {
		wsconn.Close()
		return
	}

	if m.MsgId() < uint16(knetpb.Msg_RESERVED_END) {
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

func (c *wsClient[T, M]) handleMsg(wsconn *knet.WSConn, m M) error {
	switch m.MsgId() {
	case uint16(knetpb.Msg_HANDSHAKE):
		return c.handleHandshake(wsconn, m)
	default:
		return errors.New("unknow message")
	}
}

func (c *wsClient[T, M]) handshake(wsconn *knet.WSConn) error {
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
	return wsconn.Send(data)
}

func (c *wsClient[T, M]) handleHandshake(wsconn *knet.WSConn, m M) error {
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
