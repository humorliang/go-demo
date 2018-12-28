package main

import (
	"github.com/gin-gonic/gin"
	"ginCms/routers"
	"ginCms/middleware"
)

func init() {
	//设置运行环境
	//gin.SetMode("release")

}

func main() {

	//基础路由(不包含任何中间件的路由)
	router := gin.New()

	//中间件注册
	//logger中间件
	router.Use(middleware.Logger())
	//异常恢复中间件
	router.Use(gin.Recovery())

	//路由初始化
	routers.InitRouter(router)

	//设置运行模式
	//gin.SetMode(gin.ReleaseMode)
	router.Run()
}
