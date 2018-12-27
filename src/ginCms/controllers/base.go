package controllers

import (
	"github.com/gin-gonic/gin"
	"ginCms/comm"
)

func Index(ctx *gin.Context) {
	comm.Log("debug").Println("test log")
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": "",
	})

}
