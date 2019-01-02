package main

import (
	"github.com/gin-gonic/gin"
	"ginCms/middleware"
	"flag"
	"ginCms/comm/setting"
	"ginCms/routers"
	"ginCms/db"
)

func main() {
	var mode = flag.String("mode", "dev", "this is run mode options")
	//命令行解析
	flag.Parse()

	//设置运行模式
	//命令行需要解指针,读取运行模式
	if *mode == "pro" {
		setting.SetUp("pro")
	} else {
		setting.SetUp("dev")
	}
	//初始化数据库
	db.Setup()

	//设置模式
	gin.SetMode(setting.ServerSetting.RunMode)

	//基础路由(不包含任何中间件的路由)
	router := routers.SetupRouter()

	//中间件注册
	//logger中间件
	router.Use(middleware.Logger())
	//异常恢复中间件
	router.Use(gin.Recovery())

	//路由初始化
	router.Run()

}
