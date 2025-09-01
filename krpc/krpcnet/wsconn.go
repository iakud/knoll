package krpcnet

import (
	"context"
	"net"

	"github.com/iakud/knoll/knet"
)

type wsConn struct {
	id       uint64
	wsconn   *knet.WSConn
	backend  bool
	rt       *roundTrip
	hash     uint64
	userdata any
}

func newWSConn(id uint64, wsconn *knet.WSConn, backend bool) *wsConn {
	c := &wsConn{
		id:      id,
		wsconn:  wsconn,
		backend: backend,
		rt:      newRoundTrip(),
	}
	return c
}

func (s *wsConn) Id() uint64 {
	return s.id
}

func (c *wsConn) Hash() uint64 {
	return c.hash
}

func (c *wsConn) Close() error {
	return c.wsconn.Close()
}

func (c *wsConn) LocalAddr() net.Addr {
	return c.wsconn.LocalAddr()
}

func (c *wsConn) RemoteAddr() net.Addr {
	return c.wsconn.RemoteAddr()
}

func (c *wsConn) NewMsg() Msg {
	return NewMsg(c.backend)
}

func (c *wsConn) Send(m Msg) error {
	data := make([]byte, m.Size())
	if _, err := m.Marshal(data); err != nil {
		return err
	}
	return c.wsconn.Send(data)
}

func (c *wsConn) Request(ctx context.Context, m Msg) (Msg, error) {
	return c.rt.request(ctx, c, m)
}

func (c *wsConn) Reply(reqId uint32, reply Msg) error {
	reply.Header().SetFlagReply()
	reply.Header().SetReqId(reqId)
	return c.Send(reply)
}

func (c *wsConn) SetUserdata(userdata any) {
	c.userdata = userdata
}

func (c *wsConn) GetUserdata() any {
	return c.userdata
}
