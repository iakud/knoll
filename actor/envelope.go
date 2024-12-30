package actor

type Envelope struct {
	Message interface{}
	Sender  ActorRef
}

func NewEnvelope(sender ActorRef, message interface{}) {

}
