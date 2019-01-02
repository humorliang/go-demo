package routers

import (
	"github.com/gin-gonic/gin"
	"ginCms/controllers/user"
	"ginCms/controllers/tags"
	"ginCms/middleware"
	"ginCms/controllers/test"
	"ginCms/controllers/article"
	"fmt"
)

//初始化路由映射函数
func SetupRouter() *gin.Engine {
	fmt.Println("router")
	router := gin.New()
	//认证路由
	authRouterGroup := router.Group("/auth", middleware.JWTAuth())
	//测试认证路由
	authRouterGroup.GET("/test", test.ReTest)

	//用户路由
	//登录
	router.POST("/login", user.Login)
	//注册
	router.POST("/register", user.Register)

	//分类路由
	//获取分类
	router.GET("/tag", tags.GetTags)
	//添加分类
	router.POST("/tag", tags.AddTag)
	//修改分类
	router.PUT("/tag", tags.UpdateTag)
	//删除分类
	router.DELETE("/tag", tags.DeleteTag)

	//文章路由
	//获取分类文章
	router.GET("/post/list/tag", article.GetTagArticle)
	//获取推荐文章
	router.GET("/post/list/recom", article.GetRecomArticle)
	//获取文章详情
	router.GET("/post", article.GetDescArticle)
	//添加文章
	router.POST("/post", article.AddArticle)
	//删除文章
	router.DELETE("/post", article.DeleteArticle)

	return router
}
