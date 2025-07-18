package knet

import (
	"errors"
	"log"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

var (
	ErrWSClientClosed = errors.New("knet: Websocket client closed")
)

type WSClient struct {
	Url     string
	Handler WSHandler
	retry   bool

	mutex  sync.Mutex
	conn   *WSConn
	closed bool
}

func NewWSClient(url string, handler WSHandler) *WSClient {
	client := &WSClient{
		Url:     url,
		Handler: handler,
	}
	return client
}

func (c *WSClient) EnableRetry()  { c.retry = true }
func (c *WSClient) DisableRetry() { c.retry = false }

func DialAndServeWS(url string, handler WSHandler) error {
	client := &WSClient{Url: url, Handler: handler}
	return client.DialAndServe()
}

func (c *WSClient) DialAndServe() error {
	if c.isClosed() {
		return ErrWSClientClosed
	}

	handler := c.Handler
	if handler == nil {
		handler = DefaultWSHandler
	}

	var tempDelay time.Duration // how long to sleep on connect failure
	for {
		wsconn, err := websocket.Dial(c.Url, "", "http://localhost/")
		if err != nil {
			if c.isClosed() {
				return ErrWSClientClosed
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
			log.Printf("knet: WebSocket client dial error: %v; retrying in %v", err, tempDelay)
			time.Sleep(tempDelay)
			continue
		}
		tempDelay = 0

		conn := newWSConn(wsconn)
		if err := c.newConn(conn); err != nil {
			conn.Close()
			return err
		}
		if err := c.serveConn(conn, handler); err != nil {
			return err
		}
	}
}

func (c *WSClient) isClosed() bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.closed
}

func (c *WSClient) newConn(conn *WSConn) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.closed {
		return ErrWSClientClosed
	}
	c.conn = conn
	return nil
}

func (c *WSClient) serveConn(conn *WSConn, handler WSHandler) error {
	conn.serve(handler)
	// remove conn
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.closed {
		return ErrWSClientClosed
	}
	c.conn = nil
	return nil
}

func (c *WSClient) GetConn() *WSConn {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.closed {
		return nil
	}
	return c.conn
}

func (c *WSClient) Close() {
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
