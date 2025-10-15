package nrpc

import (
	"context"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

func Call(ctx context.Context, nc *nats.Conn, subj string, msg proto.Message, reply proto.Message) error {
	data, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	m, err := nc.RequestWithContext(ctx, subj, data)
	if err != nil {
		return err
	}

	return proto.Unmarshal(m.Data, reply)
}
