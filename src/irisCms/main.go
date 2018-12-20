/*
addr:127.0.0.1:8080

*/
package main

import (
	"github.com/kataras/iris"
	"flag"
	"irisCms/routers"
)

func main() {
	var app = iris.New()

	//中间件注册

	//路由注册
	routers.New(app)


	//启动配置项
	//命令行参数配置
	var configName = flag.String("configName", "dev", "this is setting application run config")
	//解析cmd
	flag.Parse()
	if *configName == "dev" {
		app.Run(
			iris.Addr("127.0.0.1:8080"),
			iris.WithConfiguration(iris.YAML("./configs/pro.yml")),
		)
	} else if *configName == "pro" {
		app.Run(
			iris.Addr("127.0.0.1:8080"),
			iris.WithConfiguration(iris.YAML("./configs/pro.yml")),
		)
	}
}
