package actor

import "context"

type Context struct {
	pid      *PID
	system   *System
	actor    Actor
	envelope Envelope
	response chan response
}

func newContext(pid *PID, system *System, actor Actor) *Context {
	return &Context{
		pid:    pid,
		system: system,
		actor:  actor,
	}
}

func (c *Context) Sender() *PID {
	return c.envelope.Sender
}

func (c *Context) Message() any {
	if req, ok := c.envelope.Message.(*request); ok {
		return req.message
	}
	return c.envelope.Message
}

func (c *Context) Send(pid *PID, message any) {
	c.system.SendWithSender(pid, message, c.pid)
}

type request struct {
	message  any
	response chan response
}

type response struct {
	resp any
	err  error
}

func (c *Context) Request(ctx context.Context, pid *PID, message any) (any, error) {
	respc := make(chan response, 1)
	c.system.SendWithSender(pid, &request{message, respc}, c.pid)
	select {
	case resp := <-respc:
		return resp.resp, resp.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (c *Context) Respond(message any) {
	if c.envelope.Sender == nil {
		return
	}
	if req, ok := c.envelope.Message.(*request); ok {
		select {
		case req.response <- response{message, nil}:
		default:
		}
		return
	}
	c.system.SendWithSender(c.envelope.Sender, message, c.pid)
}

func (c *Context) System() *System {
	return c.system
}
