package nrpc

import (
	"context"
	"fmt"
	"log/slog"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/iakud/knoll/nrpc/codes"
	"github.com/iakud/knoll/nrpc/nrpcutil"
	"github.com/iakud/knoll/nrpc/status"
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
	sub, err := s.nc.QueueSubscribe(s.subj, "", s.handleMsg)
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
	var ctx context.Context
	var cancel context.CancelFunc
	if value := msg.Header.Get(timeoutHdr); len(value) > 0 {
		var timeout time.Duration
		var err error
		if timeout, err = nrpcutil.DecodeDuration(value); err != nil {
			s.respondStatus(msg, codes.Internal, fmt.Sprintf("malformed nrpc-timeout: %v", err))
			return
		}
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}
	defer cancel()

	sm := msg.Header.Get(methodHdr)
	if sm != "" && sm[0] == '/' {
		sm = sm[1:]
	}
	pos := strings.LastIndex(sm, "/")
	if pos == -1 {
		s.respondStatus(msg, codes.Unimplemented, fmt.Sprintf("malformed nrpc-method: %q", sm))
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
			return status.Errorf(codes.Internal, "nrpc: error unmarshalling request: %v", err.Error())
		}
		return nil
	}
	reply, err := md.Handler(srv.serviceImpl, ctx, df)
	if err != nil {
		appStatus, ok := status.FromError(err)
		if !ok {
			appStatus = status.FromContextError(err)
			err = appStatus.Err()
		}
		s.respondStatus(msg, appStatus.Code, appStatus.Message)
		return err
	}
	data, err := proto.Marshal(reply.(proto.Message))
	if err != nil {
		s.respondStatus(msg, codes.Internal, fmt.Sprintf("nrpc: error marshaling reply: %v", err))
		return err
	}
	return s.respond(msg, data)
}

func (s *Server) respondStatus(msg *nats.Msg, code codes.Code, message string) error {
	replyMsg := nats.NewMsg("")
	replyMsg.Header.Set(statusHdr, strconv.Itoa(int(code)))
	replyMsg.Header.Set(messageHdr, message)
	if err := msg.RespondMsg(replyMsg); err != nil {
		slog.Warn("nrpc: Server.respondStatus failed to respond msg", "error", err.Error())
		return err
	}
	return nil
}

func (s *Server) respond(msg *nats.Msg, data []byte) error {
	if err := msg.Respond(data); err != nil {
		slog.Warn("nrpc: Server.respond failed to respond", "error", err.Error())
		return err
	}
	return nil
}
