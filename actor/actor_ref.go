package actor

type ActorRef struct {
	pid string
}

func newActorRef(name string) ActorRef {
	return ActorRef{name}
}
