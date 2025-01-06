package actor

type Actor interface {
	Receive(ctx *Context)
}

type funcReceiver func(*Context)

func (fr funcReceiver) Receive(c *Context) {
	fr(c)
}

func newFuncReceiver(f func(*Context)) Actor {
	return funcReceiver(f)
}
