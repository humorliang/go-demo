package routers

import (
	"github.com/gin-gonic/gin"
	"ginCms/controllers/user"
	"ginCms/controllers/tag"
	"ginCms/controllers/test"
	"ginCms/controllers/article"
	"ginCms/controllers/comment"
)

//初始化路由映射函数
func SetupRouter(router *gin.Engine) {

	//认证路由
	//authRouterGroup := router.Group("/auth", middleware.JWTAuth())
	//测试认证路由
	//authRouterGroup.GET("/test", test.ReTest)

	//用户路由
	//登录
	router.POST("/login", user.Login)
	//注册
	router.POST("/register", user.Register)

	//分类路由
	//获取分类
	router.GET("/tag", tag.GetTags)
	//添加分类
	router.POST("/tag", tag.AddTag)
	//修改分类
	router.PUT("/tag", tag.UpdateTag)
	//删除分类
	router.DELETE("/tag", tag.DeleteTag)

	//文章路由
	//获取分类文章
	router.GET("/post/list/tag", article.GetTagArticle)
	//获取推荐文章
	router.GET("/post/list/recom", article.GetRecomArticle)
	//获取文章详情
	router.GET("/post/desp", article.GetDescArticle)
	//添加文章
	router.POST("/post", article.AddArticle)
	//删除文章
	router.DELETE("/post", article.DeleteArticle)

	//test路由
	router.GET("/test/log", test.ReTestLogger)

	//文章评论
	//获取评论
	router.GET("/comment", comment.GetPostComment)
	//发表评论
	router.POST("/comment", comment.AddComment)

	//评论回复
	//添加回复评论
	router.POST("/reply", comment.AddReplyComment)

}
