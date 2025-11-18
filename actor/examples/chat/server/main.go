package main

import (
	"flag"
	"log/slog"

	"github.com/iakud/knoll/actor"
	"github.com/iakud/knoll/actor/examples/chat/message"
	"github.com/iakud/knoll/actor/remote"
)

type clientMap map[string]*actor.PID
type userMap map[string]string

type server struct {
	clients clientMap // key: address value: *pid
	users   userMap   // key: address value: username
}

func newServer() *server {
	return &server{
		clients: make(clientMap),
		users:   make(userMap),
	}
}

func (s *server) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case *message.Message:
		slog.Info("message received", "msg", msg.Msg, "from", ctx.Sender())
		s.handleMessage(ctx)
	case *message.Disconnect:
		cAddr := ctx.Sender().GetAddress()
		pid, ok := s.clients[cAddr]
		if !ok {
			slog.Warn("unknown client disconnected", "client", pid.Address)
			return
		}
		username, ok := s.users[cAddr]
		if !ok {
			slog.Warn("unknown user disconnected", "client", pid.Address)
			return
		}
		slog.Info("client disconnected", "username", username)
		delete(s.clients, cAddr)
		delete(s.users, username)
	case *message.Connect:
		cAddr := ctx.Sender().GetAddress()
		if _, ok := s.clients[cAddr]; ok {
			slog.Warn("client already connected", "client", ctx.Sender().GetID())
			return
		}
		if _, ok := s.users[cAddr]; ok {
			slog.Warn("user already connected", "client", ctx.Sender().GetID())
			return
		}
		s.clients[cAddr] = ctx.Sender()
		s.users[cAddr] = msg.Username
		slog.Info("new client connected",
			"id", ctx.Sender().GetID(), "addr", ctx.Sender().GetAddress(), "sender", ctx.Sender(),
			"username", msg.Username,
		)
	}
}

func (s *server) handleMessage(ctx *actor.Context) {
	for _, pid := range s.clients {
		if pid.Equals(ctx.Sender()) {
			continue
		}
		slog.Info("forwarding message", "pid", pid.ID, "addr", pid.Address, "msg", ctx.Message())
		ctx.Forward(pid)
	}
}

func main() {
	var (
		listenAt = flag.String("listen", "127.0.0.1:4000", "")
	)
	flag.Parse()
	r := remote.New(*listenAt)
	s := actor.NewSystemWithConfig(actor.WithRemote(r))
	r.Start(s)

	s.Spawn("server"+"/"+"primary", newServer())

	select {}
}
