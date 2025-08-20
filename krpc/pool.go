package krpc

import "sync"

type messagePool struct {
	pool sync.Pool
}

func (p *messagePool) Get() Message {
	return p.pool.Get().(Message)
}

func (p *messagePool) Put(m Message) {
	p.pool.Put(m)
}

var BMessagePool = messagePool{
	pool: sync.Pool{
		New: func() any {
			return &BMessage{}
		},
	},
}

var CMessagePool = messagePool{
	pool: sync.Pool{
		New: func() any {
			return &CMessage{}
		},
	},
}
