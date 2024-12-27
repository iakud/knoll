package actor

import (
	"context"
)

type Actor[T any] interface {
	OnStart()
	OnClose()
	Receive(ctx context.Context, message T)
}

/*
type Actor struct {
	Context interface{}

	mutex    sync.Mutex
	cond     *sync.Cond
	functors []func()
	closed   bool
}

func New() *Actor {
	a := new(Actor)
	a.cond = sync.NewCond(&a.mutex)
	return a
}

func (a *Actor) RunInActor(functor func()) {
	a.mutex.Lock()
	a.functors = append(a.functors, functor)
	a.mutex.Unlock()

	a.cond.Signal()
}

func (a *Actor) Loop() {
	defer func() {
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			// log.Printf("eventloop: panic loop: %v\n%s", err, buf)
		}
	}()
	var closed bool
	for !closed {
		var functors []func()
		a.mutex.Lock()
		for !a.closed && len(a.functors) == 0 {
			a.cond.Wait()
		}
		functors, a.functors = a.functors, nil // swap
		closed = a.closed
		a.mutex.Unlock()

		for _, functor := range functors {
			functor()
		}
	}
}

func (a *Actor) Close() {
	a.mutex.Lock()
	if a.closed {
		a.mutex.Unlock()
		return
	}
	a.closed = true
	a.mutex.Unlock()

	a.cond.Signal()
}
*/
