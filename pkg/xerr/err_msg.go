package xerr

var errMsgMap = map[int]string{
	SERVER_COMMON_ERROR: "服务器异常,请稍后再试",
	REQUEST_PARAM_ERROR: "请求参数异常",
	DB_ERROR:            "数据库异常",
}

func ErrMsg(code int) string {
	if msg, ok := errMsgMap[code]; ok {
		return msg
	}
	return errMsgMap[SERVER_COMMON_ERROR]
}
