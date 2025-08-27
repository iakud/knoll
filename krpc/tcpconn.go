package krpc

import (
	"github.com/iakud/knoll/knet"
)

type tcpConn struct {
	id         uint64
	tcpconn    *knet.TCPConn
	hash       uint64
	userdata   any
	newMessage func() Message
}

func newTCPConn(id uint64, tcpconn *knet.TCPConn, newMessage func() Message) *tcpConn {
	return &tcpConn{id: id, tcpconn: tcpconn, newMessage: newMessage}
}

func (s *tcpConn) Id() uint64 {
	return s.id
}

func (c *tcpConn) setHash(hash uint64) {
	c.hash = hash
}

func (c *tcpConn) Hash() uint64 {
	return c.hash
}

func (c *tcpConn) NewMessage() Message {
	return c.newMessage()
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
	reply.Header().setFlagReply()
	reply.Header().setReqId(reqId)
	return c.Send(reply)
}

func (c *tcpConn) SetUserdata(userdata any) {
	c.userdata = userdata
}

func (c *tcpConn) GetUserdata() any {
	return c.userdata
}
