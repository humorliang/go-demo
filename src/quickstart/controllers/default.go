package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	var jsonData = map[string]string{"key1": "values", "key2": "values"}
	c.Data["json"] = jsonData
	c.ServeJSON()
}
