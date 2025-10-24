package nrpc

import (
	"context"
	"fmt"
	"log/slog"
	"reflect"
	"strconv"
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
		panic(fmt.Sprintf("nrpc: Server.RegisterService found duplicate service registration for %q", sd.ServiceName))
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
	sub, err := s.nc.Subscribe(s.subj, s.handleMsg)

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

func (s *Server) handleMsg(msg *nats.Msg) {
	ctx := context.TODO()
	sm := msg.Header.Get(methodHdr)
	if sm != "" && sm[0] == '/' {
		sm = sm[1:]
	}
	pos := strings.LastIndex(sm, "/")
	if pos == -1 {
		errDesc := fmt.Sprintf("malformed method name: %q", sm)
		s.respondError(msg.Reply, Unimplemented, errDesc)
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
}

func (s *Server) processMsg(ctx context.Context, msg *nats.Msg, srv *serviceInfo, md *MethodDesc) error {
	df := func(v interface{}) error {
		if err := proto.Unmarshal(msg.Data, v.(proto.Message)); err != nil {
			return Errorf(Internal, "nrpc: error unmarshalling request: %v", err)
		}
		return nil
	}
	reply, err := md.Handler(srv.serviceImpl, ctx, df)
	if err != nil {
		appStatus, ok := FromError(err)
		if !ok {
			appStatus = FromContextError(err)
			err = appStatus.Err()
		}
		s.respondError(msg.Reply, appStatus.Code, appStatus.Message)
		return err
	}
	data, err := proto.Marshal(reply.(proto.Message))
	if err != nil {
		errDesc := fmt.Sprintf("nrpc: error marshaling reply: %v", err)
		s.respondError(msg.Reply, Internal, errDesc)
		return err
	}
	if err := msg.Respond(data); err != nil {
		slog.Warn("nrpc: Server.processMsg failed to respond", "error", err)
		return err
	}
	return nil
}

func (s *Server) respondError(reply string, code Code, message string) error {
	m := nats.NewMsg(reply)
	m.Header.Set(statusHdr, strconv.Itoa(int(code)))
	m.Header.Set(messageHdr, message)

	if err := s.nc.PublishMsg(m); err != nil {
		slog.Warn("nrpc: Server.respondError failed to publish msg", "error", err)
		return err
	}
	return nil
}
