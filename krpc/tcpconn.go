package krpc

import (
	"github.com/iakud/knoll/knet"
)

type tcpConn struct {
	id       uint64
	tcpconn  *knet.TCPConn
	hash     uint64
	userdata any
}

func newTCPConn(id uint64, tcpconn *knet.TCPConn, hash uint64) *tcpConn {
	return &tcpConn{id: id, tcpconn: tcpconn, hash: hash}
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

func (c *tcpConn) Send(m Message) error {
	data := make([]byte, m.Size())
	if _, err := m.Marshal(data); err != nil {
		return err
	}
	return c.tcpconn.Send(data)
}

func (c *tcpConn) Reply(reqId uint32, reply Message) error {
	reply.setFlagReply()
	reply.setReqId(reqId)
	return c.Send(reply)
}

func (c *tcpConn) SetUserdata(userdata any) {
	c.userdata = userdata
}

func (c *tcpConn) GetUserdata() any {
	return c.userdata
}
