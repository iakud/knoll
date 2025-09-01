package krpc

import (
	"context"
	"net"

	"github.com/iakud/knoll/knet"
)

type tcpConn struct {
	id       uint64
	tcpconn  *knet.TCPConn
	backend  bool
	rt       *roundTrip
	hash     uint64
	userdata any
}

func newTCPConn(id uint64, tcpconn *knet.TCPConn, backend bool) *tcpConn {
	c := &tcpConn{
		id:      id,
		tcpconn: tcpconn,
		backend: backend,
		rt:      newRoundTrip(),
	}
	return c
}

func (s *tcpConn) Id() uint64 {
	return s.id
}

func (c *tcpConn) Hash() uint64 {
	return c.hash
}

func (c *tcpConn) Close() error {
	return c.tcpconn.Close()
}

func (c *tcpConn) LocalAddr() net.Addr {
	return c.tcpconn.LocalAddr()
}

func (c *tcpConn) RemoteAddr() net.Addr {
	return c.tcpconn.RemoteAddr()
}

func (c *tcpConn) NewMsg() Msg {
	return NewMsg(c.backend)
}

func (c *tcpConn) Send(m Msg) error {
	data := make([]byte, m.Size())
	if _, err := m.Marshal(data); err != nil {
		return err
	}
	return c.tcpconn.Send(data)
}

func (c *tcpConn) Request(ctx context.Context, m Msg) (Msg, error) {
	return c.rt.request(ctx, c, m)
}

func (c *tcpConn) Reply(reqId uint32, reply Msg) error {
	reply.Header().SetFlagReply()
	reply.Header().SetReqId(reqId)
	return c.Send(reply)
}

func (c *tcpConn) SetUserdata(userdata any) {
	c.userdata = userdata
}

func (c *tcpConn) GetUserdata() any {
	return c.userdata
}
