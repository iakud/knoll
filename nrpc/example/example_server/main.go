package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"

	"github.com/iakud/knoll/nrpc"
	"github.com/iakud/knoll/nrpc/example"
	"github.com/nats-io/nats.go"
)

type server struct {
	num atomic.Int32
}

func (s *server) Test(ctx context.Context, req *example.TestRequest) (*example.TestReply, error) {
	slog.Info("server test", "req", req)
	reply := &example.TestReply{}
	reply.SetText(req.GetText())
	reply.SetNum(s.num.Add(1))
	return reply, nil
}

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	s := nrpc.NewServer(nc, "example1")
	example.RegisterExampleServer(s, &server{})
	if err := s.Start(); err != nil {
		panic(err)
	}
	slog.Info("server started")
	defer slog.Info("server stopped")
	// signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	slog.Info("app", "signal", <-ch)
}
