package tag

import (
	"github.com/gin-gonic/gin"
	"ginCms/controllers"
	"ginCms/db"
	"log"
)

type Tag struct {
	Id      int64  `json:"id"`
	TagName string `json:"tag_name"`
}
type AddTagJson struct {
	Data struct {
		TagName string `json:"tag_name"`
	} `json:"data"`
}
type UpTagJson struct {
	Data struct {
		ID      int    `json:"id"`
		TagName string `json:"tag_name"`
	} `json:"data"`
}
type DelTagJson struct {
	Data struct {
		ID int `json:"id"`
	} `json:"data"`
}

//获取分类
func GetTags(c *gin.Context) {
	ctx := controllers.Context{c}
	var tag Tag
	var tags []Tag
	//数据查询
	rows, err := db.Con.Query("SELECT id,tag_name FROM tag")
	if err != nil {
		ctx.Fail(500, "10001", "获取分类失败！")
	} else {
		//读取记录集合
		for rows.Next() {
			err := rows.Scan(&tag.Id, &tag.TagName)
			if err != nil {
				ctx.Fail(500, "10001", "获取分类失败！")
			}
			tags = append(tags, tag)
		}
		ctx.Success(gin.H{
			"tag_list": tags,
		})
	}
}

//增加分类
func AddTag(c *gin.Context) {
	ctx := controllers.Context{c}
	//json数据绑定
	tag := &AddTagJson{}
	err := ctx.BindJSON(tag)
	if err != nil {
		log.Println(err)
		ctx.Fail(500, "10001", "添加分类失败！")
	}
	//执行插入语句
	res, err := db.Con.Exec("INSERT INTO tag (tag_name) VALUES (?)", tag.Data.TagName)
	if err != nil {
		log.Println(err)
		ctx.Fail(500, "10001", "添加分类失败！")
	} else {
		//获取插入tagId
		tagId, err := res.LastInsertId()
		if err != nil {
			log.Println(err)
			ctx.Fail(500, "10001", "获取分类Id失败")
		} else {
			ctx.Success(gin.H{
				"tag_id":   tagId,
				"tag_name": tag.Data.TagName,
			})
		}
	}

}

//更新分类
func UpdateTag(c *gin.Context) {
	ctx := controllers.Context{c}
	//绑定json数据
	upTag := &UpTagJson{}
	err := ctx.BindJSON(upTag)
	if err != nil {
		ctx.Fail(500, "10001", "更新分类失败！")
	}
	//更新数据SQL
	_, err = db.Con.Exec("UPDATE tag SET tag_name=? WHERE id=?", upTag.Data.TagName, upTag.Data.ID)
	if err != nil {
		ctx.Fail(500, "10001", "更新分类失败！")
	} else {
		if err != nil {
			ctx.Fail(500, "10001", "更新分类失败！")
		} else {
			ctx.Success("更新成功！")
		}
	}
}

//删除分类
func DeleteTag(c *gin.Context) {
	ctx := controllers.Context{c}
	//json数据绑定
	delTag := &DelTagJson{}
	err := ctx.BindJSON(delTag)
	if err != nil {
		ctx.Fail(500,"10001","删除标签失败！")
	}
	//执行删除SQL
	_,err:=db.Con.Exec("DELETE FROM tag WHERE id=?",delTag.Data.ID)
	if err!=nil {
		ctx.Fail("")
	}
}
