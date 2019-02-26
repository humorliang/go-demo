package db

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"blog/comm/setting"
	"log"
) //mysql数据库连接驱动

var Con *sql.DB

func MysqlInit() {
	var err error
	//数据库链接地址
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		setting.DbCfg.User,
		setting.DbCfg.Password,
		setting.DbCfg.Host,
		setting.DbCfg.Name, )
	Con, err = sql.Open(setting.DbCfg.Type, dataSource)
	if err != nil {
		log.Fatal("mysql connect error:", err)
	}
	//测试连接
	err = Con.Ping()
	if err != nil {
		log.Fatal("mysql ping connect error:", err)
	}
}
