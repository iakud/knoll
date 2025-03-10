package remote

import (
	"errors"
	"io"
	"log/slog"

	"github.com/iakud/knoll/actor"
)

type endpointReader struct {
	system     *actor.System
	serializer Serializer
}

func newEndpointReader(system *actor.System) *endpointReader {
	return &endpointReader{
		system:     system,
		serializer: ProtoSerializer{},
	}
}

func (r *endpointReader) Receive(stream Remote_ReceiveServer) error {
	for {
		envelope, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				r.system.Logger().Info("remote: Reader stream closed")
				return nil
			}
			r.system.Logger().Error("remote: Reader failed to read", slog.Any("error", err))
			return err
		}

		message, err := r.serializer.Deserialize(envelope.TypeName, envelope.Message)
		if err != nil {
			r.system.Logger().Error("remote: Reader failed to deserialize", slog.Any("error", err))
			return err
		}

		r.system.SendLocal(envelope.Target, message, envelope.Sender)
	}
}

func (r *endpointReader) mustEmbedUnimplementedRemoteServer() {}
