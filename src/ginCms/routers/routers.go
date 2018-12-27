package routers

import (
	"github.com/gin-gonic/gin"
	"ginCms/controllers"
)

//初始化路由映射函数
func InitRouter(router *gin.Engine) {
	//传入handleFunc
	router.GET("/test", controllers.Index)
}
