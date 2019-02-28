package dbops

import (
	"testing"
	"blog/controller/data"
	"blog/dbops/gmysql/db"
)


func clearTable() {
	con := db.Con
	con.Exec("truncate table ll_admin")
}

func TestMain(m *testing.M) {
	clearTable()
	m.Run()
	clearTable()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add user", testAddUser)
	t.Run("User Exist", testUserIsExist)
	t.Run("User Auth", testUserAuth)
	t.Run("User Info", testGetUserInfo)
	t.Run("User List", testGetUserList)
}

func testAddUser(t *testing.T) {
	u := data.Admin{
		Username: "user2",
		Password: "123456",
		Realname: "张三",
	}
	err := AddUser(u)
	if err != nil {
		t.Error(err)
	}
}
func testUserIsExist(t *testing.T) {

}
func testUserAuth(t *testing.T) {

}
func testGetUserInfo(t *testing.T) {

}
func testGetUserList(t *testing.T) {

}
