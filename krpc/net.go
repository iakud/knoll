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
	setHash(hash uint64)
	Hash() uint64
	Close() error
	Send(msg Message) error
	Reply(reqId uint32, reply Message) error
	NewMessage() Message

	SetUserdata(userdata any)
	GetUserdata() any
}

type Handler interface {
	Connect(conn Conn, connected bool)
	Receive(conn Conn, message Message)
	Handshake(conn Conn, hash uint64)
	UserOnline(conn Conn, userId uint64)
}

type Codec interface {
	Unmarshal(data []byte) (Message, error)
	Marshal(m Message) ([]byte, error)
}
