package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//数据类型体
type PostInfo struct {
	Url       string   `json:"url" binding:"required"` //binding:"required" 必须绑定，就是请求时候，必须带上该参数，
	Imgs      []string `json:"imgs"`
	TimeStamp string   `json:"time_stamp" binding:"required"`
	Title     string   `json:"title" binding:"required"`
}

func Index(ctx *gin.Context) {
	var p PostInfo
	//将结构体与json绑定
	err := ctx.BindJSON(&p)
	//数据json绑定是否成功
	if err == nil {
		ctx.JSON(http.StatusOK,
			gin.H{
				"code": "200",
				"msg":  "success",
				"data": p,
			})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": "40001",
			"msg":  "文章数据有误！",
			"data": "",
		})
	}
}
