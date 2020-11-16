package e

var Msg = map[int]string{
	Success:     "成功",
	Fail:        "失败",
	ErrorCode:   "无效错误码",
	ParamsErr:   "参数错误",
	InvalidUser: "无效用户",
	ServerErr:   "服务器错误",
}

func GetMsg(code int) string {
	msg, ok := Msg[code]
	if !ok {
		return Msg[Fail]
	}
	return msg
}
