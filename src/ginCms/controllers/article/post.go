package article

import (
	"github.com/gin-gonic/gin"
	"ginCms/controllers"
	"ginCms/db"
	"ginCms/comm"
	"ginCms/comm/setting"
)

type AddArticleJson struct {
	Data struct {
		Content       string `json:"content" binding:"required"`
		Desp          string `json:"desp" binding:"required"`
		TagID         int    `json:"tagId" binding:"required"`
		Title         string `json:"title" binding:"required"`
		UserID        int    `json:"userId" binding:"required"`
		PreviewImgUrl string `json:"previewImgUrl"`
		IsRecommend   int    `json:"isRecommend"`
	} `json:"data"`
}
type GetArticleListJson struct {
	Data struct {
		PageNum int `json:"pageNum" binding:"required"`
		TagID   int `json:"tagId" binding:"required"`
	} `json:"data"`
}
type ArticleItemInfo struct {
	PostId        int     `json:"postId"`
	Desp          string  `json:"desp"`
	Title         string  `json:"title"`
	Content       string  `json:"content,omitempy"`
	Date          string  `json:"date"`
	PreviewImgUrl *string `json:"previewImgUrl"` //当数据库解析字段为空值时，变为nil
	UserId        int     `json:"userId"`
	PenName       string  `json:"penName"`
	TagId         int     `json:"tagId"`
	TagName       string  `json:"tagName"`
}
type ArticleId struct {
	Data struct {
		PostID int `json:"postId" binding:"required"`
	} `json:"data"`
}

//获取文章
func GetTagArticle(c *gin.Context) {
	ctx := controllers.Context{c}
	//json数据结构化
	article := &GetArticleListJson{}
	err := ctx.BindJSON(&article)
	if err != nil {
		comm.Log("error").Println(err)
		ctx.Fail(500, "10001", "获取文章失败")
		return
	}
	//Sql查询
	pgSize := setting.AppSetting.PageSize
	pgNumSize := (article.Data.PageNum - 1) * pgSize
	rows, err := db.Con.Query("SELECT a.id,a.desp,a.title,"+
		"a.date,a.preview_img_url,a.user_id,a.tag_id,u.pen_name,t.tag_name "+
		"FROM user u,article a,tag t "+
		"WHERE a.user_id=u.id AND a.tag_id=t.id AND a.tag_id=? "+
		"LIMIT ?,?",
		article.Data.TagID, pgNumSize, pgSize)
	defer rows.Close()
	if err != nil {
		comm.Log("error").Println(err)
		ctx.Fail(500, "10001", "获取推荐文章失败")
		return
	}

	//sql结果解析
	var art ArticleItemInfo
	var artList []ArticleItemInfo
	for rows.Next() {
		err := rows.Scan(&art.PostId, &art.Desp, &art.Title, &art.Date,
			&art.PreviewImgUrl, &art.UserId, &art.TagId, &art.PenName, &art.TagName)
		if err != nil {
			comm.Log("error").Println(err)
			ctx.Fail(500, "10001", "获取推荐文章失败！")
			return
		} else {
			artList = append(artList, art)
		}
	}
	ctx.Success(artList)
}

//获取推荐文章
func GetRecomArticle(c *gin.Context) {
	ctx := controllers.Context{c}
	//sql查询
	rows, err := db.Con.Query("SELECT a.id,a.desp,a.title," +
		"a.date,a.preview_img_url,a.user_id,a.tag_id,u.pen_name,t.tag_name " +
		"FROM user u,article a,tag t " +
		"WHERE a.user_id=u.id AND a.tag_id=t.id AND a.is_recommend=1 ORDER BY a.date DESC LIMIT 6")
	defer rows.Close()
	if err != nil {
		comm.Log("error").Println(err)
		ctx.Fail(500, "10001", "获取推荐文章失败")
		return
	}

	//sql结果解析
	var art ArticleItemInfo
	var artList []ArticleItemInfo
	for rows.Next() {
		err := rows.Scan(&art.PostId, &art.Desp, &art.Title, &art.Date,
			&art.PreviewImgUrl, &art.UserId, &art.TagId, &art.PenName, &art.TagName)
		if err != nil {
			comm.Log("error").Println(err)
			ctx.Fail(500, "10001", "获取推荐文章失败！")
			return
		} else {
			artList = append(artList, art)
		}
	}
	ctx.Success(artList)
}

//获取文章详情
func GetDescArticle(c *gin.Context) {
	ctx := controllers.Context{c}
	//json结构化
	artId := &ArticleId{}
	err := ctx.BindJSON(artId)
	if err != nil {
		ctx.Fail(500, "10001", "获取文章详情失败！")
		return
	}
	//sql
	rows, err := db.Con.Query("SELECT a.id,a.desp,a.title,a.content,"+
		"a.date,a.preview_img_url,a.user_id,a.tag_id,u.pen_name,t.tag_name "+
		"FROM user u,article a,tag t "+
		"WHERE a.user_id=u.id AND a.tag_id=t.id AND a.id=?", artId.Data.PostID)
	if err != nil {
		comm.Log("error").Println(err)
		ctx.Fail(500, "10001", "获取文章详情失败！")
		return
	}
	//fmt.Println(rows, artDesp.Data.PostID)
	//sql结果解析
	var art ArticleItemInfo
	var artList []ArticleItemInfo
	for rows.Next() {
		err := rows.Scan(&art.PostId, &art.Desp, &art.Title, &art.Content, &art.Date,
			&art.PreviewImgUrl, &art.UserId, &art.TagId, &art.PenName, &art.TagName)
		if err != nil {
			comm.Log("error").Println(err)
			ctx.Fail(500, "10001", "获取推荐文章失败！")
			return
		} else {
			artList = append(artList, art)
		}
	}
	if len(artList) != 0 {
		ctx.Success(artList[0])
	} else {
		ctx.Fail(500, "10001", "没有文章详情")
	}
}

//添加文章
func AddArticle(c *gin.Context) {
	ctx := controllers.Context{c}
	//绑定文章结构体
	article := &AddArticleJson{}
	err := ctx.BindJSON(article)
	if err != nil {
		comm.Log("error").Println(err)
		ctx.Fail(500, "10001", "添加文章失败")
		return
	}
	//SQL
	res, err := db.Con.Exec("INSERT INTO article (title,desp,content,user_id,tag_id,is_recommend) VALUES (?,?,?,?,?,?)",
		article.Data.Title,
		article.Data.Desp,
		article.Data.Content,
		article.Data.UserID,
		article.Data.TagID,
		article.Data.IsRecommend)
	if err != nil {
		comm.Log("error").Println(err)
		ctx.Fail(500, "10001", "添加文章失败")
	} else {
		pId, err := res.LastInsertId()
		if err != nil {
			comm.Log("error").Println(err)
			ctx.Fail(500, "10001", "添加文章失败")
		} else {
			ctx.Success(gin.H{
				"postId": pId,
				"title":  article.Data.Title,
			})
		}
	}
}

//删除文章
func DeleteArticle(c *gin.Context) {
	ctx := controllers.Context{c}
	//json结构化
	artId := &ArticleId{}
	err := ctx.BindJSON(artId)
	if err != nil {
		comm.Log("error").Println(err)
		ctx.Fail(500, "10001", "删除文章失败！")
		return
	}
	//sql
	_, err = db.Con.Exec("DELETE FROM article WHERE id=?", artId.Data.PostID)
	if err != nil {
		comm.Log("error").Println(err)
		ctx.Fail(500, "10001", "删除文章失败！")
		return
	} else {
		ctx.Success("文章删除成功！")
	}
}
