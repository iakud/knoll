package knet

type TCPHandler interface {
	Connect(conn *TCPConn, connected bool)
	Receive(conn *TCPConn, data []byte)
}

type defaultTCPHandler struct {
}

func (*defaultTCPHandler) Connect(*TCPConn, bool) {

}

func (*defaultTCPHandler) Receive(*TCPConn, []byte) {

}

var DefaultTCPHandler = &defaultTCPHandler{}
