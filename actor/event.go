package actor

type RemoteUnreachableEvent struct {
	ListenAddr string
}

type DeadLetterEvent struct {
	Target  *PID
	Message any
	Sender  *PID
}
