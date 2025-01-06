package actor

import (
	"context"
	"log/slog"
	"testing"
	"time"
)

type MessageTest struct {
	CmdId   int32
	Message string
}

type ActorTest struct {
}

func (a *ActorTest) Receive(ctx *Context) {
	switch msg := ctx.Message().(type) {
	case Started:
		slog.Info("actor start")
	case Stopped:
		slog.Info("actor close")
	case MessageTest:
		slog.Info("receive message:", "CmdId", msg.CmdId, "Message", msg.Message)
	default:
		slog.Info("receive message unknow type")
	}
}

func TestActor(t *testing.T) {
	actorSystem := NewSystem()
	pid, err := actorSystem.Spawn("hello", &ActorTest{})
	if err != nil {
		t.Fatal(err)
	}
	actorSystem.Send(pid, MessageTest{2, "hello"})
	actorSystem.Shutdown(context.Background(), pid)
}

/*
func TestChan(t *testing.T) {
	c1 := make(chan struct{}, 1)
	close(c1)

	select {
	case c1 <- struct{}{}:
		t.Log("ok")
	default:
		t.Log("default")
	}
}
*/

type ActorStart struct {
	pid *PID
}

type Actor2 struct {
}

func (a *Actor2) OnStart() {
}

func (a *Actor2) OnClose() {
}

func (a *Actor2) Receive(ctx *Context) {
	switch msg := ctx.Message().(type) {
	case Started:
		slog.Info("actor2 started")
	case Stopped:
		slog.Info("actor2 stopped")
	case string:
		slog.Info("actor2 receive,", "message", msg)
		ctx.Respond("456")
	default:
		slog.Info("receive message unknow type")
	}
}

type Actor1 struct {
}

func (a *Actor1) OnStart() {
}

func (a *Actor1) OnClose() {
}

func (a *Actor1) Receive(ctx *Context) {
	switch msg := ctx.Message().(type) {
	case Started:
		slog.Info("actor1 started")
	case Stopped:
		slog.Info("actor1 stopped")
	case *ActorStart:
		ctx1, cancel := context.WithTimeout(context.Background(), time.Second)
		_ = cancel
		if resp, err := ctx.Request(ctx1, msg.pid, "123"); err == nil {
			slog.Info("actor1 get response,", "message", resp.(string))
		}
	case string:
		slog.Info("actor1 receive,", "message", msg)
	default:
		slog.Info("receive message unknow type")
	}
}

func TestRequest(t *testing.T) {
	actorSystem := NewSystem()
	pid1, err := actorSystem.Spawn("actor1", &Actor1{})
	if err != nil {
		t.Fatal(err)
	}
	pid2, err := actorSystem.Spawn("actor2", &Actor2{})
	if err != nil {
		t.Fatal(err)
	}
	actorSystem.Send(pid1, &ActorStart{pid2})
	actorSystem.Shutdown(context.Background(), pid1)
	actorSystem.Shutdown(context.Background(), pid2)
}
