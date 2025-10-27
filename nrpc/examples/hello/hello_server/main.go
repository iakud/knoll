package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/iakud/knoll/nrpc"
	"github.com/iakud/knoll/nrpc/examples/hello/hello"
	"github.com/nats-io/nats.go"
)

type helloServer struct {
}

func (s *helloServer) SayHello(ctx context.Context, req *hello.SayHelloRequest) (*hello.SayHelloReply, error) {
	slog.Info("sayhello", "req", req)
	reply := &hello.SayHelloReply{}
	reply.SetMessage(fmt.Sprintf("hello %s", req.GetName()))
	return reply, nil
}

func main() {
	nc, err := nats.Connect("nats://10.10.45.25:4222")
	if err != nil {
		panic(err)
	}
	s := nrpc.NewServer(nc, "hello", "hello")
	hello.RegisterHelloServer(s, &helloServer{})
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
