package main

import (
	"irisDemo/app"
	"github.com/kataras/iris"
	"flag"
	"fmt"
	"github.com/kataras/iris/core/errors"
)

func main() {
	//配置命令行
	var configName = flag.String("configName", "dev", "this is setting application run config")
	//解析cmd
	flag.Parse()
	//获取命令行参数
	fmt.Println("configName:", *configName)

	//应用实例化
	var instanceApp = app.InitApp()

	instanceApp.Handle("GET", "/welcome", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
		ctx.Next()
	})

	//环境配置
	if *configName == "dev" {
		instanceApp.Run(
			//服务器地址
			iris.Addr("127.0.0.1:8080"),
			//服务器环境配置文件
			iris.WithConfiguration(iris.YAML("./conf/dev.yml")),
			// 按下CTRL / CMD + C时跳过错误的服务器：
			iris.WithoutServerError(iris.ErrServerClosed),
		)
	} else if *configName == "pro" {
		instanceApp.Run(
			iris.Addr("127.0.0.1:8080"),
			iris.WithConfiguration(iris.YAML("./conf/pro.yml")),
			iris.WithoutServerError(iris.ErrServerClosed),
		)
	} else {
		errors.New("this is config name not have yml file!")
	}

}
