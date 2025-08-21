package krpc

import (
	"github.com/iakud/knoll/knet"
)

type wsConn struct {
	id       uint64
	wsconn   *knet.WSConn
	hash     uint64
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
