package controllers

import "github.com/gin-gonic/gin"

//请求内容格式
type Context struct {
	*gin.Context
}


//请求成功响应格式
func (ctx *Context) Success(data interface{}) {
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}

//请求失败响应格式
func (ctx *Context) Fail(status int, code interface{}, msg string) {
	ctx.JSON(status, gin.H{
		"code": code,
		"msg":  msg,
		"data": "",
	})
}
