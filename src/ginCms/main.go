package main

import (
	"github.com/gin-gonic/gin"
	"ginCms/routers"
	"ginCms/middleware"
	"flag"
	"ginCms/comm/setting"
	"fmt"
)

func main() {
	var mode = flag.String("mode", "dev", "this is run mode options")
	//命令行解析
	flag.Parse()

	//设置运行模式
	//gin.SetMode(gin.ReleaseMode)

	//命令行需要解指针,读取运行模式
	if *mode == "pro" {
		setting.SetUp("pro")
	} else {
		setting.SetUp("dev")
	}
	//设置模式
	gin.SetMode(setting.ServerSetting.RunMode)

	//基础路由(不包含任何中间件的路由)
	router := gin.New()

	//中间件注册
	//logger中间件
	router.Use(middleware.Logger())
	//异常恢复中间件
	router.Use(gin.Recovery())

	//路由初始化
	routers.InitRouter(router)
	router.Run()

	fmt.Sprintf("run")
}
