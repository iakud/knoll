package knet

import (
	"errors"
	"net"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

var (
	ErrWSConnectionPendingSendFull = errors.New("knet: WebSocket connection pending send full")
)

type WSConn struct {
	wsconn *websocket.Conn

	bufs        [][]byte
	pendingSend int
	mutex       sync.Mutex
	cond        *sync.Cond
	closed      bool
}

func newWSConn(wsconn *websocket.Conn) *WSConn {
	conn := &WSConn{wsconn: wsconn}
	conn.cond = sync.NewCond(&conn.mutex)
	return conn
}

func (c *WSConn) serve(handler WSHandler) {
	defer c.wsconn.Close()

	// start write
	c.startBackgroundWrite()
	defer c.stopBackgroundWrite()

	// conn event
	handler.Connect(c, true)
	defer handler.Connect(c, false)
	for {
		var data []byte
		if err := websocket.Message.Receive(c.wsconn, &data); err != nil {
			c.wsconn.Close()
			break
		}
		handler.Receive(c, data)
	}
}

func (c *WSConn) startBackgroundWrite() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.closed {
		return
	}
	go c.backgroundWrite()
}

func (c *WSConn) backgroundWrite() {
	for closed := false; !closed; {
		var bufs [][]byte

		c.mutex.Lock()
		for !c.closed && len(c.bufs) == 0 {
			c.cond.Wait()
		}
		bufs, c.bufs = c.bufs, bufs // swap
		closed = c.closed
		c.mutex.Unlock()

		for _, message := range bufs {
			if err := websocket.Message.Send(c.wsconn, message); err != nil {
				c.closeWrite()
				c.wsconn.Close()
				return
			}
		}
	}
	// not writing now
	c.wsconn.Close()
}

func (c *WSConn) stopBackgroundWrite() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.closed {
		return
	}
	c.closed = true
	c.cond.Signal()
}

func (c *WSConn) closeWrite() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.closed {
		return
	}
	c.closed = true
}

func (c *WSConn) LocalAddr() net.Addr {
	return c.wsconn.LocalAddr()
}

func (c *WSConn) RemoteAddr() net.Addr {
	return c.wsconn.RemoteAddr()
}

func (c *WSConn) SetPendingSend(pendingSend int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.pendingSend = pendingSend
}

func (c *WSConn) Send(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.closed {
		return nil
	}
	if c.pendingSend > 0 && len(c.bufs) >= c.pendingSend {
		return ErrWSConnectionPendingSendFull
	}
	c.bufs = append(c.bufs, data)
	c.cond.Signal()
	return nil
}

func (c *WSConn) Shutdown() {
	c.stopBackgroundWrite()
}

func (c *WSConn) Close() {
	c.wsconn.Close()
}
func (c *WSConn) CloseWithTimeout(timeout time.Duration) {
	time.AfterFunc(timeout, c.Close)
}
