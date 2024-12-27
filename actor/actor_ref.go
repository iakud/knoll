package actor

type ActorRef struct {
	id string
}

func newActorRef(name string) ActorRef {
	return ActorRef{name}
}
