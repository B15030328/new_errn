package errn

import (
	"encoding/json"

	"google.golang.org/protobuf/proto"
)

const UnknownCode = 500

//go:generate protoc --go_out=paths=source_relative:. errors.proto

// Error implements the error interface.
func (e Error) Error() string {
	return e.Message
}

//// ResetMsg 重置msg
//func (e *Error) ResetMsg(msg string) *Error {
//	Msgs[int(e.Errs)] = msg
//	e.Message = msg
//	return e
//}

// New returns an error object for the code, message.
func New(code int, message string) *Error {
	//1、判断该code是否已绑定错误信息
	if _, ok := Msgs[code]; ok {
		panic("code has been difined")
		//return nil, errors.New("code has been difined")
	}
	//2、写入Msg并返回Error对象
	Msgs[code] = message
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
	for k, v := range Msgs {
		if v == err.Error() {
			return &Error{
				Errs:    int32(k),
				Message: err.Error(),
			}
		}
	}
	return &Error{
		Errs:    int32(UnknownCode),
		Message: err.Error(),
	}
}

// FromProto returns a proto.
func FromProto(ms []byte) *Error {
	se := new(Error)
	err := proto.Unmarshal(ms, se)
	if err != nil {
		panic("FromProto error")
	}
	return se
}
