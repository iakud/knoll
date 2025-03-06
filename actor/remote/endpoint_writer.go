package remote

import (
	"context"
	"log/slog"
	"time"

	"github.com/iakud/knoll/actor"
	grpc "google.golang.org/grpc"
)

const kMaxRetryCount int = 3

type endpointWriter struct {
	address    string
	system     *actor.System
	router     *actor.PID
	serializer Serializer

	conn   *grpc.ClientConn
	stream Remote_ReceiveClient
}

func newEndpointWriter(system *actor.System, router *actor.PID, address string) *endpointWriter {
	return &endpointWriter{
		address:    address,
		system:     system,
		router:     router,
		serializer: ProtoSerializer{},
	}
}

func (w *endpointWriter) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case actor.Started:
		w.start()
	case actor.Stopped:
		w.stop()
	case *remoteDeliver:
		w.send(msg)
	}
}

func (w *endpointWriter) start() {
	now := time.Now()
	var tempDelay time.Duration
	for i := 0; i < kMaxRetryCount; i++ {
		if err := w.connect(); err != nil {
			if tempDelay == 0 {
				tempDelay = 5 * time.Millisecond
			} else {
				tempDelay *= 2
			}
			slog.Error("EndpointWriter failed to connect", slog.String("address", w.address), slog.Any("error", err), slog.Int("retry", i))
			time.Sleep(tempDelay)
			continue
		}
		slog.Info("EndpointWriter connected", slog.String("address", w.address), slog.Duration("cost", time.Since(now)))
		return
	}
	terminated := &RemoteUnreachableEvent{Address: w.address}
	w.system.Send(w.router, terminated)
}

func (w *endpointWriter) connect() error {
	conn, err := grpc.Dial(w.address)
	if err != nil {
		return err
	}
	w.conn = conn
	client := NewRemoteClient(conn)
	stream, err := client.Receive(context.Background())
	if err != nil {
		slog.Error("EndpointWriter failed to create receive stream", slog.String("address", w.address), slog.Any("error", err))
		return err
	}
	w.stream = stream

	go func() {
		if _, err := w.stream.Recv(); err != nil {
			slog.Error("EndpointWriter lost connection", slog.String("address", w.address), slog.Any("error", err))
		} else {
			slog.Info("EndpointWriter disconnected from remote", slog.String("address", w.address))
		}
		terminated := &RemoteUnreachableEvent{Address: w.address}
		w.system.Send(w.router, terminated)
	}()
	return nil
}

func (w *endpointWriter) stop() {
	if w.stream != nil {
		w.stream.CloseSend()
		w.stream = nil
	}
	if w.conn != nil {
		w.conn.Close()
		w.conn = nil
	}
}

func (w *endpointWriter) send(msg *remoteDeliver) {
	data, err := w.serializer.Serialize(msg.message)
	if err != nil {
		slog.Error("EndpointWriter failed to serialize message", slog.String("address", w.address), slog.Any("error", err), slog.Any("message", msg.message))
		return
	}
	tname := w.serializer.TypeName(msg)

	envelope := &Envelope{}
	envelope.Target = msg.target
	envelope.Sender = msg.target
	envelope.TypeName = tname
	envelope.Message = data
	if err := w.stream.Send(envelope); err != nil {
		terminated := &RemoteUnreachableEvent{Address: w.address}
		w.system.Send(w.router, terminated)
		slog.Debug("gRPC Failed to send", slog.String("address", w.address), slog.Any("error", err))
	}
}
