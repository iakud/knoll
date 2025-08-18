package knet

import (
	"bufio"
	"errors"
	"log"
	"net"
	"runtime"
	"sync"
	"time"
)

var (
	ErrConnectionPendingSendFull = errors.New("knet: Connection pending send full")
)

type TCPConn struct {
	tcpconn *net.TCPConn

	bufs        [][]byte
	pendingSend int
	mutex       sync.Mutex
	cond        *sync.Cond
	closed      bool

	Userdata interface{}
}

func newTCPConn(tcpconn *net.TCPConn) *TCPConn {
	conn := &TCPConn{
		tcpconn: tcpconn,
	}
	conn.cond = sync.NewCond(&conn.mutex)
	return conn
}

func (c *TCPConn) serve(handler TCPHandler, codec Codec) {
	defer func() {
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			log.Printf("knet: panic serving %v: %v\n%s", c.RemoteAddr(), err, buf)
			c.Close()
		}
	}()

	// start write
	c.startBackgroundWrite(codec)
	defer c.stopBackgroundWrite()
	// conn event
	handler.Connect(c, true)
	defer handler.Connect(c, false)
	// loop read
	r := bufio.NewReader(c.tcpconn)
	for {
		b, err := codec.Read(r)
		if err != nil {
			c.Close()
			return
		}
		handler.Receive(c, b)
	}
}

func (c *TCPConn) startBackgroundWrite(codec Codec) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.closed {
		return
	}
	go c.backgroundWrite(codec)
}

func (c *TCPConn) backgroundWrite(codec Codec) {
	defer func() {
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			log.Printf("knet: panic serving %v: %v\n%s", c.RemoteAddr(), err, buf)
			c.Close()
		}
	}()

	// loop write
	w := bufio.NewWriter(c.tcpconn)
	for closed := false; !closed; {
		var bufs [][]byte

		c.mutex.Lock()
		for !c.closed && len(c.bufs) == 0 {
			c.cond.Wait()
		}
		bufs, c.bufs = c.bufs, bufs // swap
		closed = c.closed
		c.mutex.Unlock()

		for _, b := range bufs {
			if err := codec.Write(w, b); err != nil {
				c.closeWrite()
				c.Close()
				return
			}
		}
		if err := w.Flush(); err != nil {
			c.closeWrite()
			c.Close()
			return
		}
	}
	// not writing now
	c.tcpconn.CloseWrite() // only SHUT_WR
}

func (c *TCPConn) stopBackgroundWrite() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.closed {
		return
	}
	c.closed = true
	c.cond.Signal()
}

func (c *TCPConn) closeWrite() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.closed {
		return
	}
	c.closed = true
}

func (c *TCPConn) LocalAddr() net.Addr {
	return c.tcpconn.LocalAddr()
}

func (c *TCPConn) RemoteAddr() net.Addr {
	return c.tcpconn.RemoteAddr()
}

func (c *TCPConn) SetNoDelay(noDelay bool) error {
	return c.tcpconn.SetNoDelay(noDelay)
}

func (c *TCPConn) SetPendingSend(pendingSend int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.pendingSend = pendingSend
}

func (c *TCPConn) Send(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.closed {
		return nil
	}
	if c.pendingSend > 0 && len(c.bufs) >= c.pendingSend {
		return ErrConnectionPendingSendFull
	}
	c.bufs = append(c.bufs, b)
	c.cond.Signal()
	return nil
}

func (c *TCPConn) Shutdown() {
	c.stopBackgroundWrite() // stop write
}

func (c *TCPConn) Close() error {
	return c.tcpconn.Close()
}

func (c *TCPConn) CloseWithTimeout(timeout time.Duration) {
	time.AfterFunc(timeout, func() {
		c.Close()
	})
}
