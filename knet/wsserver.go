package knet

import (
	"errors"
	"net/http"
	"sync"

	"golang.org/x/net/websocket"
)

var (
	ErrWSServerClosed = errors.New("knet: WebSocket server closed")
)

type WSServer struct {
	Handler WSHandler
	server  *http.Server

	mutex  sync.Mutex
	conns  map[*WSConn]struct{}
	closed bool
}

func NewWSServer(addr string, handler WSHandler) *WSServer {
	server := &WSServer{
		Handler: handler,
		conns:   make(map[*WSConn]struct{}),
	}
	server.server = &http.Server{Addr: addr, Handler: websocket.Server{Handler: server.serveWebSocket}}
	return server
}

func (s *WSServer) ListenAndServe() error {
	return s.server.ListenAndServe()
}

func (s *WSServer) serveWebSocket(wsconn *websocket.Conn) {
	handler := s.Handler
	if handler == nil {
		handler = DefaultWSHandler
	}

	conn := newWSConn(wsconn)
	if err := s.newConn(conn); err != nil {
		conn.Close() // close
		return
	}
	s.serveConn(conn, handler)
}

func (s *WSServer) newConn(conn *WSConn) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.closed {
		return ErrWSServerClosed
	}
	s.conns[conn] = struct{}{}
	return nil
}

func (s *WSServer) serveConn(conn *WSConn, handler WSHandler) {
	conn.serve(handler)
	// remove conn
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.closed {
		return
	}
	delete(s.conns, conn)
}

func (s *WSServer) Close() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.closed {
		return nil
	}
	s.closed = true
	err := s.server.Close()
	for conn := range s.conns {
		conn.Close()
		delete(s.conns, conn)
	}
	return err
}
