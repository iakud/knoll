package krpc

import "slices"

type ClientMsg struct {
	Header
	payload []byte
}

func (m *ClientMsg) Payload() []byte {
	return m.payload
}

func (m *ClientMsg) SetPayload(payload []byte) {
	m.payload = payload
}

func (m *ClientMsg) Size() int {
	return HeaderSize + len(m.payload)
}

func (m *ClientMsg) Marshal(buf []byte) (int, error) {
	n, err := m.Header.Marshal(buf)
	if err != nil {
		return n, err
	}

	n += copy(buf[n:], m.payload)
	return n, nil
}

func (m *ClientMsg) Unmarshal(buf []byte) (int, error) {
	n, err := m.Header.Unmarshal(buf)
	if err != nil {
		return n, err
	}

	m.payload = slices.Clone(buf[n:])
	n += len(m.payload)
	return n, nil
}
