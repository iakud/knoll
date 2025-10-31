package krpcnet

import (
	"encoding/binary"
	"errors"
	"net/netip"
	"slices"
	"sync"
)

type Msg interface {
	Header() *Header
	Size() int
	ConnId() uint64
	SetConnId(connId uint64)
	ConnIP() netip.Addr
	SetConnIP(ip netip.Addr)
	UserId() uint64
	SetUserId(userId uint64)
	Payload() []byte
	SetPayload(payload []byte)
	Marshal(buf []byte) (int, error)
	Unmarshal(buf []byte) (int, error)
}

func NewMsg(backend bool) Msg {
	if backend {
		return NewBackendMsg()
	}
	return NewUserMsg()
}

var _ Msg = (*userMsg)(nil)

var userMsgPool = sync.Pool{New: func() any { return &userMsg{} }}

func NewUserMsg() Msg {
	return userMsgPool.Get().(Msg)
}

type userMsg struct {
	header  Header
	payload []byte
}

func (m *userMsg) Header() *Header {
	return &m.header
}

func (m *userMsg) ConnId() uint64 {
	return 0
}

func (m *userMsg) SetConnId(connId uint64) {
}

func (m *userMsg) ConnIP() netip.Addr {
	return netip.Addr{}
}

func (m *userMsg) SetConnIP(ip netip.Addr) {
}

func (m *userMsg) UserId() uint64 {
	return 0
}

func (m *userMsg) SetUserId(userId uint64) {
}

func (m *userMsg) Payload() []byte {
	return m.payload
}

func (m *userMsg) SetPayload(payload []byte) {
	m.payload = payload
}

func (m *userMsg) Size() int {
	return HeaderSize + len(m.payload)
}

func (m *userMsg) Marshal(buf []byte) (int, error) {
	if len(buf) < HeaderSize+len(m.payload) {
		return 0, errors.New("buffer too small")
	}
	n, err := m.header.Marshal(buf)
	if err != nil {
		return n, err
	}

	n += copy(buf[n:], m.payload)
	return n, nil
}

func (m *userMsg) Unmarshal(buf []byte) (int, error) {
	if len(buf) < HeaderSize {
		return 0, errors.New("buffer too small")
	}
	n, err := m.header.Unmarshal(buf)
	if err != nil {
		return n, err
	}

	m.payload = slices.Clone(buf[n:])
	n += len(m.payload)
	return n, nil
}

const (
	kMsgConnIdSize   = 8
	kMsgConnAddrSize = 16
	kMsgUserIdSize   = 8
)

var _ Msg = (*backendMsg)(nil)

var backendMsgPool = sync.Pool{New: func() any { return &backendMsg{} }}

func NewBackendMsg() Msg {
	return backendMsgPool.Get().(Msg)
}

type backendMsg struct {
	header  Header
	connId  uint64
	connIP  [16]byte
	userId  uint64
	payload []byte
}

func (m *backendMsg) Header() *Header {
	return &m.header
}

func (m *backendMsg) ConnId() uint64 {
	return m.connId
}

func (m *backendMsg) SetConnId(connId uint64) {
	m.connId = connId
}

func (m *backendMsg) ConnIP() netip.Addr {
	return netip.AddrFrom16(m.connIP)
}

func (m *backendMsg) SetConnIP(ip netip.Addr) {
	m.connIP = ip.As16()
}

func (m *backendMsg) UserId() uint64 {
	return m.userId
}

func (m *backendMsg) SetUserId(userId uint64) {
	m.userId = userId
}

func (m *backendMsg) Payload() []byte {
	return m.payload
}

func (m *backendMsg) SetPayload(payload []byte) {
	m.payload = payload
}

func (m *backendMsg) Size() int {
	return HeaderSize + kMsgConnIdSize + kMsgConnAddrSize + kMsgUserIdSize + len(m.payload)
}

func (m *backendMsg) Marshal(buf []byte) (int, error) {
	if len(buf) < HeaderSize+kMsgConnIdSize+kMsgConnAddrSize+kMsgUserIdSize+len(m.payload) {
		return 0, errors.New("buffer too small")
	}
	n, err := m.header.Marshal(buf)
	if err != nil {
		return n, err
	}

	binary.BigEndian.PutUint64(buf[n:], m.connId)
	n += kMsgConnIdSize
	copy(buf[n:], m.connIP[:])
	n += kMsgConnAddrSize
	binary.BigEndian.PutUint64(buf[n:], m.userId)
	n += kMsgUserIdSize

	n += copy(buf[n:], m.payload)
	return n, nil
}

func (m *backendMsg) Unmarshal(buf []byte) (int, error) {
	if len(buf) < HeaderSize+kMsgConnIdSize+kMsgConnAddrSize+kMsgUserIdSize {
		return 0, errors.New("buffer too small")
	}
	n, err := m.header.Unmarshal(buf)
	if err != nil {
		return n, err
	}

	m.connId = binary.BigEndian.Uint64(buf[n:])
	n += kMsgConnIdSize
	copy(m.connIP[:], buf[n:])
	n += kMsgConnAddrSize
	m.userId = binary.BigEndian.Uint64(buf[n:])
	n += kMsgUserIdSize

	m.payload = slices.Clone(buf[n:])
	n += len(m.payload)
	return n, nil
}
