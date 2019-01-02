package article

//func Index(ctx *gin.Context) {
//	var p controllers.PostInfo
//	//将结构体与json绑定
//	err := ctx.BindJSON(&p)
//	//数据json绑定是否成功
//	if err == nil {
//		ctx.JSON(http.StatusOK,
//			gin.H{
//				"code": "200",
//				"msg":  "success",
//				"data": p,
//			})
//	} else {
//		ctx.JSON(http.StatusBadRequest, gin.H{
//			"code": "40001",
//			"msg":  "文章数据有误！",
//			"data": "",
//		})
//	}
//}
