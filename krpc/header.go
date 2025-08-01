package krpc

import (
	"encoding/binary"
	"errors"
)

type MsgFlag uint16

const MF_None MsgFlag = 0x00

const (
	MF_Request MsgFlag = 0x01 << iota
	MF_Reply
)

const (
	OffsetFlag  = 0
	OffsetMsgID = 2
	OffsetReqID = 4

	HeaderSize = 8
)

type Header struct {
	flag  uint16
	msgId uint16
	reqId uint32
}

func (h *Header) FlagRequest() bool {
	return h.flag&uint16(MF_Request) != 0
}

func (h *Header) setFlagRequest() {
	h.flag |= uint16(MF_Request)
}

func (h *Header) ClearFlagRequest() {
	h.flag &^= uint16(MF_Request)
}

func (h *Header) FlagReply() bool {
	return h.flag&uint16(MF_Reply) != 0
}

func (h *Header) setFlagReply() {
	h.flag |= uint16(MF_Reply)
}

func (h *Header) ClearFlagReply() {
	h.flag &^= uint16(MF_Reply)
}

func (h *Header) MsgId() uint16 {
	return h.msgId
}

func (h *Header) SetMsgId(msgId uint16) {
	h.msgId = msgId
}

func (h *Header) ReqId() uint32 {
	return h.reqId
}

func (h *Header) setReqId(reqId uint32) {
	h.reqId = reqId
}

func (h *Header) Marshal(buf []byte) (int, error) {
	if len(buf) < HeaderSize {
		return 0, errors.New("buffer too small")
	}
	binary.BigEndian.PutUint16(buf[OffsetFlag:], h.flag)
	binary.BigEndian.PutUint16(buf[OffsetMsgID:], h.msgId)
	binary.BigEndian.PutUint32(buf[OffsetReqID:], h.reqId)
	return HeaderSize, nil
}

func (h *Header) Unmarshal(buf []byte) (int, error) {
	if len(buf) < HeaderSize {
		return 0, errors.New("buffer too small")
	}
	h.flag = binary.BigEndian.Uint16(buf[OffsetFlag:])
	h.msgId = binary.BigEndian.Uint16(buf[OffsetMsgID:])
	h.reqId = binary.BigEndian.Uint32(buf[OffsetReqID:])
	return HeaderSize, nil
}
