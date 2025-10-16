package nrpc

import (
	"context"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	nc *nats.Conn
}

func NewClient(nc *nats.Conn) *Client {
	return &Client{nc}
}

func (c *Client) Call(ctx context.Context, subj string, method string, req proto.Message, reply proto.Message) error {
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}
	var msg nats.Msg
	msg.Subject = subj
	msg.Header.Add("method", method)
	msg.Data = data
	m, err := c.nc.RequestMsgWithContext(ctx, &msg)
	if err != nil {
		return err
	}

	// c.nc.Subscribe("1", func(msg *nats.Msg) {
	// msg.Respond()
	// })

	return proto.Unmarshal(m.Data, reply)
}
