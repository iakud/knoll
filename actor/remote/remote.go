package remote

import (
	"fmt"
	"log/slog"
	"net"
	sync "sync"
	"sync/atomic"

	"github.com/iakud/knoll/actor"
	"google.golang.org/grpc"
)

type Remote struct {
	addr   string
	system *actor.System
	router *actor.PID
	stopCh chan struct{} // Stop closes this channel to signal the remote to stop listening.
	stopWg *sync.WaitGroup
	state  atomic.Uint32
}

const (
	stateInvalid uint32 = iota
	stateInitialized
	stateRunning
	stateStopped
)

func New(addr string) *Remote {
	r := &Remote{
		addr: addr,
	}
	r.state.Store(stateInitialized)
	return r
}

func (r *Remote) Start(system *actor.System) error {
	if r.state.Load() != stateInitialized {
		return fmt.Errorf("remote already started")
	}
	r.state.Store(stateRunning)
	r.system = system
	ln, err := net.Listen("tcp", r.addr)
	if err != nil {
		return fmt.Errorf("remote failed to listen: %w", err)
	}
	slog.Debug("listening", "addr", r.addr)

	r.router = system.Spawn("router", newEndpointRouter(system))
	slog.Debug("server started", "listenAddr", r.addr)

	s := grpc.NewServer()
	RegisterRemoteServer(s, newEndpointReader(r))
	slog.Info("Starting Proto.Actor server", slog.String("address", r.addr))
	go s.Serve(ln)

	return nil
}

func (r *Remote) Shutdown() {

}

func (r *Remote) Stop() {

}

func (r *Remote) Send(pid *actor.PID, msg any, sender *actor.PID) {
	rd := &remoteDeliver{
		target:  pid,
		sender:  sender,
		message: msg,
	}
	r.system.Send(r.router, rd)
}

func (r *Remote) Address() string {
	return r.addr
}
