package nrpc

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type MethodHandler func(srv any, ctx context.Context, dec func(interface{}) error) (interface{}, error)

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

type serviceInfo struct {
	// Contains the implementation for the methods in this service.
	serviceImpl any
	methods     map[string]*MethodDesc
	// streams     map[string]*StreamDesc
	mdata any
}

type Server struct {
	mu   sync.Mutex // guards following
	nc   *nats.Conn
	subj string
	sub  *nats.Subscription

	services map[string]*serviceInfo
}

func NewServer(nc *nats.Conn, subj string) *Server {
	return &Server{nc: nc, subj: subj, services: make(map[string]*serviceInfo)}
}

type ServiceRegistrar interface {
	RegisterService(desc *ServiceDesc, impl any)
}

func (s *Server) RegisterService(sd *ServiceDesc, ss any) {
	if ss != nil {
		ht := reflect.TypeOf(sd.HandlerType).Elem()
		st := reflect.TypeOf(ss)
		if !st.Implements(ht) {
			panic(fmt.Sprintf("nrpc: Server.RegisterService found the handler of type %v that does not satisfy %v", st, ht))
		}
	}
	s.register(sd, ss)
}

func (s *Server) register(sd *ServiceDesc, ss any) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.sub != nil {
		panic(fmt.Sprintf("nrpc: Server.RegisterService after Server.Serve for %q", sd.ServiceName))
	}
	if _, ok := s.services[sd.ServiceName]; ok {
		panic(fmt.Sprintf("grpc: Server.RegisterService found duplicate service registration for %q", sd.ServiceName))
	}
	info := &serviceInfo{
		serviceImpl: ss,
		methods:     make(map[string]*MethodDesc),
		// streams:     make(map[string]*StreamDesc),
		mdata: sd.Metadata,
	}
	for i := range sd.Methods {
		d := &sd.Methods[i]
		info.methods[d.MethodName] = d
	}
	/*
		for i := range sd.Streams {
			d := &sd.Streams[i]
			info.streams[d.StreamName] = d
		}
	*/
	s.services[sd.ServiceName] = info
}

func (s *Server) Start() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	sub, err := s.nc.Subscribe(s.subj, func(msg *nats.Msg) {
		ctx := context.TODO()
		sm := msg.Header.Get("method")
		if sm != "" && sm[0] == '/' {
			sm = sm[1:]
		}
		pos := strings.LastIndex(sm, "/")
		if pos == -1 {
			// FIXME: error
			return
		}
		service := sm[:pos]
		method := sm[pos+1:]
		srv, knownService := s.services[service]
		if knownService {
			if md, ok := srv.methods[method]; ok {
				s.processMsg(ctx, msg, srv, md)
				return
			}
		}
	})

	if err != nil {
		return err
	}
	s.sub = sub
	return nil
}

func (s *Server) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	sub := s.sub
	if sub == nil {
		return nil
	}
	err := sub.Unsubscribe()
	if err != nil {
		return err
	}
	s.sub = nil
	return nil
}

func (s *Server) processMsg(ctx context.Context, msg *nats.Msg, srv *serviceInfo, md *MethodDesc) {
	df := func(v interface{}) error {
		return proto.Unmarshal(msg.Data, v.(proto.Message))
	}
	reply, err := md.Handler(srv.serviceImpl, ctx, df)
	if err != nil {
		// FIXME: err
		return
	}
	data, err := proto.Marshal(reply.(proto.Message))
	if err != nil {
		// FIXME: err
		return
	}
	if err := msg.Respond(data); err != nil {
		// FIXME: err
	}
}
