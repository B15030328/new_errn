package errn

// Msgs:错误信息
var Msgs = map[int]string{
	0:     "succ",
	10001: "系统异常",
	10002: "json解析异常",
	10004: "参数错误",
	10009: "系统异常",
	10100: "DB异常",
	10011: "无效数据",
	10012: "参数错误",
}

var (
	BadRequest         = &Error{Errs: 400, Message: "请求出错"}
	Unauthorized       = &Error{Errs: 401, Message: "未授权错误"}
	Forbidden          = &Error{Errs: 403, Message: "禁止访问错误"}
	NotFound           = &Error{Errs: 404, Message: "访问链接不存在错误"}
	InternalServer     = &Error{Errs: 500, Message: "服务器异常错误"}
	ServiceUnavailable = &Error{Errs: 503, Message: "服务不可用错误"}
	Conflict           = &Error{Errs: 409, Message: "请求冲突错误"}
	System             = &Error{Errs: 10001, Message: "系统异常"}
	JsonFail           = &Error{Errs: 10002, Message: "json解析异常"}
	InvalidParams      = &Error{Errs: 10004, Message: "参数错误"}
	Deprecated         = &Error{Errs: 10009, Message: "系统异常"}
	DB                 = &Error{Errs: 10100, Message: "DB异常"}
	InvalidData        = &Error{Errs: 10011, Message: "无效数据"}
	BindErr            = &Error{Errs: 10012, Message: "参数错误"}
)
