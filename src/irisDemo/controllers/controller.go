package controllers

import (

	"github.com/kataras/iris"
	"fmt"
)

func InitRouter(app *iris.Application) {
	fmt.Println("init router ......")
	//路由注册
	app.Handle("get","/", func(ctx iris.Context) {
		ctx.HTML("<h1>index</h1>")
		//ctx.WriteString("test")
	})
}
