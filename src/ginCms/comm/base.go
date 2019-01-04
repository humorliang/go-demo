package comm

import "ginCms/controllers"

//检查是否是错误
func IsCheckError(ctx *controllers.Context, f interface{}, err error) (c bool) {
	if err != nil {
		return true
	} else {
		return false
	}
}
