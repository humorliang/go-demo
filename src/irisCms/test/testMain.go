package main

import (
	"os"
	"log"
	"irisCms/utils"
)

func main() {
	//f := newLogFile()
	//f.Write([]byte("main!\n"))
	//today:=time.Now().Format("2016-10-10")
	//fmt.Println(today)
	//month:=time.Now().Month()
	//fmt.Println(month.String())
	//cur,_:=time.Parse("2018-10-19",time.Now().String())
	//fmt.Println(cur)
	//fmt.Println(strings.Split(time.Now().String(), " ")[0] + ".log")
	//utils.CreateAppendWriteFile("./dataTest/","tests.txt",[]byte("hello world\n"))
	//fmt.Println(comm.DbConfigRead("./configs/", "dev.ini"))
	utils.CreateAppendWriteFile("./data/log/info/", utils.TodayFileName(), nil)
}

//创建一个新的日志文件
func newLogFile() *os.File {
	//文件名
	filename := "/Users/datatist001/gopath/data/test2.txt"
	//打开文件以创建写和追加的方式打开
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("newLogFile........")
	f.Write([]byte("hello world!\n"))
	return f
}
