package main

import (
	"net"
	"log"
)

//全局变量
type client chan<- string //只写通道类型
var (
	entering = make(chan client) //用户进入
	leaving  = make(chan client) //用户离开
	messages = make(chan string) //消息存储
)

//服务端
func main() {
	//监听
	lister, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal("lister error:", err)
	}
	//处理客户端
	go broadcaster()
	//接受连接
	for {
		con, err := lister.Accept()
		if err != nil {
			log.Println("con error:", con)
			continue
		}
		go handlecon()
	}
}

//客户端管理程序 广播处理
func broadcaster() {
	//定义所有客户端
	clients := make(map[client]bool)
	for {
		//多个case准备就绪时会随机选择一个
		select {
		//全局消息通道读消息
		case msg := <-messages:
			//通知客户端
			for cl := range clients {
				cl <- msg
			}

		}
	}
}

//客户端请求处理程序
func handlecon() {

}
