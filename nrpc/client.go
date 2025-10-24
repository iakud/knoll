package nrpc

import (
	"context"
	"strconv"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type ClientConnInterface interface {
	Invoke(ctx context.Context, method string, args proto.Message, reply proto.Message) error
}

// Assert *ClientConn implements ClientConnInterface.
var _ ClientConnInterface = (*ClientConn)(nil)

type ClientConn struct {
	nc   *nats.Conn
	subj string
}

func NewClient(nc *nats.Conn, subj string) *ClientConn {
	return &ClientConn{nc, subj}
}

func (c *ClientConn) Invoke(ctx context.Context, method string, args proto.Message, reply proto.Message) error {
	data, err := proto.Marshal(args)
	if err != nil {
		return err
	}
	msg := nats.NewMsg(c.subj)
	msg.Header.Set(methodHdr, method)
	msg.Data = data
	m, err := c.nc.RequestMsgWithContext(ctx, msg)
	if err != nil {
		return err
	}
	status := m.Header.Get(statusHdr)
	if len(status) > 0 {
		c, err := strconv.Atoi(status)
		if err != nil {
			return err
		}
		message := m.Header.Get(messageHdr)
		return New(Code(c), message).Err()
	}

	return proto.Unmarshal(m.Data, reply)
}
