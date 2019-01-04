package test

import (
	"github.com/gin-gonic/gin"
)

func ReTest(c *gin.Context) {
	//ctx := controllers.Context{c}
	//ctx.Success("要认证的路由！")
	//ctx.Success(setting.ServerSetting)
}

func ReTestLogger(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "test",
	})
	//c.JSON(200, gin.H{
	//	"data": "test2",
	//})
}
