package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //mysql数据库连接驱动
	"ginCms/comm"
)

var DbCon *sql.DB

func init() {
	var err error
	//open会自动创建一个数据库连接池，只有在query,exce才会去连接
	//username:password@protocol(address)/dbname?param=value
	DbCon, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/android_app?charset=utf8")
	if err != nil {
		//fmt.Println(err)
		comm.Log("error").Fatal("数据库连接池错误：", err)
	}

	//不能够打开一个连接
	err = DbCon.Ping()
	if err != nil {
		//fmt.Println(err)
		comm.Log("error").Fatal("数据库连接错误：", err)
	}
}

//func main()  {
//	//关闭连接当函数结束时
//	defer db.DbCon.Close()
//	//err := db.DbCon.Prepare() //定义一个sql解析器
//	//返回一个记录结果集指针*rows
//	rows, err := db.DbCon.Query("SELECT id,username FROM user")
//	if err != nil {
//		log.Print(err)
//	}
//	//用户信息
//	type userInfo struct {
//		id       int
//		username string
//	}
//	//切片存储信息
//	userSlice := []userInfo{}
//	//遍历查询信息 下一个记录的bool值
//	for rows.Next() {
//		//创建user类型
//		user := userInfo{}
//		//用户信息添加到容器
//		rows.Scan(&user.id, &user.username)
//		userSlice = append(userSlice, user)
//	}
//	fmt.Println(userSlice[0].username)
//}
