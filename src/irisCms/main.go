/*
addr:127.0.0.1:8080

*/
package main

import (
	"github.com/kataras/iris"
	"irisCms/routers"
	"flag"
)

func main() {
	app := iris.New()

	//中间件注册

	//路由注册
	routers.New(app)
	//初始化命令行
	//命令行参数配置
	var AppEnvName = flag.String("appEnvName", "dev", "this is set application run env")
	flag.Parse()
	if *AppEnvName == "dev" {
		app.Run(
			iris.Addr("127.0.0.1:8080"),
			iris.WithConfiguration(iris.YAML("./configs/dev.yml")),
		)
		//fmt.Println(app.ConfigurationReadOnly().GetFireMethodNotAllowed())
	} else if *AppEnvName == "pro" {
		app.Run(
			iris.Addr("127.0.0.1:8080"),
			iris.WithConfiguration(iris.YAML("./configs/pro.yml")),
		)
	}
}
