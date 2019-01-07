package test

import (
	"github.com/gin-gonic/gin"
	"ginCms/controllers"
	"fmt"
)

func ReTest(c *gin.Context) {
	ctx := controllers.Context{c}
	if v, ok := ctx.Get("userId"); ok {
		ctx.Success(gin.H{
			"userId": v,
		})
	} else {
		ctx.Fail(400, "10001", "用户未认证！")
	}
	//ctx.Success("要认证的路由！")
	//ctx.Success(setting.ServerSetting)
}

func ReTestLogger(c *gin.Context) {
	ctx := controllers.Context{c}
	fmt.Println(ctx.Get("k"))
	c.JSON(200, gin.H{
		"data": "test",
	})
	//c.JSON(200, gin.H{
	//	"data": "test2",
	//})
}
