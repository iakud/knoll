package nrpc

import (
	"context"
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
			return
			// panic("nrpc: Server.RegisterService found the handler of type %v that does not satisfy %v", st, ht)
		}
	}
	s.register(sd, ss)
}

func (s *Server) register(sd *ServiceDesc, ss any) {
	s.mu.Lock()
	defer s.mu.Unlock()
	/*
		s.printf("RegisterService(%q)", sd.ServiceName)
		if s.serve {
			panic("nrpc: Server.RegisterService after Server.Serve for %q", sd.ServiceName)
		}
	*/
	if _, ok := s.services[sd.ServiceName]; ok {
		return
		// fmt.fprintf("grpc: Server.RegisterService found duplicate service registration for %q", sd.ServiceName)
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
		}
	})
	_ = sub
	// sub.Unsubscribe()
	if err != nil {
		return err
	}
	return nil
}
