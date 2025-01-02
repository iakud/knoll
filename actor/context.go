package actor

import "context"

type Context interface {
	Sender() *PID
	Message() interface{}
	Send(pid *PID, message any)
	Request(ctx context.Context, pid *PID, message any) (interface{}, error)
	Respond(message any)
	System() *System
}

type actorContext struct {
	envelope Envelope
	system   *System
	pid      *PID
	response chan response
}

func (ctx *actorContext) Sender() *PID {
	return ctx.envelope.Sender
}

func (c *actorContext) Message() interface{} {
	if req, ok := c.envelope.Message.(*request); ok {
		return req.message
	}
	return c.envelope.Message
}

func (c *actorContext) Send(pid *PID, message any) {
	c.system.SendWithSender(pid, message, c.pid)
}

type request struct {
	message  interface{}
	response chan response
}

type response struct {
	resp interface{}
	err  error
}

func (c *actorContext) Request(ctx context.Context, pid *PID, message any) (interface{}, error) {
	respc := make(chan response, 1)
	c.system.SendWithSender(pid, &request{message, respc}, c.pid)
	select {
	case resp := <-respc:
		return resp.resp, resp.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (c *actorContext) Respond(message any) {
	if c.envelope.Sender == nil {
		return
	}
	if req, ok := c.envelope.Message.(*request); ok {
		select {
		case req.response <- response{message, nil}:
		default:
		}
	}
	c.system.SendWithSender(c.envelope.Sender, message, c.pid)
}

func (c *actorContext) System() *System {
	return c.system
}
