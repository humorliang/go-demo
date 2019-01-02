package tags

import (
	"github.com/gin-gonic/gin"
	"ginCms/controllers"
	"ginCms/db"
)

type Tag struct {
	TagName string `json:"tag_name"`
}

//获取分类
func GetTags(c *gin.Context) {
	ctx := controllers.Context{c}

	ctx.Success("成功")
}

//增加分类
func AddTag(c *gin.Context) {
	ctx := controllers.Context{c}
	//json数据绑定
	tag := &Tag{}
	tagName := c.BindJSON(tag)
	//执行插入语句
	res, err := db.Con.Exec("INSERT INTO tag_name VALUES ?", tagName)
	if err != nil {
		ctx.Fail(500, "10001", "添加分类失败！")
	} else {
		//获取插入tagId
		tagId, err := res.LastInsertId()
		if err != nil {
			ctx.Fail(500, "10001", "获取分类Id失败")
		} else {
			ctx.Success(gin.H{
				"tag_id":   tagId,
				"tag_name": tagName,
			})
		}
	}

}

//更新分类
func UpdateTag(c *gin.Context) {
	ctx := controllers.Context{c}
	ctx.Success("成功")
}

//删除分类
func DeleteTag(c *gin.Context) {
	ctx := controllers.Context{c}
	ctx.Success("成功")
}
