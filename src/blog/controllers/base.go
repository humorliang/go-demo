package controllers

import (
	"github.com/astaxie/beego"
)

//路由控制
type BaseController struct {
	beego.Controller //beego基础
	IsLogin    bool  //是否登录
	UserUserId int64 //用户id
}

// 基础控制准备方法
func (c *BaseController) Prepare() {
	userLogin := c.GetSession("userLogin")
	//判断用户是否登录
	if userLogin == nil {
		c.IsLogin = false
	} else {
		c.IsLogin = true
	}
}
