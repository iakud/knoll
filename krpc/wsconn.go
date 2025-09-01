package krpc

import (
	"context"
	"net"

	"github.com/iakud/knoll/knet"
)

type wsConn struct {
	id         uint64
	wsconn     *knet.WSConn
	hash       uint64
	rt         *roundTrip
	userdata   any
	newMessage func() Message
}

func newWSConn(id uint64, wsconn *knet.WSConn, newMessage func() Message) *wsConn {
	c := &wsConn{
		id:         id,
		wsconn:     wsconn,
		rt:         newRoundTrip(),
		newMessage: newMessage,
	}
	return c
}

func (c *wsConn) LocalAddr() net.Addr {
	return c.wsconn.LocalAddr()
}

func (c *wsConn) RemoteAddr() net.Addr {
	return c.wsconn.RemoteAddr()
}

func (s *wsConn) Id() uint64 {
	return s.id
}

func (c *wsConn) setHash(hash uint64) {
	c.hash = hash
}

func (c *wsConn) Hash() uint64 {
	return c.hash
}

func (c *wsConn) Request(ctx context.Context, m Message) (Message, error) {
	return c.rt.request(ctx, c, m)
}

func (c *wsConn) NewMessage() Message {
	return c.newMessage()
}

func (c *wsConn) Close() error {
	return c.wsconn.Close()
}

func (c *wsConn) Send(m Message) error {
	data := make([]byte, m.Size())
	if _, err := m.Marshal(data); err != nil {
		return err
	}
	return c.wsconn.Send(data)
}

func (c *wsConn) Reply(reqId uint32, reply Message) error {
	reply.Header().setFlagReply()
	reply.Header().setReqId(reqId)
	return c.Send(reply)
}

func (c *wsConn) SetUserdata(userdata any) {
	c.userdata = userdata
}

func (c *wsConn) GetUserdata() any {
	return c.userdata
}
