package article

import (
	"github.com/gin-gonic/gin"
	"ginCms/controllers"
)

//获取文章
func GetTagArticle(c *gin.Context) {
	ctx:=controllers.Context{c}
	ctx.Success("成功")
}

//获取推荐文章
func GetRecomArticle(c *gin.Context) {
	ctx:=controllers.Context{c}
	ctx.Success("成功")
}

//获取文章详情
func GetDescArticle(c *gin.Context) {
	ctx:=controllers.Context{c}
	ctx.Success("成功")
}

//添加文章
func AddArticle(c *gin.Context) {
	ctx:=controllers.Context{c}
	ctx.Success("成功")
}

//删除文章
func DeleteArticle(c *gin.Context) {
	ctx:=controllers.Context{c}
	ctx.Success("成功")
}
