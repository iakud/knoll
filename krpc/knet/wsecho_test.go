package knet_test

import (
	"log"
	"testing"

	"github.com/iakud/knoll/krpc/knet"
)

type wsEchoServer struct {
}

func newWSEchoServer() *wsEchoServer {
	echoServer := &wsEchoServer{}
	return echoServer
}

func (srv *wsEchoServer) Connect(connection *knet.WSConn, connected bool) {
	if connected {
		log.Printf("echo server: %v connected.\n", connection.RemoteAddr())
	} else {
		log.Printf("echo server: %v disconnected.\n", connection.RemoteAddr())
	}
}

func (srv *wsEchoServer) Receive(connection *knet.WSConn, data []byte) {
	message := string(data)
	log.Printf("echo server: %v receive %v\n", connection.RemoteAddr(), message)
	log.Println("echo server: send", message)
	connection.Send(data)
	connection.Shutdown()
}

type wsEchoClient struct {
	Client *knet.WSClient
}

func newWSEchoClient() *wsEchoClient {
	echoClient := &wsEchoClient{}
	return echoClient
}

func (c *wsEchoClient) Connect(connection *knet.WSConn, connected bool) {
	const message string = "hello"
	if connected {
		log.Printf("echo client: %v connected.\n", connection.RemoteAddr())
		log.Println("echo client: send", message)
		connection.Send([]byte(message))
	} else {
		log.Printf("echo client: %v disconnected.\n", connection.RemoteAddr())
		c.Client.Close()
	}
}

func (c *wsEchoClient) Receive(connection *knet.WSConn, data []byte) {
	log.Printf("echo client: %v receive %v\n", connection.RemoteAddr(), string(data))
}

func TestWSEcho(t *testing.T) {
	log.Println("test start")
	wsServer := knet.NewWSServer("localhost:8000", newWSEchoServer())
	go func() {
		client := newWSEchoClient()
		wsClient := knet.NewWSClient("ws://localhost:8000", client)
		client.Client = wsClient
		wsClient.EnableRetry() // 启用Retry
		if err := wsClient.DialAndServe(); err != nil {
			log.Println(err)
		}
		wsServer.Close()
		wsServer.Close()
	}()
	wsServer.ListenAndServe()
}
