package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //mysql数据库连接驱动
	"ginCms/comm"
	"fmt"
	"ginCms/comm/setting"
)

var Con *sql.DB

func Setup() {
	var err error
	//open会自动创建一个数据库连接池，只有在query,exce才会去连接，是一个惰性连接
	//
	//username:password@protocol(address)/dbname?param=value
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name)
	fmt.Println(dataSource)
	Con, err = sql.Open("mysql", dataSource)
	if err != nil {
		//fmt.Println(err)
		comm.Log("error").Fatal("数据库连接错误：", err)
	}

	//不能够打开一个连接
	err = Con.Ping()
	if err != nil {
		//fmt.Println(err)
		comm.Log("error").Fatal("数据库连接测试错误：", err)
	}
}

//通过用户名进行数据库查询是否存在
func IsQueryUserByUserName(username string, con *sql.DB) (h bool, err error) {
	//打开数据库
	rows, err := con.Query("SELECT id FROM user WHERE user_name = ?", username)
	if err != nil {
		comm.Log("error").Println("用户名查询错误：", err)
	}
	//释放链接
	defer rows.Close()
	//查询字段解析
	var id int64
	//var idContain []int64
	for rows.Next() {
		err = rows.Scan(&id)
		//idContain = append(idContain, id)
		if err != nil {
			comm.Log("error").Println("用户名查询数据解析错误：", err)
		}
	}
	if id == 0 {
		return false, err
	} else {
		return true, err
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
