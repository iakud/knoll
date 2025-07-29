package knet

import (
	"net"
	"time"
)

type Conn interface {
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	SetPendingSend(pendingSend int)
	Send(data []byte) error
	Shutdown()
	Close()
	CloseWithTimeout(timeout time.Duration)
}
