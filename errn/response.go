package errn

//JSONResult json数据
type JSONResult struct {
	Error int         `json:"error"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}

// Convert convert err to Error
func ConvertJson(err error, msg string, data interface{}) JSONResult {
	error := FromError(err)
	if msg == "" {
		return JSONResult{
			Error: int(error.GetErrs()),
			Msg:   error.GetMessage(),
			Data:  data,
		}
	}
	return JSONResult{
		Error: int(error.GetErrs()),
		Msg:   msg,
		Data:  data,
	}
}

func ErrorConvert(err error) *Error {
	return FromError(err)
}
