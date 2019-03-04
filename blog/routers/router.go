package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/humorliang/go-demo/blog/controller/admin/user"
)

func Register(router *gin.Engine) {
	v2 := router.Group("/v2")
	v2.GET("/user", user.Login)
}
