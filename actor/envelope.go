package actor

type Envelope[T any] struct {
	sender   ActorRef
	receiver ActorRef

	message T
}

func NewEnvelope[T any](sender, receiver ActorRef, message T) {

}
