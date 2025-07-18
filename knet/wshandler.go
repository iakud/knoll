package knet

type WSHandler interface {
	Connect(conn *WSConn, connected bool)
	Receive(conn *WSConn, data []byte)
}

type defaultWSHandler struct {
}

func (*defaultWSHandler) Connect(*WSConn, bool) {

}

func (*defaultWSHandler) Receive(*WSConn, []byte) {

}

var DefaultWSHandler = &defaultWSHandler{}
