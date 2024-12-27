package actor

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
)

type ActorSystem[T any] struct {
	lock     sync.RWMutex
	mailboxs map[string]*mailbox[T]
	stop     chan struct{}
	logger   *slog.Logger
}

func (as *ActorSystem[T]) Logger() *slog.Logger {
	return as.logger
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
func (as *ActorSystem[T]) Stop() {
	close(as.stop)
}

func (as *ActorSystem[T]) IsStopped() bool {
	select {
	case <-as.stop:
		return true
	default:
		return false
	}
}

func NewActorSystem[T any]() *ActorSystem[T] {
	return NewActorSystemWithConfig[T]()
}

func NewActorSystemWithConfig[T any]() *ActorSystem[T] {
	system := &ActorSystem[T]{}
	system.mailboxs = make(map[string]*mailbox[T])
	system.logger = slog.Default()
	system.stop = make(chan struct{})
	return system
}

func (as *ActorSystem[T]) Spawn(name string, actor Actor[T]) (ActorRef, error) {
	actorRef := newActorRef(name)
	mailbox := NewMailbox[T]()

	as.lock.Lock()
	defer as.lock.Unlock()
	if _, ok := as.mailboxs[name]; ok {
		return actorRef, nil
	}
	as.mailboxs[name] = mailbox
	go receive(mailbox, actor)
	return actorRef, nil
}

func receive[T any](mailbox *mailbox[T], actor Actor[T]) {
	defer func() {
		if r := recover(); r != nil {
			slog.Info("Receive recovered in %v", r)
		}
	}()
	actor.OnStart()
	defer actor.OnClose()

	for {
		select {
		case message := <-mailbox.MessageC():
			actor.Receive(context.Background(), message)
		case <-mailbox.Done():
			return
		}
	}
}

func (as *ActorSystem[T]) Send(ctx context.Context, actorRef ActorRef, message T) error {
	mailbox, ok := as.mailboxs[actorRef.id]
	if !ok {
		return fmt.Errorf("Mailbox failed")
	}
	return mailbox.Send(ctx, message)
}

func (as *ActorSystem[T]) Close(actorRef ActorRef) {
	mailbox, ok := as.mailboxs[actorRef.id]
	if !ok {
		return
	}
	mailbox.Close()
	delete(as.mailboxs, actorRef.id)
}
