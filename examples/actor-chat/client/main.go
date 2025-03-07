package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log/slog"
	"math/rand"
	"os"

	"github.com/iakud/knoll/actor"
	"github.com/iakud/knoll/actor/remote"
	"github.com/iakud/knoll/examples/actor-chat/messages"
)

type client struct {
	username  string
	serverPID *actor.PID
}

func newClient(username string, serverPID *actor.PID) *client {
	return &client{
		username:  username,
		serverPID: serverPID,
	}
}

func (c *client) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case *messages.Message:
		fmt.Printf("%s: %s\n", msg.Username, msg.Msg)
	case actor.Started:
		ctx.Send(c.serverPID, &messages.Connect{
			Username: c.username,
		})
	case actor.Stopped:
		slog.Info("client stopped")
	}
}

func main() {
	var (
		listenAt  = flag.String("listen", "", "specify address to listen to, will pick a random port if not specified")
		connectTo = flag.String("connect", "127.0.0.1:4000", "the address of the server to connect to")
		username  = flag.String("username", os.Getenv("USER"), "")
	)
	flag.Parse()
	if *listenAt == "" {
		*listenAt = fmt.Sprintf("127.0.0.1:%d", rand.Int31n(50000)+10000)
	}

	r := remote.New(*listenAt)
	s := actor.NewSystemWithConfig(actor.WithRemote(r))
	r.Start(s)
	var (
		serverPID = actor.NewPID(*connectTo, "server/primary")
		clientPID = s.Spawn("client"+"/"+*username, newClient(*username, serverPID))
		scanner   = bufio.NewScanner(os.Stdin)
	)
	fmt.Println("Type 'quit' and press return to exit.")
	for scanner.Scan() {
		msg := &messages.Message{
			Msg:      scanner.Text(),
			Username: *username,
		}
		if msg.Msg == "quit" {
			break
		}
		s.SendWithSender(serverPID, msg, clientPID)
	}
	if err := scanner.Err(); err != nil {
		slog.Error("failed to read message from stdin", "err", err)
	}

	s.SendWithSender(serverPID, &messages.Disconnect{}, clientPID)
	s.Shutdown(context.Background(), clientPID)
	r.Stop()
	slog.Info("client disconnected")
}
