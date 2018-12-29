package user

import (
	"github.com/gin-gonic/gin"
	"ginCms/db"
	"ginCms/comm"
	"ginCms/controllers"
	"fmt"
	"ginCms/utils"
)

//用户类
type User struct {
	Id       int    `json:"id"`
	Username string `json:"username",form:"username"`
	Password string `json:"password",form:"password"`
	PenName  string `json:"pen_name",form:"pen_name"`
	Email    string `json:"email",form:"email"`
}

//登录控制函数
func Login(c *gin.Context) {
	//context类型定义
	ctx := controllers.Context{c}
	//获取表单信息
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	//判断用户名和密码
	if h, err := db.IsQueryUserByUserName(username, db.Con); h {
		//查询异常判断
		if err != nil {
			comm.Log("error").Println(err)
			ctx.Fail(500, 10002, "数据查询失败！")
		} else {
			//数据库结构解析
			var user User
			var userAll []User
			//查询数据没有
			//如果没有结果会返回error
			//err := db.Con.QueryRow("SELECT id,user_name,pen_name,email FROM user WHERE user_name=? AND pass_word=?", username, password)
			rows, err := db.Con.Query("SELECT id,user_name,pen_name,email FROM user WHERE user_name=? AND pass_word=?", username, utils.Md5Encrypt(password))
			//查询异常判断
			if err != nil {
				comm.Log("error").Println(err)
				ctx.Fail(500, 10002, "数据查询失败！")
			} else {
				//fmt.Println("2")
				//fmt.Println(rows)
				//查询结果遍历
				for rows.Next() {
					fmt.Println("3")
					err := rows.Scan(&user.Id, &user.Username, &user.PenName, &user.Email)
					if err != nil {
						comm.Log("error").Println(err)
						ctx.Fail(500, 10002, "数据查询失败！")
					} else {
						userAll = append(userAll, user)
					}
				}

			}
			//对查询数据进行判断
			if len(userAll) == 0 {
				ctx.Fail(401, 10002, "密码错误！")
			} else {


				ctx.Success(gin.H{
					"userId":   userAll[0].Id,
					"userName": userAll[0].Username,
					"pen_name": userAll[0].PenName,
					"email":    userAll[0].Email,
				})
			}
		}
	} else {
		ctx.Fail(400, 40001, "用户名不存在！")
	}
}

//注册控制函数
func Register(c *gin.Context) {
	//context类型定义
	ctx := controllers.Context{c}

	//获取注册表单信息
	username := ctx.PostForm("username")
	penname := ctx.PostForm("penName")
	password := ctx.PostForm("password")
	email := ctx.PostForm("email")

	//用户判断
	if b, err := db.IsQueryUserByUserName(username, db.Con); err != nil {
		comm.Log("error").Println(err)
		ctx.Fail(500, 10002, "数据查询有误")
	} else {
		if b {
			ctx.Fail(400, 40001, "用户名已存在")
		} else {
			//执行一次命令不返回任何结果
			//Exec可以直接处理完自动释放到sql连接池
			re, err := db.Con.Exec("INSERT INTO user(user_name,pen_name,pass_word,email) VALUES(?,?,?,?)", username, penname, utils.Md5Encrypt(password), email)
			if err != nil {
				comm.Log("error").Println("注册信息插入有误：", err)
				ctx.Fail(500, 10005, "数据写入失败！")
			} else {
				//返回插入的ID
				id, err := re.LastInsertId()
				if err != nil {
					comm.Log("error").Println("注册成功返回Id错误：", err)
				}
				ctx.Success(gin.H{
					"userId": id,
				})
			}
		}
	}
}


