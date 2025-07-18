package knet

import (
	"errors"
	"log"
	"net"
	"sync"
	"time"
)

var (
	ErrClientClosed = errors.New("knet: Client closed")
)

type TCPClient struct {
	addr  string
	retry bool

	mutex  sync.Mutex
	conn   *TCPConn
	closed bool
}

func NewTCPClient(addr string) *TCPClient {
	client := &TCPClient{
		addr:  addr,
		retry: false,
	}
	return client
}

func (c *TCPClient) EnableRetry()  { c.retry = true }
func (c *TCPClient) DisableRetry() { c.retry = false }

func dialTCP(addr string) (*net.TCPConn, error) {
	raddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}
	return net.DialTCP("tcp", nil, raddr)
}

func (c *TCPClient) DialAndServe(handler TCPHandler, codec Codec) error {
	if c.isClosed() {
		return ErrClientClosed
	}

	if handler == nil {
		handler = DefaultTCPHandler
	}
	if codec == nil {
		codec = DefaultCodec
	}

	var tempDelay time.Duration // how long to sleep on connect failure
	for {
		tcpconn, err := dialTCP(c.addr)
		if err != nil {
			if c.isClosed() {
				return ErrClientClosed
			}
			if !c.retry {
				return err
			}

			if tempDelay == 0 {
				tempDelay = 1 * time.Second
			} else {
				tempDelay *= 2
			}
			if max := 1 * time.Minute; tempDelay > max {
				tempDelay = max
			}
			log.Printf("network: TCPClient dial error: %v; retrying in %v", err, tempDelay)
			time.Sleep(tempDelay)
			continue
		}
		tempDelay = 0

		conn := newTCPConn(tcpconn)
		if err := c.newConn(conn); err != nil {
			conn.Close()
			return err
		}
		if err := c.serveConn(conn, handler, codec); err != nil {
			return err
		}
	}
}

func (c *TCPClient) isClosed() bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.closed
}

func (c *TCPClient) newConn(conn *TCPConn) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.closed {
		return ErrClientClosed
	}
	c.conn = conn
	return nil
}

func (c *TCPClient) serveConn(conn *TCPConn, handler TCPHandler, codec Codec) error {
	conn.serve(handler, codec)
	// remove conn
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.closed {
		return ErrClientClosed
	}
	c.conn = nil
	return nil
}

func (c *TCPClient) GetConn() *TCPConn {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.closed {
		return nil
	}
	return c.conn
}

func (c *TCPClient) Close() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.closed {
		return
	}
	c.closed = true
	if c.conn == nil {
		return
	}
	c.conn.Close()
	c.conn = nil
}
