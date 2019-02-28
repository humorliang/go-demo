package dbops

import (
	"database/sql"
	"fmt"
	"blog/comm/setting"
	"log"
	_ "github.com/go-sql-driver/mysql"
) //mysql数据库连接驱动

var dbCon *sql.DB

func init() {
	fmt.Println("db init")
	var err error
	if setting.DbCfg.User != "" {
		fmt.Println("ser")
		//数据库链接地址
		dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
			setting.DbCfg.User,
			setting.DbCfg.Password,
			setting.DbCfg.Host,
			setting.DbCfg.Name, )
		dbCon, err = sql.Open(setting.DbCfg.Type, dataSource)
		if err != nil {
			log.Fatal("mysql connect error:", err)
		}
		//测试连接
		err = dbCon.Ping()
		if err != nil {
			log.Fatal("mysql ping connect error:", err)
		}
	}else {
		dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
			"root", "123456", "127.0.0.1:3306", "ll_cms")
		dbCon, err = sql.Open("mysql", dataSource)
		if err != nil {
			log.Fatal("mysql connect2 error:", err)
		}
		//测试连接
		err = dbCon.Ping()
		if err != nil {
			log.Fatal("mysql ping connect error:", err)
		}
	}
}
