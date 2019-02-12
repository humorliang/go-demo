package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
) //mysql数据库连接驱动

func main() {
	con, err := sql.Open("mysql", "root:123456@tcp(192.168.199.192:6666)/mysql?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	err = con.Ping()
	fmt.Println(err)
}
