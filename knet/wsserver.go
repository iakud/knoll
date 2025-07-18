package knet

import (
	"errors"
	"fmt"
	"net/http"
	"sync"

	"golang.org/x/net/websocket"
)

var (
	ErrWSServerClosed = errors.New("knet: WebSocket server closed")
)

type WSServer struct {
	Handler WSHandler
	server  websocket.Server

	mutex  sync.Mutex
	conns  map[*WSConn]struct{}
	closed bool
}

func NewWSServer(handler WSHandler) *WSServer {
	server := &WSServer{
		Handler: handler,
		conns:   make(map[*WSConn]struct{}),
	}
	server.server = websocket.Server{Handler: server.serveWebSocket, Handshake: checkOrigin}
	return server
}

func checkOrigin(config *websocket.Config, req *http.Request) (err error) {
	config.Origin, err = websocket.Origin(config, req)
	if err == nil && config.Origin == nil {
		return fmt.Errorf("null origin")
	}
	return err
}

func (s *WSServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.server.ServeHTTP(w, req)
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

func (s *WSServer) Close() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.closed {
		return
	}
	s.closed = true
	for conn := range s.conns {
		conn.Close()
		delete(s.conns, conn)
	}
}
