package actor

import (
	"context"
)

type Context struct {
	pid      *PID
	system   *System
	actor    Actor
	envelope Envelope
}

func newContext(pid *PID, system *System, actor Actor) *Context {
	return &Context{
		pid:    pid,
		system: system,
		actor:  actor,
	}
}

func (c *Context) PID() *PID {
	return c.pid
}

func (c *Context) Sender() *PID {
	return c.envelope.Sender
}

func (c *Context) Message() any {
	return c.envelope.Message
}

func (c *Context) Send(pid *PID, message any) {
	c.system.SendWithSender(pid, message, c.pid)
}

func (c *Context) Forward(pid *PID) {
	c.system.SendWithSender(pid, c.envelope.Message, c.pid)
}

func (c *Context) Request(ctx context.Context, pid *PID, message any) (any, error) {
	return c.system.Request(ctx, pid, message)
}

func (c *Context) Respond(message any) {
	if c.envelope.Sender == nil {
		return
	}
	c.system.SendWithSender(c.envelope.Sender, message, c.pid)
}

func (c *Context) System() *System {
	return c.system
}
