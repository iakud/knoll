package actor

type Processer interface {
	Send(message any, sender *PID)
	Start()
	Stop()
}

type process struct {
	context *Context
	mailbox *mailbox
	stopped bool
}

func newProcess(context *Context) Processer {
	p := &process{
		context: context,
	}
	p.mailbox = newMailbox(p)
	return p
}

func (p *process) Start() {
	p.mailbox.Send(Envelope{started, nil})
}

func (p *process) Stop() {
	p.mailbox.Send(Envelope{stopped, nil})
}

func (p *process) Send(message any, sender *PID) {
	p.mailbox.Send(Envelope{message, sender})
}

func (p *process) Invoke(envelope Envelope) {
	message := envelope.Message
	switch message.(type) {
	case Started:
		p.handleStarted()
	case Stopped:
		p.handleStopped()
	case PoisonPill:
		p.context.system.Stop(p.context.pid)
	default:
		p.invokeMessage(envelope)
	}
}

func (p *process) handleStarted() {
	p.invokeMessage(Envelope{started, nil})
}

func (p *process) handleStopped() {
	if p.stopped {
		return
	}
	p.stopped = true
	p.context.system.registry.Remove(p.context.pid)
	p.invokeMessage(Envelope{stopped, nil})
}

func (p *process) invokeMessage(envelope Envelope) {
	p.context.envelope = envelope
	p.context.actor.Receive(p.context)
	p.context.envelope = Envelope{}
}
