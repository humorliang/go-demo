package comment

import (
	"github.com/gin-gonic/gin"
	"ginCms/controllers"
	"ginCms/db"
	"ginCms/comm"
	"ginCms/comm/setting"
)

type Comment struct {
	Data struct {
		Content   string `json:"content" binding:"required"`
		UserId    int    `json:"userId" binding:"required"`
		ArticleId int    `json:"articleId" binding:"required"`
	} `json:"data"`
}

type CommentList struct {
	Data struct {
		PostId  int `json:"postId" binding:"required"`
		PageNum int `json:"pageNum" binding:"required"`
	} `json:"data"`
}

type CommentInfo struct {
	CommentId int    `json:"commentId"`
	Content   string `json:"content"`
	Date      string `json:"date"`
	PenName   string `json:"penName"`
}

//发表评论
func AddComment(c *gin.Context) {
	ctx := controllers.Context{c}

	//Json结构化绑定
	cmt := &Comment{}
	err := ctx.BindJSON(cmt)
	if err != nil {
		comm.Log("error").Println(err)
		ctx.Fail(400, "100001", "发表评论失败！")
		return
	}
	//sql操作
	res, err := db.Con.Exec("INSERT INTO comment (content,user_id,article_id) Values (?,?,?)", cmt.Data.Content, cmt.Data.UserId, cmt.Data.ArticleId)
	if err != nil {
		comm.Log("error").Println(err)
		ctx.Fail(500, "100001", "发表评论失败！")
		return
	}
	//返回发表成功评论Id
	cId, err := res.LastInsertId()
	if err != nil {
		comm.Log("error").Println(err)
		ctx.Fail(500, "100001", "发表评论失败！")
		return
	} else {
		ctx.Success(gin.H{
			"commentId": cId,
		})
	}
}

//获取评论
func GetPostComment(c *gin.Context) {
	ctx := controllers.Context{c}
	//JSON结构化绑定
	cmtList := &CommentList{}
	err := ctx.BindJSON(cmtList)
	if err != nil {
		comm.Log("error").Println(err)
		ctx.Fail(400, "100001", "JSON数据错误！")
		return
	}
	//Sql
	ctPgSize := setting.AppSetting.CommentPageSize
	ctPgNumSize := (cmtList.Data.PageNum - 1) * ctPgSize
	rows, err := db.Con.Query("SELECT c.id,c.content,c.date,u.pen_name"+
		"FROM comment c,user u WHERE u.id=(SELECT user_id From article WHERE article_id=?) " +
		"AND c.article_id=? ORDER BY c.date LIMIT ?,?",
		cmtList.Data.PostId, cmtList.Data.PostId, ctPgNumSize, ctPgSize)
	if err != nil {
		comm.Log("error").Println(err)
		ctx.Fail(500, "100001", "获取评论失败！")
		return
	}
	//处理数据
	var cmt CommentInfo
	var cmtAll []CommentInfo
	for rows.Next() {
		err = rows.Scan(&cmt.CommentId, &cmt.Content, &cmt.Date, &cmt.PenName)
		if err != nil {
			comm.Log("error").Println(err)
			ctx.Fail(500, "100001", "数据解析错误！")
			return
		}
		cmtAll = append(cmtAll, cmt)
	}
	ctx.Success(cmtAll)
}

//回复评论
func ReplyComment(c *gin.Context) {
	ctx := controllers.Context{c}
	ctx.Success("")
}
