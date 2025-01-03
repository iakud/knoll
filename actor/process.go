package actor

import (
	"context"
)

type Processer interface {
	Send(message any, sender *PID)
	Start()
	Stop()
}

type process struct {
	context *Context
	mailbox *mailbox
	actor   Actor
}

func newProcess(pid *PID, system *System, actor Actor) Processer {
	p := &process{
		context: newContext(pid, system),
		actor:   actor,
	}
	p.mailbox = newMailbox(p)
	return p
}

func (p *process) Send(message any, sender *PID) {
	p.mailbox.Send(context.Background(), Envelope{message, sender})
}

func (p *process) Start() {
	/*
		p.context.envelope.Message = started
		p.InvokeMessage()
		p.context.envelope.Message = nil
	*/
	p.mailbox.Start()
}

func (p *process) Stop() {
	p.mailbox.Send(context.Background(), Envelope{stopped, nil})
}

func (p *process) Invoke(envelope Envelope) {
	message := envelope.Message
	switch message.(type) {
	case Started:

	case Stopped:
		p.handleStop()
	case PoisonPill:
		p.Stop()
	default:
		p.context.envelope = envelope
		p.InvokeMessage()
		p.context.envelope = Envelope{}
	}
}

func (p *process) InvokeMessage() {
	p.actor.Receive(p.context)
}

func (p *process) handleStop() {
	p.context.system.registry.Remove(p.context.pid)
	p.mailbox.Stop()
}
