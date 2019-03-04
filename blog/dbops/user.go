package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/humorliang/go-demo/blog/controller/data"
	"github.com/humorliang/go-demo/blog/utils/encryption"
)

/*
用户数据操作API
*/
//用户是否存在
func UserIsExist(username string) (bool, error) {
	//sql语句解析
	stmt, err := dbCon.Prepare("SELECT id FROM ll_admin WHERE username=?")
	defer stmt.Close()
	if err != nil {
		return false, err
	}
	//执行
	res, err := stmt.Exec(username)
	if err != nil {
		return false, err
	}
	num, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	if num == 0 {
		return false, nil
	}
	return true, nil
}

//获取用户信息
func GetUserInfo(id int) (user data.Admin, err error) {
	//sql解析
	sql := `SELECT id,username,realname,email,mobile,description,
			login_time,login_count,status,parent_id,is_super,is_delete,
			create_time,update_time FROM ll_admin WHERE id=? `
	rows, err := dbCon.Query(sql, id)
	if err != nil {
		return data.Admin{}, err
	}
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Username, &user.Realname,
			&user.Email, &user.Mobile, &user.Description, &user.LoginTime,
			&user.LoginCount, &user.Status, &user.ParentId, &user.IsSuper,
			&user.IsDelete, &user.CreateTime, &user.UpdateTime)
		if err != nil {
			return data.Admin{}, err
		}
	}
	defer rows.Close()
	return user, nil
}

//获取用户列表
func GetUserList(pageNum int, pageSize int) (users []data.Admin, count int, err error) {
	//sql解析
	sql := `SELECT id,username,realname,email,mobile,description,
			login_time,login_count,status,parent_id,is_super,is_delete,
			create_time,update_time FROM ll_admin limit ?,? `
	sql2 := "select FOUND_ROWS() as count limit ? "
	rows, err := dbCon.Query(sql, (pageNum-1)*pageSize, pageSize)
	rows2, err := dbCon.Query(sql2, 1)
	if err != nil {
		return nil, 0, err
	}
	for rows.Next() {
		var user data.Admin
		err = rows.Scan(&user.Id, &user.Username, &user.Realname,
			&user.Email, &user.Mobile, &user.Description, &user.LoginTime,
			&user.LoginCount, &user.Status, &user.ParentId, &user.IsSuper,
			&user.IsDelete, &user.CreateTime, &user.UpdateTime)
		if err != nil {
			return nil, 0, err
		}
	}
	for rows2.Next() {
		err = rows2.Scan(&count)
		if err != nil {
			return nil, 0, err
		}
	}
	defer rows.Close()
	defer rows2.Close()
	return users, count, nil
}

//用户账号验证
func UserAuth(user data.Admin) (bool, error) {
	var id int
	//sql语句
	sql := "SELECT id FROM ll_admin WHERE username=? AND password=? "
	err := dbCon.QueryRow(sql, user.Username, user.Password).Scan(&id)
	if err != nil {
		return false, err
	}
	return true, err
}

//添加用户信息
func AddUser(user data.Admin) (err error) {
	//sql
	sql := `INSERT INTO ll_admin (username,password,realname,email,
		mobile,description,parent_id)VALUES(?,?,?,?,?,?,?)`
	stmt, err := dbCon.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return err
	}
	res, err := stmt.Exec(user.Username, encryption.Md5Encrypt(user.Password), user.Realname,
		user.Email, user.Mobile, user.Description, user.ParentId)
	if n, err := res.RowsAffected(); err != nil {
		return err
	} else {
		if n == 0 {
			return fmt.Errorf("insert user fail")
		}
	}
	return nil
}
