package status

import (
	"context"
	"errors"
	"fmt"

	"github.com/iakud/knoll/nrpc/codes"
)

type Status struct {
	Code    codes.Code
	Message string
}

func (s *Status) String() string {
	return fmt.Sprintf("rpc error: code = %s desc = %s", s.Code, s.Message)
}

func (s *Status) Err() error {
	if s.Code == codes.OK {
		return nil
	}
	return &Error{s: s}
}

func New(code codes.Code, msg string) *Status {
	return &Status{Code: code, Message: msg}
}

func FromContextError(err error) *Status {
	if err == nil {
		return nil
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return New(codes.DeadlineExceeded, err.Error())
	}
	if errors.Is(err, context.Canceled) {
		return New(codes.Canceled, err.Error())
	}
	return New(codes.Unknown, err.Error())
}

type Error struct {
	s *Status
}

func (e *Error) Error() string {
	return e.s.String()
}

func (e *Error) NRPCStatus() *Status {
	return e.s
}

func (e *Error) Is(target error) bool {
	if _, ok := target.(*Error); ok {
		return true
	}
	return false
}

func Errorf(code codes.Code, format string, a ...any) error {
	return New(code, fmt.Sprintf(format, a...)).Err()
}

func FromError(err error) (s *Status, ok bool) {
	if err == nil {
		return nil, true
	}
	type nrpcstatus interface{ NRPCStatus() *Status }
	if gs, ok := err.(nrpcstatus); ok {
		nrpcStatus := gs.NRPCStatus()
		if nrpcStatus == nil {
			return New(codes.Unknown, err.Error()), false
		}
		return nrpcStatus, true
	}
	var gs nrpcstatus
	if errors.As(err, &gs) {
		nrpcStatus := gs.NRPCStatus()
		if nrpcStatus == nil {
			return New(codes.Unknown, err.Error()), false
		}
	}
	return New(codes.Unknown, err.Error()), false
}
