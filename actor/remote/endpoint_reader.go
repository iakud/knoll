package remote

import (
	"context"
	"errors"
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
	defer slog.Debug("streamreader terminated")

	for {
		envelope, err := stream.Recv()
		if err != nil {
			if errors.Is(err, context.Canceled) {
				break
			}
			slog.Error("EndpointReader receive", "err", err)
			return err
		}

		payload, err := r.serializer.Deserialize(envelope.TypeName, envelope.Message)
		if err != nil {
			slog.Error("EndpointReader deserialize", "err", err)
			return err
		}

		r.remote.system.SendLocal(envelope.Target, payload, envelope.Sender)
	}
	return nil
}

func (r *endpointReader) mustEmbedUnimplementedRemoteServer() {

}
