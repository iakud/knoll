package actor

import (
	"context"
	"log/slog"
)

type Processer interface {
	Start()
	Stop()
	Close()
	Send(message any, sender *PID)
}

type process struct {
	pid     *PID
	system  *System
	mailbox *mailbox

	actor Actor
}

func newProcess(pid *PID, system *System, actor Actor) Processer {
	return &process{
		pid:     pid,
		system:  system,
		mailbox: newMailbox(),

		actor: actor,
	}
}

func (proc *process) Send(message any, sender *PID) {
	proc.mailbox.Send(context.Background(), Envelope{message, sender})
}

func (proc *process) Start() {
	go proc.process()
}

func (proc *process) Stop() {

}

func (proc *process) process() {
	defer func() {
		if r := recover(); r != nil {
			slog.Info("Receive recovered in %v", r)
		}
	}()
	proc.actor.OnStart()
	defer proc.actor.OnClose()

	for {
		select {
		case envelope := <-proc.mailbox.MessageC():
			proc.Invoke(envelope)
		case <-proc.mailbox.Done():
			return
		}
	}
}

func (proc *process) Close() {

}

func (proc *process) Invoke(envelope Envelope) {
	message := envelope.Message
	switch message.(type) {
	case PoisonPill:
	}
	ctx := &actorContext{
		envelope: envelope,
		system:   proc.system,
		pid:      proc.pid,
	}
	proc.actor.Receive(ctx, envelope.Message)
}
