package krpc

type Msg interface {
	FlagRequest() bool
	setFlagRequest()
	FlagReply() bool
	setFlagReply()
	MsgId() uint16
	SetMsgId(msgId uint16)
	ReqId() uint32
	setReqId(reqId uint32)
	Size() int
	Marshal(buf []byte) (int, error)
	Unmarshal(buf []byte) (int, error)
}

func Marshal(msg Msg) ([]byte, error) {
	buf := make([]byte, msg.Size())
	if _, err := msg.Marshal(buf); err != nil {
		return nil, err
	}
	return buf, nil
}

func Unmarshal(buf []byte, msg Msg) error {
	if _, err := msg.Unmarshal(buf); err != nil {
		return err
	}
	return nil
}
