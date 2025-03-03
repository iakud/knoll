package remote

import (
	"context"
	"errors"
	"log/slog"
)

type streamReader struct {
	remote     *Remote
	serializer Serializer
}

func newStreamReader(r *Remote) *streamReader {
	return &streamReader{
		remote:     r,
		serializer: ProtoSerializer{},
	}
}

func (r *streamReader) Receive(stream Remote_ReceiveServer) error {
	defer slog.Debug("streamreader terminated")

	for {
		envelope, err := stream.Recv()
		if err != nil {
			if errors.Is(err, context.Canceled) {
				break
			}
			slog.Error("streamReader receive", "err", err)
			return err
		}

		payload, err := r.serializer.Deserialize(envelope.TypeName, envelope.Message)
		if err != nil {
			slog.Error("streamReader deserialize", "err", err)
			return err
		}

		r.remote.system.SendLocal(envelope.Target, payload, envelope.Sender)
	}
	return nil
}
