package actor

import (
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

func (a *ActorTest) OnStart() {
	slog.Info("actor start")
}

func (a *ActorTest) OnClose() {
	slog.Info("actor close")
}

func (a *ActorTest) Receive(ctx Context, message any) {
	switch msg := message.(type) {
	case *MessageTest:
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
	context := actorSystem.Context()
	context.Send(pid, &MessageTest{2, "hello"})
	time.Sleep(time.Second)
}
