/*
addr:127.0.0.1:8080

*/
package main

import (
	"github.com/kataras/iris"
	"irisCms/routers"
	"flag"
	"irisCms/middleware"
	recover2 "github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()

	//中间件注册
	app.Use(middleware.NewRequestLogger())
	//异常恢复
	app.Use(recover2.New())


	//路由注册
	routers.New(app)
	//初始化命令行
	//命令行参数配置
	var AppEnvName = flag.String("appEnvName", "dev", "this is set application run env")
	flag.Parse()
	if *AppEnvName == "dev" {
		app.Run(
			iris.Addr("127.0.0.1:8080"),
			iris.WithConfiguration(iris.YAML("./configs/dev.ini")),
		)
		//fmt.Println(app.ConfigurationReadOnly().GetFireMethodNotAllowed())
	} else if *AppEnvName == "pro" {
		app.Run(
			iris.Addr("127.0.0.1:8080"),
			iris.WithConfiguration(iris.YAML("./configs/pro.ini")),
		)
	}
}
