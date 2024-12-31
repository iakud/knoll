package actor

type Context interface {
	Sender() *PID
	Message() interface{}
	Send(pid *PID, message any)
	Close(pid *PID)
}

type actorContext struct {
	envelope Envelope
	system   *System
	pid      *PID
}

func (ctx *actorContext) Sender() *PID {
	return ctx.envelope.Sender
}

func (ctx *actorContext) Message() interface{} {
	return ctx.envelope.Message
}

func (c *actorContext) Send(pid *PID, message any) {
	proc := c.system.registry.Get(pid)
	if proc == nil {
		return
	}
	proc.Send(message, c.pid)
}

func (c *actorContext) Close(pid *PID) {
	proc := c.system.registry.Get(pid)
	if proc == nil {
		return
	}
	proc.Close()
}
