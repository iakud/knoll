package nrpc

import (
	"context"
	"errors"
	"fmt"
	"strconv"
)

type Code int32

const (
	OK Code = 0

	Canceled           Code = 1
	Unknown            Code = 2
	InvalidArgument    Code = 3
	DeadlineExceeded   Code = 4
	NotFound           Code = 5
	AlreadyExists      Code = 6
	PermissionDenied   Code = 7
	ResourceExhausted  Code = 8
	FailedPrecondition Code = 9
	Aborted            Code = 10
	OutOfRange         Code = 11
	Unimplemented      Code = 12
	Internal           Code = 13
	Unavailable        Code = 14
	DataLoss           Code = 15
	Unauthenticated    Code = 16

	_maxCode = 17
)

func (c Code) String() string {
	switch c {
	case OK:
		return "OK"
	case Canceled:
		return "Canceled"
	case Unknown:
		return "Unknown"
	case InvalidArgument:
		return "InvalidArgument"
	case DeadlineExceeded:
		return "DeadlineExceeded"
	case NotFound:
		return "NotFound"
	case AlreadyExists:
		return "AlreadyExists"
	case PermissionDenied:
		return "PermissionDenied"
	case ResourceExhausted:
		return "ResourceExhausted"
	case FailedPrecondition:
		return "FailedPrecondition"
	case Aborted:
		return "Aborted"
	case OutOfRange:
		return "OutOfRange"
	case Unimplemented:
		return "Unimplemented"
	case Internal:
		return "Internal"
	case Unavailable:
		return "Unavailable"
	case DataLoss:
		return "DataLoss"
	case Unauthenticated:
		return "Unauthenticated"
	default:
		return "Code(" + strconv.FormatInt(int64(c), 10) + ")"
	}
}

type Status struct {
	Code    Code
	Message string
}

func (s *Status) String() string {
	return fmt.Sprintf("rpc error: code = %s desc = %s", s.Code, s.Message)
}

func (s *Status) Err() error {
	if s.Code == OK {
		return nil
	}
	return &Error{s: s}
}

func New(c Code, msg string) *Status {
	return &Status{Code: c, Message: msg}
}

func FromContextError(err error) *Status {
	if err == nil {
		return nil
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return New(DeadlineExceeded, err.Error())
	}
	if errors.Is(err, context.Canceled) {
		return New(Canceled, err.Error())
	}
	return New(Unknown, err.Error())
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

func Errorf(c Code, format string, a ...any) error {
	return New(c, fmt.Sprintf(format, a...)).Err()
}

func FromError(err error) (s *Status, ok bool) {
	if err == nil {
		return nil, true
	}
	type nrpcstatus interface{ NRPCStatus() *Status }
	if gs, ok := err.(nrpcstatus); ok {
		nrpcStatus := gs.NRPCStatus()
		if nrpcStatus == nil {
			return New(Unknown, err.Error()), false
		}
		return nrpcStatus, true
	}
	var gs nrpcstatus
	if errors.As(err, &gs) {
		nrpcStatus := gs.NRPCStatus()
		if nrpcStatus == nil {
			return New(Unknown, err.Error()), false
		}
	}
	return New(Unknown, err.Error()), false
}
