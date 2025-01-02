package actor

import (
	"context"
	"log/slog"
)

type Processer interface {
	Start()
	Stop()
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

func (p *process) Send(message any, sender *PID) {
	p.mailbox.Send(context.Background(), Envelope{message, sender})
}

func (p *process) Start() {
	go p.process()
}

func (p *process) Stop() {
	p.mailbox.Close()
}

func (p *process) process() {
	defer func() {
		if r := recover(); r != nil {
			slog.Info("Receive recovered in %v", r)
		}
	}()
	p.actor.OnStart()
	defer p.actor.OnClose()

	for {
		select {
		case envelope := <-p.mailbox.MessageC():
			p.Invoke(envelope)
		case <-p.mailbox.Done():
			return
		}
	}
}

func (p *process) Invoke(envelope Envelope) {
	message := envelope.Message
	switch message.(type) {
	case PoisonPill:
		p.Stop()
	default:
		p.InvokeMessage(envelope)
	}
}

func (p *process) InvokeMessage(envelope Envelope) {
	ctx := &actorContext{
		envelope: envelope,
		system:   p.system,
		pid:      p.pid,
	}
	p.actor.Receive(ctx)
}
