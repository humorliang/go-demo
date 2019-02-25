package main

import (
	"net"
	"time"
	"github.com/smallnest/rpcx/log"
	"fmt"
)

func handler(con net.Conn) {
	//处理完进行连接关闭
	defer con.Close()
	//固定时间 06 01 02 15:04:05
	for true {
		now := time.Now().Format("2006-01-02 15:04:05")
		_, err := con.Write([]byte(now + "\n"))
		if err != nil {
			fmt.Println("con writer err:", err)
		}
		//定义缓冲区
		buffer := make([]byte, 4096)
		//将内容读取到缓冲区，
		rN, err := con.Read(buffer)
		//bt, err := ioutil.ReadAll(con)
		if err != nil {
			fmt.Println("read err:", err)
		}
		fmt.Println("read content:", string(buffer[:rN]))
		time.Sleep(1 * time.Second)
	}
}

func main() {
	//监听
	lister, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal("lister err:", err)
	}
	fmt.Println("*********** listen port:8000 ****************")
	//等待请求接入
	for true {
		con, err := lister.Accept()
		if err != nil {
			log.Info("con error:", err)
			continue
		}
		fmt.Println("Request IP:", con.RemoteAddr())
		go handler(con)
	}
}
