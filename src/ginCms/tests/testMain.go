package main

func main() {
	//comm.Log()
	//log.Print("test")
	//path := "/data/log/test.log"
	////path := "test.log"
	//pathSlice := strings.Split(path, "/")
	//pathStr:=strings.Join(pathSlice[:len(pathSlice)-1],"/")
	//fmt.Println(pathStr)
	//fmt.Println(pathSlice[len(pathSlice)-1])

	//utils.OpenFile("../data/log/test.log")
	//utils.OpenFile("test2.log")
	//fmt.Println(time.Now().String())
	//fmt.Println(strings.Split(time.Now().String()," ")[0])
	//fmt.Println(db.DbCon)

	////关闭连接当函数结束时
	//defer db.DbCon.Close()
	////err := db.DbCon.Prepare() //定义一个sql解析器
	////返回一个记录结果集指针*rows
	//rows, err := db.DbCon.Query("SELECT id,username FROM user")
	//if err != nil {
	//	log.Print(err)
	//}
	////用户信息
	//type userInfo struct {
	//	id       int
	//	username string
	//}
	////切片存储信息
	//userSlice := []userInfo{}
	////遍历查询信息 下一个记录的bool值
	//for rows.Next() {
	//	//创建user类型
	//	user := userInfo{}
	//	//用户信息添加到容器
	//	rows.Scan(&user.id, &user.username)
	//	userSlice = append(userSlice, user)
	//}
	//fmt.Println(userSlice[0].username)
}
