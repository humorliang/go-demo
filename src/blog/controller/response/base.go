package response

func SuccessResponse(data interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	m["code"] = 0
	m["msg"] = "SUCCESS"
	m["data"] = data
	return m
}

func FailResponse(code int, data interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	m["code"] = code
	m["msg"] = "FAILS"
	m["data"] = data
	return m
}
