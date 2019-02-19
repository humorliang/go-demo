package main

import (
	"net"
	"github.com/smallnest/rpcx/log"
	"fmt"
)

func main() {
	//生成tcp地址
	addr, err := net.ResolveTCPAddr("tcp", ":8000")
	//监听
	lister, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal("lister is fail:", err)
	}
	//192.168.199.191
	fmt.Println("*********** tcp listen 8000 *********")
	for true {
		//连接
		con, err := lister.Accept()
		if err != nil {
			log.Info(err)
			continue
		}
		//并发处理链接
		go func() {
			fmt.Println("请求client地址：", con.RemoteAddr())
			//写内容
			con.Write([]byte("hello client"))
			con.Close()
		}()
	}
}
