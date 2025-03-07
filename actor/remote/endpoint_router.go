package remote

import (
	"context"
	"log/slog"

	"github.com/iakud/knoll/actor"
)

type endpointRouter struct {
	system      *actor.System
	connections map[string]*actor.PID
	pid         *actor.PID
}

func newEndpointRouter(system *actor.System) *endpointRouter {
	return &endpointRouter{
		connections: make(map[string]*actor.PID),
		system:      system,
	}
}

func (r *endpointRouter) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case actor.Started:
		r.pid = ctx.PID()
	case actor.Stopped:
		r.stop()
	case *remoteDeliver:
		r.remoteDeliver(msg)
	case *RemoteUnreachableEvent:
		r.removeEndpoint(msg.Address)
	}
}

func (r *endpointRouter) removeEndpoint(address string) {
	edpWriter, ok := r.connections[address]
	if !ok {
		return
	}
	delete(r.connections, address)
	slog.Info("remote: EndpointRouter remove endpoint", "address", address)
	r.system.Stop(edpWriter)
}

func (r *endpointRouter) remoteDeliver(msg *remoteDeliver) {
	address := msg.target.Address
	edpWriter, ok := r.connections[address]
	if !ok {
		edpWriter = r.system.Spawn("endpoint"+"/"+address, newEndpointWriter(r.system, r.pid, address))
		r.connections[address] = edpWriter
	}
	r.system.Send(edpWriter, msg)
}

func (r *endpointRouter) stop() {
	for _, writerPID := range r.connections {
		r.system.Shutdown(context.Background(), writerPID)
	}
	clear(r.connections)
}
