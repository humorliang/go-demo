package e

var msg = map[int]string{
	//http code
	SUCCESS:             "success",
	INVALID_PARAMS:      "请求参数错误",
	USER_NOT_AUTH:       "用户未认证",
	USER_NOT_PERMISSION: "用户无权限",
	NOT_FOUND:           "请求资源不存在",
	INTERSERVER_ERROR:   "服务器错误",

	// api code
}
//获取错误消息
func GetMsg(code int) string {
	m, ok := msg[code]
	if ok {
		return m
	}
	return msg[500]
}
