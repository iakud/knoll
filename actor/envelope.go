package actor

type Envelope struct {
	Message any
	Sender  *PID
}
