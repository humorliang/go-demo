package comment

import (
	"github.com/gin-gonic/gin"

	"fmt"
	"github.com/humorliang/go-demo/ginCms/controllers"
	"github.com/humorliang/go-demo/ginCms/comm"
	"github.com/humorliang/go-demo/ginCms/db"
	"github.com/humorliang/go-demo/ginCms/comm/setting"
)

type Comment struct {
	Data struct {
		Content   string `json:"content" binding:"required"`
		UserId    int    `json:"userId" binding:"required"`
		ArticleId int    `json:"articleId" binding:"required"`
	} `json:"car"`
}

type CommentList struct {
	Data struct {
		PostId  int `json:"postId" binding:"required"`
		PageNum int `json:"pageNum" binding:"required"`
	} `json:"car"`
}

//评论信息
type CommentInfo struct {
	CommentId int     `json:"commentId"`
	Content   string  `json:"content"`
	Date      string  `json:"date"`
	PenName   string  `json:"penName"`
	ReplyAll  []Reply `json:"replyCommentAll"`
}

//回复评论
type Reply struct {
	ReplyId int    `json:"replyId"`
	Content string `json:"content"`
	Date    string `json:"date"`
	PenName string `json:"penName"`
}
type ReplyJson struct {
	Data struct {
		CommentID int    `json:"commentId" binding:"required"`
		Content   string `json:"content" binding:"required"`
		UserID    int    `json:"userId" binding:"required"`
	} `json:"car"`
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
	rows, err := db.Con.Query("SELECT c.id,c.content,c.date,u.pen_name "+
		"FROM comment c,user u,reply r WHERE u.id=c.user_id AND c.article_id=? "+
		"ORDER BY c.date LIMIT ?,?",
		cmtList.Data.PostId, ctPgNumSize, ctPgSize)
	defer rows.Close()
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
		//局部作用域
		//定义回复变量
		var reply Reply
		var replyAll []Reply
		// 回复评论SQL查询
		rowsReply, err := db.Con.Query("SELECT r.id,r.content,r.date,u.pen_name "+
			"FROM user u,reply r WHERE u.id=r.user_id AND r.comment_id=? "+
			"ORDER BY r.date", cmt.CommentId)
		defer rowsReply.Close()
		if err != nil {
			comm.Log("error").Println(err)
			ctx.Fail(500, "10001", "回复评论获取失败！")
			return
		}
		// 遍历回复评论查询
		for rowsReply.Next() {
			//遍历回复评论
			err = rowsReply.Scan(&reply.ReplyId, &reply.Content, &reply.Date, &reply.PenName)
			if err != nil {
				comm.Log("error").Println(err)
				ctx.Fail(500, "10001", "数据解析错误！")
				return
			}
			replyAll = append(replyAll, reply)
		}
		//fmt.Println("评论：", cmt.CommentId)
		//fmt.Println("回复：", replyAll)
		// 将评论给文章
		cmt.ReplyAll = replyAll
		// 添加到文章对象
		cmtAll = append(cmtAll, cmt)
	}
	fmt.Println(cmtAll)
	ctx.Success(cmtAll)
}

//添加回复评论
func AddReplyComment(c *gin.Context) {
	ctx := controllers.Context{c}
	rePJ := &ReplyJson{}
	err := ctx.BindJSON(rePJ)
	if err != nil {
		comm.Log("error").Println(err)
		ctx.Fail(400, "10001", "回复评论失败！")
		return
	}
	//SQL执行
	res, err := db.Con.Exec("INSERT INTO reply (content,user_id,comment_id) VALUES (?,?,?)",
		rePJ.Data.Content, rePJ.Data.UserID, rePJ.Data.CommentID)
	if err != nil {
		comm.Log("error").Println(err)
		ctx.Fail(500, "10001", "回复评论失败！")
		return
	}
	//获取回复ID
	id, err := res.LastInsertId()
	if err != nil {
		comm.Log("error").Println(err)
		ctx.Fail(500, "10001", "回复评论失败！")
		return
	} else {
		ctx.Success(gin.H{
			"ReplyId": id,
		})
	}
}
