package routers

import (
	"github.com/gin-gonic/gin"
	"ginCms/controllers/user"
	"ginCms/controllers/test"
	"ginCms/middleware"
)

//初始化路由映射函数
func InitRouter(router *gin.Engine) {
	//用户
	//登录
	router.POST("/login", user.Login)
	//注册
	router.POST("/register", user.Register)

	//认证路由
	authRouterGroup := router.Group("/auth", middleware.JWTAuth())
	//测试认证路由
	authRouterGroup.GET("/test", test.ReTest)

}
