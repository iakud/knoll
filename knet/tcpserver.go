package knet

import (
	"errors"
	"log"
	"net"
	"sync"
	"time"
)

var (
	ErrServerClosed = errors.New("knet: Server closed")
)

type TCPServer struct {
	addr    string
	handler TCPHandler
	codec   Codec

	mutex    sync.Mutex
	listener *net.TCPListener
	conns    map[*TCPConn]struct{}
	closed   bool
}

func NewTCPServer(addr string, handler TCPHandler, codec Codec) *TCPServer {
	server := &TCPServer{
		addr:    addr,
		handler: handler,
		codec:   codec,
	}
	return server
}

func listenTCP(addr string) (*net.TCPListener, error) {
	if addr == "" {
		addr = ":0"
	}
	laddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}
	return net.ListenTCP("tcp", laddr)
}

func (s *TCPServer) ListenAndServe() error {
	if s.isClosed() {
		return ErrServerClosed
	}
	ln, err := listenTCP(s.addr)
	if err != nil {
		return err
	}

	defer ln.Close()

	if err := s.newListener(ln); err != nil {
		return err
	}

	handler := s.handler
	if handler == nil {
		handler = DefaultTCPHandler
	}
	codec := s.codec
	if codec == nil {
		codec = DefaultCodec
	}

	var tempDelay time.Duration // how long to sleep on accept failure
	for {
		wsconn, err := ln.AcceptTCP()
		if err != nil {
			if s.isClosed() {
				return ErrServerClosed
			}
			if ne, ok := err.(net.Error); ok && ne.Timeout() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				log.Printf("knet: TCPServer accept error: %v; retrying in %v", err, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			log.Printf("knet: TCPServer error: %v", err)
			return err
		}
		tempDelay = 0

		conn := newTCPConn(wsconn)
		if err := s.newConn(conn); err != nil {
			conn.Close() // close
			return err
		}
		go s.serveConn(conn, handler, codec)
	}
}

func (s *TCPServer) isClosed() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.closed
}

func (s *TCPServer) newListener(ln *net.TCPListener) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.closed {
		return ErrServerClosed
	}
	s.listener = ln
	s.conns = make(map[*TCPConn]struct{})
	return nil
}

func (s *TCPServer) newConn(conn *TCPConn) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.closed {
		return ErrServerClosed
	}
	s.conns[conn] = struct{}{}
	return nil
}

func (s *TCPServer) serveConn(conn *TCPConn, handler TCPHandler, codec Codec) {
	conn.serve(handler, codec)
	// remove conn
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.closed {
		return
	}
	delete(s.conns, conn)
}

func (s *TCPServer) Close() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.closed {
		return nil
	}
	s.closed = true
	if s.listener == nil {
		return nil
	}
	err := s.listener.Close()
	s.listener = nil
	for conn := range s.conns {
		conn.Close()
		delete(s.conns, conn)
	}
	return err
}
