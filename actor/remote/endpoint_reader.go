package remote

import (
	"errors"
	"io"
	"log/slog"
)

type endpointReader struct {
	remote     *Remote
	serializer Serializer
}

func newEndpointReader(r *Remote) *endpointReader {
	return &endpointReader{
		remote:     r,
		serializer: ProtoSerializer{},
	}
}

func (r *endpointReader) Receive(stream Remote_ReceiveServer) error {
	for {
		envelope, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				slog.Info("remote: EndpointReader stream closed")
				return nil
			}
			slog.Error("remote: EndpointReader failed to read", slog.Any("error", err))
			return err
		}

		message, err := r.serializer.Deserialize(envelope.TypeName, envelope.Message)
		if err != nil {
			slog.Error("remote: EndpointReader failed to deserialize", slog.Any("error", err))
			return err
		}

		r.remote.system.SendLocal(envelope.Target, message, envelope.Sender)
	}
}

func (r *endpointReader) mustEmbedUnimplementedRemoteServer() {}
