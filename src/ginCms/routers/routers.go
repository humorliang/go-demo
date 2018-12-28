package routers

import (
	"github.com/gin-gonic/gin"
	"ginCms/controllers/user"
)

//初始化路由映射函数
func InitRouter(router *gin.Engine) {
	//用户
	//登录
	router.POST("/login", user.Login)
	//注册
	router.POST("/register", user.Register)
}
