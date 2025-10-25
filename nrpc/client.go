package nrpc

import (
	"context"
	"strconv"
	"time"

	"github.com/iakud/knoll/nrpc/codes"
	"github.com/iakud/knoll/nrpc/nrpcutil"
	"github.com/iakud/knoll/nrpc/status"
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
		return status.Errorf(codes.Internal, "nrpc: error marshaling: %v", err.Error())
	}
	msg := nats.NewMsg(c.subj)

	if dl, ok := ctx.Deadline(); ok {
		timeout := time.Until(dl)
		if timeout <= 0 {
			return status.New(codes.DeadlineExceeded, context.DeadlineExceeded.Error()).Err()
		}
		msg.Header.Set(timeoutHdr, nrpcutil.EncodeDuration(timeout))
	}

	msg.Header.Set(methodHdr, method)
	msg.Data = data
	m, err := c.nc.RequestMsgWithContext(ctx, msg)
	if err != nil {
		return err
	}
	// status
	if value := m.Header.Get(statusHdr); len(value) > 0 {
		code, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return status.Errorf(codes.Internal, "malformed nrpc-status: %v", err.Error())
		}
		message := m.Header.Get(messageHdr)
		return status.New(codes.Code(int32(code)), message).Err()
	}

	if err := proto.Unmarshal(m.Data, reply); err != nil {
		return status.Errorf(codes.Internal, "nrpc: failed to unmarshal the received message: %v", err.Error())
	}
	return nil
}
