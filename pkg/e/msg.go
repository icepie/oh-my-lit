package e

// MsgFlags 状态码map
var MsgFlags = map[int]string{
	Success:         "ok",
	Error:           "fail",
	InvalidParams:   "invalid params",
	ErrorGenToken:   "fali to generate token",
	ErrorJWTCheck:   "token error",
	ErrorJWTTimeout: "token timeout",
}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[Error]
}
