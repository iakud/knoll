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

func (a *ActorTest) OnStart() {

}

func (a *ActorTest) OnClose() {

}

func (a *ActorTest) Receive(ctx context.Context, message MessageTest) {
	slog.Info("receive message:", "CmdId", message.CmdId, "Message", message.Message)
}

func TestActor(t *testing.T) {
	actorSystem := NewActorSystem[MessageTest]()
	actorRef, err := actorSystem.Spawn("hello", &ActorTest{})
	if err != nil {
		t.Fatal(err)
	}
	actorSystem.Send(context.Background(), actorRef, MessageTest{2, "hello"})
	time.Sleep(time.Second)
}
