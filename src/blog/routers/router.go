package routers

import (
	"github.com/astaxie/beego"
	"blog/controllers/home"
)

func init() {
	beego.Router("/", &home.IndexController{})
}
