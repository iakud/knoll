package actor

import (
	"log/slog"
)

type Remoter interface {
	Address() string
	Send(pid *PID, message any, sender *PID)
}

type System struct {
	registry *registry
	remote   Remoter
	address  string
	logger   *slog.Logger
}

func (s *System) Logger() *slog.Logger {
	return s.logger
}

/*
	func (as *ActorSystem) NewLocalPID(id string) *PID {
		return NewPID(as.ProcessRegistry.Address, id)
	}

	func (as *ActorSystem) Address() string {
		return as.ProcessRegistry.Address
	}

	func (as *ActorSystem) GetHostPort() (host string, port int, err error) {
		addr := as.ProcessRegistry.Address
		if h, p, e := net.SplitHostPort(addr); e != nil {
			if addr != localAddress {
				err = e
			}

			host = localAddress
			port = -1
		} else {
			host = h
			port, err = strconv.Atoi(p)
		}

		return
	}
*/

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

func (s *System) Spawn(name string, actor Actor) (*PID, error) {
	pid := NewPID(s.address, name)
	context := newContext(pid, s, actor)
	proc := newProcess(context)

	if s.registry.Add(name, proc) {
		return pid, nil
	}
	proc.Start()
	return pid, nil
}

func (s *System) SpawnFunc(name string, f func(*Context)) (*PID, error) {
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
		s.send(pid, message, sender)
		return
	}
	if s.remote == nil {
		return
	}
	s.remote.Send(pid, message, sender)
}

func (s *System) send(pid *PID, message any, sender *PID) {
	proc := s.registry.Get(pid)
	if proc == nil {
		return
	}
	proc.Send(message, sender)
}

func (s *System) Poison(pid *PID) {
	s.SendWithSender(pid, poisonPill, nil)
}

func (s *System) Stop(pid *PID) {
	proc := s.registry.Get(pid)
	if proc == nil {
		return
	}
	proc.Stop()
}
