package controllers

import "github.com/gin-gonic/gin"

//数据类型体
//type PostInfo struct {
//	Url       string   `json:"url" binding:"required"` //binding:"required" 必须绑定，就是请求时候，必须带上该参数，
//	Imgs      []string `json:"imgs"`
//	TimeStamp string   `json:"time_stamp" binding:"required"`
//	Title     string   `json:"title" binding:"required"`
//}

//用户类型结构体
type UserInfo struct {
	Id       int64  `json:"userId"`
	Username string `json:"username",form:"username"`
	Penname  string `json:"penname"`
	Password string `json:"password",form:"password"`
	Email    string `json:"email"`
}

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

