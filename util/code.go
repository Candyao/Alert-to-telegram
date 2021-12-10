package util

var Msg = map[int]string{
	SUCCESS:              "ok",
	ERROR:                "fail",
	INVALID_PARAMS:       "请求参数错误",
	ERROR_ADD_AlERT_FAIL: "新增告警失败",
	ERROR_SEND_MSG:       "发送告警失败",
	ERROR_BIND_JSON:      "获取告警信息失败",
}

func GetMsg(code int) string {
	msg, ok := Msg[code]
	if ok {
		return msg
	}
	return Msg[ERROR]
}
