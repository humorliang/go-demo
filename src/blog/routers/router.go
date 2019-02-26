package routers

import (
	"github.com/gin-gonic/gin"
	"blog/controller/user"
)

func Register(router *gin.Engine) {
	v2 := router.Group("/v2")
	v2.GET("/user", user.Login)
}
