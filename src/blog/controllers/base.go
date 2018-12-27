package controllers

import (
	"github.com/astaxie/beego"
)


//路由控制
type MainController struct {
	beego.Controller
}

//API数据类型


//get 请求
func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
}
