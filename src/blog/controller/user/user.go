package user

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
	c.JSON(200,"login user")
}
