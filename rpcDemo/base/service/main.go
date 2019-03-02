package main

import (
	"github.com/smallnest/rpcx/server"
	"rpcDemo/base/service/sv"
	"fmt"
)

func main() {
	//新建一个web服务
	s := server.NewServer()
	//注册服务  这里简单使用了服务的 类型名称 作为 服务名。
	//s.Register(new(sv.Arith),"")
	s.RegisterName("Arith", new(sv.Arith), "")
	fmt.Println("******** Airth Server 127.0.0.1:8989 ************")
	s.Serve("tcp", "127.0.0.1:8989")
}
