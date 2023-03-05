package e

var MsgFlags = map[uint]string{
	Success:       "ok",
	Error:         "fail",
	InvalidParams: "请求参数错误",
}

// GetMsg 获取状态码信息
func GetMsg(code uint) string {
	if msg, ok := MsgFlags[code]; ok {
		return msg
	}
	return MsgFlags[Error]
}
