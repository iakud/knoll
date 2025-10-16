package nrpc

import (
	"context"

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
	var msg nats.Msg
	msg.Subject = c.subj
	msg.Header.Set("method", method)
	msg.Data = data
	m, err := c.nc.RequestMsgWithContext(ctx, &msg)
	if err != nil {
		return err
	}

	return proto.Unmarshal(m.Data, reply)
}
