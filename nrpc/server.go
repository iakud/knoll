package nrpc

import (
	"context"

	"github.com/nats-io/nats.go"
)

type MethodHandler func(srv any, ctx context.Context, dec func(any) error) (any, error)

type MethodDesc struct {
	MethodName string
	Handler    MethodHandler
}

type ServiceDesc struct {
	ServiceName string
	// The pointer to the service interface. Used to check whether the user
	// provided implementation satisfies the interface requirements.
	HandlerType any
	Methods     []MethodDesc
	//	Streams     []StreamDesc
	Metadata any
}

type Server struct {
	nc *nats.Conn
}

func NewServer(nc *nats.Conn) *Server {
	return &Server{nc}
}

type ServiceRegistrar interface {
	RegisterService(desc *ServiceDesc, impl any)
}

func (s *Server) RegisterService(sd *ServiceDesc, ss any) {

}
