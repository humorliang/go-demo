package user

import (
	"github.com/gin-gonic/gin"
	"fmt"

	"github.com/humorliang/go-demo/blog/controller/data"
	"github.com/humorliang/go-demo/blog/dbops"
	"github.com/humorliang/go-demo/blog/comm/e"
)

type Id struct {
	UserId int `json:"user_id"`
}

//登录
func Login(c *gin.Context) {
	u := data.Admin{
		Username: "user2",
		Password: "123456",
		Realname: "张三",
	}
	err := dbops.AddUser(u)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(e.SUCCESS, "")
}

//注册
func Register(c *gin.Context) {

}

//删除
func Del(c *gin.Context) {

}
