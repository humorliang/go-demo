package user

import (
	"github.com/gin-gonic/gin"
	"ginCms/comm/e"
	"blog/controller/response"
)

type Id struct {
	UserId int `json:"user_id"`
}

//登录
func Login(c *gin.Context) {
	//c.JSON(e.SUCCESS,response.SuccessResponse("login user"))
	c.JSON(e.ERROR, response.FailResponse(1000, "login user"))
}

//注册
func Register(c *gin.Context) {
}

//删除
func Del(c *gin.Context) {

}
