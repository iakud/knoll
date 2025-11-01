package knet

import (
	"errors"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	ErrWSServerClosed = errors.New("knet: WebSocket server closed")
)

type WSServer struct {
	Handler  WSHandler
	upgrader websocket.Upgrader
	server   *http.Server

	mutex  sync.Mutex
	conns  map[*WSConn]struct{}
	closed bool
}

func NewWSServer(addr string, handler WSHandler) *WSServer {
	server := &WSServer{
		Handler: handler,
		conns:   make(map[*WSConn]struct{}),
	}
	server.server = &http.Server{Addr: addr, Handler: http.HandlerFunc(server.serveWebSocket)}
	return server
}

func (s *WSServer) ListenAndServe() error {
	return s.server.ListenAndServe()
}

func (s *WSServer) serveWebSocket(w http.ResponseWriter, r *http.Request) {
	wsconn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("knet: WSServer upgrade error: %v", err)
		return
	}

	conn := newWSConn(wsconn)
	if err := s.newConn(conn); err != nil {
		conn.Close() // close
		return
	}

	handler := s.Handler
	if handler == nil {
		handler = DefaultWSHandler
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
