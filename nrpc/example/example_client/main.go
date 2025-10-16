package main

import (
	"context"
	"log/slog"

	"github.com/iakud/knoll/nrpc"
	"github.com/iakud/knoll/nrpc/example"
	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	c := nrpc.NewClient(nc, "example1")
	client := example.NewExampleClient(c)
	var req example.TestRequest
	req.SetText("hello")
	for i := 0; i < 10; i++ {
		reply, err := client.Test(context.TODO(), &req)
		if err != nil {
			slog.Info("client test", "error", err)
		} else {
			slog.Info("client test", "reply", reply)
		}
	}
}
