package main

import (
	"github.com/gin-gonic/gin"
	"ginCms/routers"
)

func init() {
	//设置运行环境
	//gin.SetMode("release")

}

func main() {

	//基础路由(不包含任何中间件的路由)
	router := gin.New()

	//注册logger中间件
	//router.Use(middleware.RequestLogger())

	//路由初始化
	routers.InitRouter(router)

	router.Run()
}
