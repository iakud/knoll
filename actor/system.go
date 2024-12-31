package actor

import (
	"log/slog"
	"sync"
)

type System struct {
	registry *registry
	lock     sync.RWMutex
	stop     chan struct{}
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
func (s *System) Stop() {
	close(s.stop)
}

func (s *System) IsStopped() bool {
	select {
	case <-s.stop:
		return true
	default:
		return false
	}
}

func NewSystem() *System {
	return NewSystemWithConfig()
}

func NewSystemWithConfig() *System {
	system := &System{}
	system.registry = newRegistry()
	system.logger = slog.Default()
	system.stop = make(chan struct{})
	return system
}

func (s *System) Spawn(name string, actor Actor) (*PID, error) {
	const address = "nohost"
	pid := NewPID(address, name)
	proc := newProcess(pid, s, actor)

	s.registry.Add(name, proc)
	proc.Start()
	return pid, nil
}

func (s *System) Context() Context {
	return &actorContext{
		envelope: Envelope{},
		system:   s,
		pid:      nil,
	}
}

func (s *System) Send(pid *PID, message any, sender *PID) {
	proc := s.registry.Get(pid)
	if proc == nil {
		return
	}
	proc.Send(message, sender)
}

func (s *System) Poison(pid *PID) {
	s.Send(pid, poisonPillMessage, nil)
}
