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
		w.sendMessage(msg)
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
			slog.Error("remote: EndpointWriter failed to connect", slog.String("address", w.address), slog.Any("error", err), slog.Int("retry", i))
			time.Sleep(tempDelay)
			continue
		}
		slog.Info("remote: EndpointWriter connected", slog.String("address", w.address), slog.Duration("cost", time.Since(now)))
		return
	}
	w.system.Send(w.router, &RemoteUnreachableEvent{Address: w.address})
}

func (w *endpointWriter) connect() error {
	conn, err := grpc.Dial(w.address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	w.conn = conn
	client := NewRemoteClient(conn)
	stream, err := client.Receive(context.Background())
	if err != nil {
		slog.Error("remote: EndpointWriter failed to create receive stream", slog.String("address", w.address), slog.Any("error", err))
		return err
	}
	w.stream = stream

	go func() {
		_, err := w.stream.Recv()
		if errors.Is(err, io.EOF) {
			slog.Debug("remote: EndpointWriter stream completed", slog.String("address", w.address))
		} else if err != nil {
			slog.Error("remote: EndpointWriter lost connection", slog.String("address", w.address), slog.Any("error", err))
		} else {
			slog.Info("remote: EndpointWriter disconnect from remote", slog.String("address", w.address))
		}
		w.system.Send(w.router, &RemoteUnreachableEvent{Address: w.address})
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

func (w *endpointWriter) sendMessage(msg *remoteDeliver) {
	typeName, data, err := w.serializer.Serialize(msg.message)
	if err != nil {
		slog.Error("remote: EndpointWriter failed to serialize", slog.String("address", w.address), slog.Any("error", err), slog.Any("message", msg.message))
		return
	}
	envelope := &Envelope{
		Target:   msg.target,
		Sender:   msg.sender,
		TypeName: typeName,
		Message:  data,
	}
	if err := w.stream.Send(envelope); err != nil {
		w.system.Send(w.router, &RemoteUnreachableEvent{Address: w.address})
		slog.Debug("remote: EndpointWriter failed to send", slog.String("address", w.address), slog.Any("error", err))
	}
}
