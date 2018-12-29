package middleware

import (
	"github.com/gin-gonic/gin"
	"ginCms/controllers"
)

type User struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := controllers.Context{c}
		//

		//鉴权完毕
		ctx.Next()
	}
}

//创建一个Token
func createToken(key string, m map[string]interface{}) string {
	//token:=
}
