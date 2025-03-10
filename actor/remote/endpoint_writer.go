package remote

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"time"

	"github.com/iakud/knoll/actor"
	grpc "google.golang.org/grpc"
)

type endpointWriter struct {
	address    string
	system     *actor.System
	routerPID  *actor.PID
	serializer Serializer

	conn   *grpc.ClientConn
	stream Remote_ReceiveClient
}

func newEndpointWriter(system *actor.System, routerPID *actor.PID, address string) *endpointWriter {
	return &endpointWriter{
		address:    address,
		system:     system,
		routerPID:  routerPID,
		serializer: ProtoSerializer{},
	}
}

func (w *endpointWriter) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case actor.Started:
		w.init()
	case actor.Stopped:
		w.closeConn()
	case *remoteDeliver:
		w.sendMessage(msg)
	}
}

func (w *endpointWriter) init() {
	now := time.Now()

	const kMaxRetryCount int = 3
	var tempDelay time.Duration
	for i := 0; i < kMaxRetryCount; i++ {
		if err := w.initConnect(); err != nil {
			if tempDelay == 0 {
				tempDelay = 5 * time.Millisecond
			} else {
				tempDelay *= 2
			}
			w.system.Logger().Error("remote: Writer failed to connect", slog.String("address", w.address), slog.Any("error", err), slog.Int("retry", i))
			time.Sleep(tempDelay)
			continue
		}
		w.system.Logger().Info("remote: Writer connected", slog.String("address", w.address), slog.Duration("cost", time.Since(now)))
		return
	}
	w.system.Send(w.routerPID, &RemoteUnreachableEvent{Address: w.address})
}

func (w *endpointWriter) initConnect() error {
	conn, err := grpc.Dial(w.address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	w.conn = conn
	client := NewRemoteClient(conn)
	stream, err := client.Receive(context.Background())
	if err != nil {
		w.system.Logger().Error("remote: Writer failed to create receive stream", slog.String("address", w.address), slog.Any("error", err))
		return err
	}
	w.stream = stream

	go func() {
		_, err := w.stream.Recv()
		if errors.Is(err, io.EOF) {
			w.system.Logger().Debug("remote: Writer stream completed", slog.String("address", w.address))
		} else if err != nil {
			w.system.Logger().Error("remote: Writer lost connection", slog.String("address", w.address), slog.Any("error", err))
		} else {
			w.system.Logger().Info("remote: Writer disconnect from remote", slog.String("address", w.address))
		}
		w.system.Send(w.routerPID, &RemoteUnreachableEvent{Address: w.address})
	}()
	return nil
}

func (w *endpointWriter) closeConn() {
	w.system.Logger().Debug("remote: Writer closing connection", slog.String("address", w.address))
	if w.stream != nil {
		w.stream.CloseSend()
		w.stream = nil
	}
	if w.conn != nil {
		w.conn.Close()
		w.conn = nil
	}
}

func (w *endpointWriter) sendMessage(msg *remoteDeliver) {
	typeName, data, err := w.serializer.Serialize(msg.message)
	if err != nil {
		w.system.Logger().Error("remote: Writer failed to serialize", slog.String("address", w.address), slog.Any("error", err), slog.Any("message", msg.message))
		return
	}
	envelope := &Envelope{
		Target:   msg.target,
		Sender:   msg.sender,
		TypeName: typeName,
		Message:  data,
	}
	if err := w.stream.Send(envelope); err != nil {
		w.system.Send(w.routerPID, &RemoteUnreachableEvent{Address: w.address})
		w.system.Logger().Debug("remote: Writer failed to send", slog.String("address", w.address), slog.Any("error", err))
	}
}
