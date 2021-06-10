package errn

import (
	"google.golang.org/grpc/codes"
)

// Msgs:错误信息
//var Msgs = map[int]string{
//	0:     "succ",
//	10001: "系统异常",
//	10002: "json解析异常",
//	10004: "参数错误",
//	10009: "系统异常",
//	10100: "DB异常",
//	10011: "无效数据",
//	10012: "参数错误",
//}

var (
	StatusOK           = &Error{Errs: 200, Message: "请求正常"}
	BadRequest         = &Error{Errs: 400, Message: "请求出错"}
	Unauthorized       = &Error{Errs: 401, Message: "未授权错误"}
	Forbidden          = &Error{Errs: 403, Message: "禁止访问错误"}
	NotFound           = &Error{Errs: 404, Message: "访问链接不存在错误"}
	InternalServer     = &Error{Errs: 500, Message: "服务器异常错误"}
	ServiceUnavailable = &Error{Errs: 503, Message: "服务不可用错误"}
)

// GRPCCodeFromStatus converts a HTTP error code into the corresponding gRPC response status.
// See: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
func GRPCCodeFromeErrs(code int32) codes.Code {
	switch code {
	case StatusOK.Errs:
		return codes.OK
	case BadRequest.Errs:
		return codes.InvalidArgument
	case Unauthorized.Errs:
		return codes.Unauthenticated
	case Forbidden.Errs:
		return codes.PermissionDenied
	case NotFound.Errs:
		return codes.NotFound
	case InternalServer.Errs:
		return codes.Internal
	case ServiceUnavailable.Errs:
		return codes.Unavailable
	}
	return codes.Unknown
}

// StatusFromGRPCCode converts a gRPC error code into the corresponding HTTP response status.
// See: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
func ErrsFromGRPCCode(code codes.Code) int32 {
	switch code {
	case codes.OK:
		return StatusOK.Errs
	case codes.InvalidArgument:
		return BadRequest.Errs
	case codes.Unauthenticated:
		return Unauthorized.Errs
	case codes.PermissionDenied:
		return Forbidden.Errs
	case codes.NotFound:
		return NotFound.Errs
	case codes.Internal:
		return InternalServer.Errs
	case codes.Unavailable:
		return ServiceUnavailable.Errs
	}
	return UnknownCode
}
