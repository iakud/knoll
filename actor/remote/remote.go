package remote

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/iakud/knoll/actor"
	"google.golang.org/grpc"
)

type Remote struct {
	address   string
	system    *actor.System
	routerPID *actor.PID
	server    *grpc.Server
}

func New(address string) *Remote {
	r := &Remote{
		address: address,
	}
	return r
}

func (r *Remote) Start(system *actor.System) {
	r.system = system
	ln, err := net.Listen("tcp", r.address)
	if err != nil {
		panic(fmt.Errorf("remote: Failed to listen: %w", err))
	}
	r.routerPID = system.Spawn("router", newEndpointRouter(system))
	r.server = grpc.NewServer()
	RegisterRemoteServer(r.server, newEndpointReader(r))
	slog.Info("remote: Starting", slog.String("address", r.address))
	go r.server.Serve(ln)
}

func (r *Remote) Shutdown() {
	r.server.GracefulStop()
	r.system.Shutdown(context.Background(), r.routerPID)
}

func (r *Remote) Stop() {
	r.server.Stop()
	r.system.Stop(r.routerPID)
}

func (r *Remote) Send(pid *actor.PID, msg any, sender *actor.PID) {
	rd := &remoteDeliver{
		target:  pid,
		sender:  sender,
		message: msg,
	}
	r.system.Send(r.routerPID, rd)
}

func (r *Remote) Address() string {
	return r.address
}
