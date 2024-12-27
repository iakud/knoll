package actor

import "sync"

type registry struct {
	localActors sync.Map
}

func (r *registry) Add(id string, mailbox *mailbox) (*PID, bool) {
	if _, ok := r.localActors.LoadOrStore(id, mailbox); ok {
		return ActorRef{id}, ok
	}
	return ActorRef{id}, false
}

func (r *registry) Remove(ref ActorRef) {
	r.localActors.Delete(ref.id)
}

func (r *registry) Get(ref ActorRef) (*mailbox, bool) {
	mailbox, ok := r.localActors.Load(ref.id)
	return mailbox, ok
}
