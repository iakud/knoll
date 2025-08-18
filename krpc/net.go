package krpc

type Server interface {
	ListenAndServe() error
	Close() error
	GetConn(id uint64) (Conn, bool)
}

type Client interface {
	DialAndServe() error
	Close() error
	GetConn() (Conn, bool)
}

type Conn interface {
	Id() uint64
	Hash() uint64
	Close() error
	Send(msg Msg) error
	Reply(reqId uint32, msg Msg) error

	SetUserdata(userdata any)
	GetUserdata() any
}

type Handler interface {
	Connect(conn Conn, connected bool)
	Receive(conn Conn, msg Msg)
}
