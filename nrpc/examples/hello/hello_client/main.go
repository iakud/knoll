package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/iakud/knoll/nrpc"
	"github.com/iakud/knoll/nrpc/examples/hello/hello"
	"github.com/nats-io/nats.go"
)

func main() {
	var natsURL = nats.DefaultURL
	if len(os.Args) == 2 {
		natsURL = os.Args[1]
	}
	nc, err := nats.Connect(natsURL)
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
