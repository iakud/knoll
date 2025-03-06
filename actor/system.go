package actor

import (
	"context"
	"log/slog"
	"strconv"
	"sync/atomic"
)

type Remoter interface {
	Address() string
	Send(pid *PID, message any, sender *PID)
}

type System struct {
	registry  *registry
	remote    Remoter
	address   string
	requestId atomic.Int32
	logger    *slog.Logger
}

func (s *System) Logger() *slog.Logger {
	return s.logger
}

func (s *System) NewLocalPID(id string) *PID {
	return NewPID(s.address, id)
}

func (s *System) Address() string {
	return s.address
}

func NewSystem() *System {
	return NewSystemWithConfig()
}

func NewSystemWithConfig(o ...Option) *System {
	opts := options{}
	for _, option := range o {
		option(&opts)
	}

	system := &System{}
	system.registry = newRegistry()
	system.logger = slog.Default()
	system.address = LocalAddress

	if opts.remote != nil {
		system.remote = opts.remote
		system.address = opts.remote.Address()
	}
	if opts.logger != nil {
		system.logger = opts.logger
	}
	return system
}

func (s *System) Spawn(name string, actor Actor) *PID {
	pid := NewPID(s.address, name)
	context := newContext(pid, s, actor)
	proc := newProcess(context)

	if !s.registry.Add(name, proc) {
		return pid
	}
	proc.Start()
	return pid
}

func (s *System) SpawnFunc(name string, f func(*Context)) *PID {
	return s.Spawn(name, newFuncReceiver(f))
}

func (s *System) Send(pid *PID, message any) {
	s.SendWithSender(pid, message, nil)
}

func (s *System) SendWithSender(pid *PID, message any, sender *PID) {
	if pid == nil {
		return
	}
	if s.address == pid.Address {
		s.SendLocal(pid, message, sender)
		return
	}
	if s.remote == nil {
		return
	}
	s.remote.Send(pid, message, sender)
}

func (s *System) SendLocal(pid *PID, message any, sender *PID) {
	proc := s.registry.Get(pid)
	if proc == nil {
		return
	}
	proc.Send(message, sender)
}

func (s *System) Stop(pid *PID) {
	proc := s.registry.Get(pid)
	if proc == nil {
		return
	}
	proc.Stop()
}

func (s *System) Poison(pid *PID) {
	proc := s.registry.Get(pid)
	if proc == nil {
		return
	}
	proc.Send(poisonPill, nil)
}

func (s *System) Shutdown(ctx context.Context, pid *PID) {
	proc := s.registry.Get(pid)
	if proc == nil {
		return
	}
	proc.Shutdown(ctx)
}

func (s *System) Request(ctx context.Context, pid *PID, message any) (any, error) {
	reqID := NewPID(s.address, "request/"+strconv.Itoa(int(s.requestId.Add(1))))
	req := newRequest()
	s.registry.Add(reqID.ID, req)
	defer s.registry.Remove(reqID)

	s.SendWithSender(pid, message, reqID)
	return req.Result(ctx)
}
