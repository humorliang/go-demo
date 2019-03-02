package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/humorliang/go-demo/ginCms/controllers"
)

func UserId() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := controllers.Context{c}
		ctx.Set("k", 1)
		ctx.Next()
	}
}
