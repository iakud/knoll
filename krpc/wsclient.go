package krpc

import (
	"errors"
	"sync"

	"github.com/iakud/knoll/knet"
	"github.com/iakud/knoll/krpc/knetpb"
	"google.golang.org/protobuf/proto"
)

type wsClient struct {
	client  *knet.WSClient
	handler ClientHandler
	backend bool
	locker  sync.RWMutex
	conn    *wsConn
}

func NewWSClient(url string, handler ClientHandler, backend bool) Client {
	c := &wsClient{
		handler: handler,
		backend: backend,
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
		conn := newWSConn(0, wsconn, c.backend)
		wsconn.Userdata = conn

		c.locker.Lock()
		c.conn = conn
		c.locker.Unlock()

		c.handler.Connect(conn, true)
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

	m := NewMsg(c.backend)
	if _, err := m.Unmarshal(data); err != nil {
		conn.Close()
		return
	}

	if m.Header().FlagReply() && conn.rt != nil {
		if err := conn.rt.handleReply(m); err != nil {
			conn.Close()
		}
		return
	}

	if m.Header().MsgId() < uint16(knetpb.Msg_RESERVED_END) {
		if err := c.handleMessage(conn, m); err != nil {
			conn.Close()
		}
		return
	}

	c.handler.Receive(conn, m)
}

func (c *wsClient) handleMessage(conn *wsConn, m Msg) error {
	switch m.Header().MsgId() {
	case uint16(knetpb.Msg_HANDSHAKE):
		return c.handleHandshake(conn, m)
	case uint16(knetpb.Msg_USER_OFFLINE_NTF):
		return c.handleUserOffline(conn, m)
	case uint16(knetpb.Msg_KICKED_OUT_NTF):
		return c.handleKickedOut(conn, m)
	default:
		return errors.New("unknow message")
	}
}

func (c *wsClient) handleHandshake(conn *wsConn, m Msg) error {
	var msg knetpb.ServerHandshake
	if err := proto.Unmarshal(m.Payload(), &msg); err != nil {
		return err
	}
	conn.hash = msg.GetHash()
	return c.handler.Handshake(conn, &msg)
}

func (c *wsClient) handleUserOffline(conn *wsConn, m Msg) error {
	var reply knetpb.UserOfflineNotify
	if err := proto.Unmarshal(m.Payload(), &reply); err != nil {
		return err
	}
	return c.handler.UserOffline(conn, &reply)
}

func (c *wsClient) handleKickedOut(conn *wsConn, m Msg) error {
	var reply knetpb.KickedOutNotify
	if err := proto.Unmarshal(m.Payload(), &reply); err != nil {
		return err
	}
	return c.handler.KickedOut(conn, &reply)
}
