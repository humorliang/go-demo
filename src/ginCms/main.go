package main

import (
	"github.com/gin-gonic/gin"
	"flag"
	"ginCms/comm/setting"
	"ginCms/db"
	"ginCms/routers"
	"ginCms/middleware"
	"strconv"
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

	//基础路由对象(不包含任何中间件的路由)
	router := gin.New()

	//中间件注册
	//logger中间件
	router.Use(middleware.Logger())
	//router.Use(middleware.UserId()
	//解决跨域问题中间件
	router.Use(middleware.Cors())
	//异常恢复中间件
	router.Use(gin.Recovery())

	//路由注册,要在注册组件之后
	routers.SetupRouter(router)

	//路由初始化
	router.Run(":" + strconv.Itoa(setting.ServerSetting.HttpPort))
}
