package actor

type DeadLetterEvent struct {
	Target  *PID
	Message any
	Sender  *PID
}
