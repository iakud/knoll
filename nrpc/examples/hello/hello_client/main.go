package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/iakud/knoll/nrpc"
	"github.com/iakud/knoll/nrpc/examples/hello/hello"
	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("nats://10.10.45.25:4222")
	if err != nil {
		panic(err)
	}
	c := nrpc.NewClient(nc, "hello")
	client := hello.NewHelloClient(c)

	for i := 0; i < 10; i++ {
		var req hello.SayHelloRequest
		req.SetName(fmt.Sprintf("world %d", i))
		reply, err := client.SayHello(context.Background(), &req)
		if err != nil {
			slog.Info("client test", "error", err)
		} else {
			slog.Info("client test", "reply", reply)
		}
	}
}
