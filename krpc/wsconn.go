package krpc

import (
	"time"

	"github.com/iakud/knoll/knet"
)

type wsConn struct {
	id       uint64
	wsconn   *knet.WSConn
	hash     uint64
	timer    *time.Timer
	userdata any
}

func newWSConn(id uint64, wsconn *knet.WSConn, hash uint64) *wsConn {
	return &wsConn{id: id, wsconn: wsconn, hash: hash}
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

func (c *wsConn) Send(msg Msg) error {
	data, err := Marshal(msg)
	if err != nil {
		return err
	}
	return c.wsconn.Send(data)
}

func (c *wsConn) SendOK() error {
	return nil
}

func (c *wsConn) SendError(err error) error {
	return nil
}

func (c *wsConn) Reply(reqId uint32, msg Msg) error {
	msg.setFlagReply()
	msg.setReqId(reqId)
	return c.Send(msg)
}

func (c *wsConn) ReplyOK(reqId uint32) error {
	return nil
}

func (c *wsConn) ReplyError(reqId uint32, err error) error {
	return nil
}

func (c *wsConn) SetUserdata(userdata any) {
	c.userdata = userdata
}

func (c *wsConn) GetUserdata() any {
	return c.userdata
}
