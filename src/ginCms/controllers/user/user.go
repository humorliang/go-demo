package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ginCms/db"
	"fmt"
)

//登录控制函数
func Login(ctx *gin.Context) {
	//获取表单信息
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	//打开表单信息
	//db := db2.DbCon
	//判断用户名和密码
	if username == "test" && password == "123" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
			"data": "登陆成功！",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "账号或密码不正确！",
			"data": "",
		})
	}
}

//注册控制函数
func Register(ctx *gin.Context) {
	//获取注册表单信息
	username := ctx.PostForm("username")
	penname := ctx.PostForm("penName")
	pwd := ctx.PostForm("password")
	email := ctx.PostForm("email")
	fmt.Println(penname, pwd, email)
	if b, err := isQueryUserByUserName(username); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": "400003",
			"msg":  "数据库操作有误",
			"data": "",
		})
	} else {
		if b {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "fails",
				"data": "用户名已纯在！",
			})
		} else {
			dbcon:=db.DbCon
			dbcon.Exec("INSERT INTO user(id,username,p)")
			ctx.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "success",
				"data": "注册成功！",
			})
		}

	}

}

func isQueryUserByUserName(username string) (h bool, err error) {
	//打开数据库
	dbcon := db.DbCon
	rows, err := dbcon.Query("SELECT id FROM user WHERE username = ?", username)
	//查询字段解析
	var id int64
	//var idContain []int64
	for rows.Next() {
		err = rows.Scan(&id)
		//idContain = append(idContain, id)
	}
	if id == 0 {
		return false, err
	} else {
		return true, err
	}
}
