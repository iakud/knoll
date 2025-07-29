package krpc

type Conn interface {
	Send(data []byte) error
}

func Send(conn Conn, msg Msg) error {
	buf := make([]byte, msg.Size())
	if _, err := msg.Marshal(buf); err != nil {
		return err
	}
	return conn.Send(buf)
}

func SendOK(conn Conn) error {
	return nil
}

func SendError(conn Conn, err error) error {
	return nil
}

func Reply(conn Conn, reqId uint32, msg Msg) error {
	msg.setFlagReply()
	msg.setReqId(reqId)
	return Send(conn, msg)
}

func ReplyOK(conn Conn, reqId uint32) error {
	return nil
}

func ReplyError(conn Conn, reqId uint32, err error) error {
	return nil
}
