package errn

import (
	"encoding/json"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"

	"google.golang.org/protobuf/proto"
)

const UnknownCode = 500

//go:generate protoc --go_out=paths=source_relative:. errors.proto

// Error implements the error interface.
func (e Error) Error() string {
	return e.Message
}

// New returns an error object for the code, message.
func New(code int, message string) *Error {
	return &Error{
		Message: message,
		Errs:    int32(code),
	}
}

// JSON returns data in JSON format
func (e *Error) JSON() ([]byte, error) {
	return json.Marshal(e)
}

// Proto returns data in proto format
func (e *Error) Proto() ([]byte, error) {
	return proto.Marshal(e)
}

// FromError try to convert an error to *Error.
// It supports wrapped errors.
func FromError(err error) *Error {
	if err == nil {
		return nil
	}
	gs, ok := status.FromError(err)
	if ok {
		return &Error{
			Errs:    ErrsFromGRPCCode(gs.Code()), //暂时写死测试，马上要修改
			Message: gs.Message(),
		}
	}
	return &Error{
		Errs:    int32(UnknownCode),
		Message: err.Error(),
	}
}

// StatusCode return an HTTP error code.
func (e *Error) StatusCode() int {
	return int(e.Errs)
}

// GRPCStatus returns the Status represented by se.
func (e *Error) GRPCStatus() *status.Status {
	s, _ := status.New(GRPCCodeFromeErrs(e.Errs), e.Message). //这里其实不能写死
									WithDetails(&errdetails.ErrorInfo{
			Reason:   "",
			Metadata: nil,
		})
	return s
}
