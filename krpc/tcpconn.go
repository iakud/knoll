package krpc

import (
	"time"

	"github.com/iakud/knoll/knet"
)

type tcpConn struct {
	id       uint64
	tcpconn  *knet.TCPConn
	hash     uint64
	timer    *time.Timer
	Userdata any
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

func (c *tcpConn) Send(msg Msg) error {
	data, err := Marshal(msg)
	if err != nil {
		return err
	}
	return c.tcpconn.Send(data)
}

func (c *tcpConn) SendOK() error {
	return nil
}

func (c *tcpConn) SendError(err error) error {
	return nil
}

func (c *tcpConn) Reply(reqId uint32, msg Msg) error {
	msg.setFlagReply()
	msg.setReqId(reqId)
	return c.Send(msg)
}

func (c *tcpConn) ReplyOK(reqId uint32) error {
	return nil
}

func (c *tcpConn) ReplyError(reqId uint32, err error) error {
	return nil
}
