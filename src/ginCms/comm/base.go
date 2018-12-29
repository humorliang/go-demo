package comm

//检查是否是错误
func IsCheckError(f interface{}, err error) (c bool) {
	if err != nil {
		return true
	} else {
		return false
	}
}
