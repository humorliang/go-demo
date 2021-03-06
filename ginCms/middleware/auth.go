package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/dgrijalva/jwt-go"

	"github.com/humorliang/go-demo/ginCms/utils"
	"github.com/humorliang/go-demo/ginCms/comm"
	"github.com/humorliang/go-demo/ginCms/controllers"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := controllers.Context{c}
		//进行token验证
		//获取token
		token := ctx.Request.Header.Get("authorization")
		if token == "" {
			ctx.Fail(500, "40001", "请求未认证")
		} else {
			//判断token
			claims, err := utils.ParseToken(token)
			if err != nil {
				comm.Log("error").Println(err)
				//token错误判断
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ctx.Fail(400, "40001", "Token已过期！")
					ctx.Abort()
					return
				default:
					ctx.Fail(400, "40002", "Token认证失败！")
					ctx.Abort()
					return
				}
			}
			ctx.Set("userId", claims.UserId)
		}
		//
		ctx.Next()
	}
}
