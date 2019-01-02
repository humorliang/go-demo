package tags

import (
	"github.com/gin-gonic/gin"
	"ginCms/controllers"
)

//获取分类
func GetTags(c *gin.Context) {
	//ctx := controllers.Context{c}
	ctx:=controllers.Context{c}
	ctx.Success("成功")
}

//增加分类
func AddTag(c *gin.Context) {
	ctx:=controllers.Context{c}
	ctx.Success("成功")
}

//更新分类
func UpdateTag(c *gin.Context) {
	ctx:=controllers.Context{c}
	ctx.Success("成功")
}

//删除分类
func DeleteTag(c *gin.Context) {
	ctx:=controllers.Context{c}
	ctx.Success("成功")
}
