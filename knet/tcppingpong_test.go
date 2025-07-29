package knet_test

import (
	"log"
	"sync/atomic"
	"testing"
	"time"

	"github.com/iakud/knoll/knet"
)

const (
	kBlockSize   = 1024 * 16
	kClientCount = 10
	kTimeout     = time.Second * 10
)

// Pingpong Server
type pingpongServer struct {
	server *knet.TCPServer
}

func newPingpongServer(addr string) *pingpongServer {
	server := &pingpongServer{}
	server.server = knet.NewTCPServer(addr, server, nil)
	return server
}

func (s *pingpongServer) ListenAndServe() {
	if err := s.server.ListenAndServe(); err != nil {
		if err == knet.ErrServerClosed {
			return
		}
		log.Println(err.Error())
	}
}

func (s *pingpongServer) Close() {
	s.server.Close()
}

func (s *pingpongServer) Connect(connection *knet.TCPConn, connected bool) {
	if connected {
		connection.SetNoDelay(true)
	}
}

func (s *pingpongServer) Receive(connection *knet.TCPConn, b []byte) {
	connection.Send(b)
}

// Pingpong Client
type pingpongClient struct {
	clients    []*knet.TCPClient
	message    []byte
	nConnected int32

	bytesRead    int64
	messagesRead int64
	done         chan struct{}
}

func newPingpongClient(addr string) *pingpongClient {
	// build message
	message := make([]byte, kBlockSize)
	for i := 0; i < kBlockSize; i++ {
		message[i] = byte(i % 128)
	}
	c := &pingpongClient{
		message: message,
		done:    make(chan struct{}),
	}
	clients := make([]*knet.TCPClient, kClientCount)
	for i := 0; i < kClientCount; i++ {
		client := knet.NewTCPClient(addr, c, nil)
		go c.serveClient(client)
		clients[i] = client
	}
	c.clients = clients
	time.AfterFunc(kTimeout, c.handleTimeout)
	return c
}

func (c *pingpongClient) serveClient(client *knet.TCPClient) {
	client.EnableRetry() // 启用retry
	if err := client.DialAndServe(); err != nil {
		if err == knet.ErrClientClosed {
			return
		}
		log.Println(err.Error())
	}
}

func (c *pingpongClient) handleTimeout() {
	for _, client := range c.clients {
		client.Close()
	}
}

func (c *pingpongClient) Connect(connection *knet.TCPConn, connected bool) {
	if connected {
		connection.SetNoDelay(true)
		connection.Send(c.message)
		if atomic.AddInt32(&c.nConnected, 1) != kClientCount {
			return
		}
		log.Println("all connected")
	} else {
		if atomic.AddInt32(&c.nConnected, -1) != 0 {
			return
		}
		bytesRead := atomic.LoadInt64(&c.bytesRead)
		messagesRead := atomic.LoadInt64(&c.messagesRead)

		log.Println("total bytes read", bytesRead)
		log.Println("total messages read", messagesRead)
		log.Println("average message size", bytesRead/messagesRead)
		timeout := int64(kTimeout / time.Second)
		log.Println("MiB/s throughput", bytesRead/(timeout*1024*1024))
		close(c.done)
	}
}

func (c *pingpongClient) Receive(connection *knet.TCPConn, b []byte) {
	connection.Send(b)
	atomic.AddInt64(&c.messagesRead, 1)
	atomic.AddInt64(&c.bytesRead, int64(len(b)))
}

func (c *pingpongClient) Done() {
	<-c.done
}

func TestPingpong(t *testing.T) {
	server := newPingpongServer("localhost:8000")
	go server.ListenAndServe()
	defer server.Close()
	time.Sleep(time.Second)
	client := newPingpongClient("localhost:8000")
	client.Done()
}
