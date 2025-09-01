package krpcnet

import (
	"encoding/binary"
	"errors"
)

type msgFlag uint16

// const mf_none msgFlag = 0x00

const (
	mf_request msgFlag = 0x01 << iota
	mf_reply
)

const (
	kOffsetFlag  = 0
	kOffsetMsgId = 2
	kOffsetReqId = 4

	HeaderSize = 8
)

type Header struct {
	flag  uint16
	msgId uint16
	reqId uint32
}

func (h *Header) FlagRequest() bool {
	return h.flag&uint16(mf_request) != 0
}

func (h *Header) SetFlagRequest() {
	h.flag |= uint16(mf_request)
}

func (h *Header) ClearFlagRequest() {
	h.flag &^= uint16(mf_request)
}

func (h *Header) FlagReply() bool {
	return h.flag&uint16(mf_reply) != 0
}

func (h *Header) SetFlagReply() {
	h.flag |= uint16(mf_reply)
}

func (h *Header) ClearFlagReply() {
	h.flag &^= uint16(mf_reply)
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

func (h *Header) SetReqId(reqId uint32) {
	h.reqId = reqId
}

func (h *Header) Marshal(buf []byte) (int, error) {
	if len(buf) < HeaderSize {
		return 0, errors.New("buffer too small")
	}
	binary.BigEndian.PutUint16(buf[kOffsetFlag:], h.flag)
	binary.BigEndian.PutUint16(buf[kOffsetMsgId:], h.msgId)
	binary.BigEndian.PutUint32(buf[kOffsetReqId:], h.reqId)
	return HeaderSize, nil
}

func (h *Header) Unmarshal(buf []byte) (int, error) {
	if len(buf) < HeaderSize {
		return 0, errors.New("buffer too small")
	}
	h.flag = binary.BigEndian.Uint16(buf[kOffsetFlag:])
	h.msgId = binary.BigEndian.Uint16(buf[kOffsetMsgId:])
	h.reqId = binary.BigEndian.Uint32(buf[kOffsetReqId:])
	return HeaderSize, nil
}
