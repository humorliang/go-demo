package main

import (
	"net/rpc"
	"log"
	"net"
	"net/rpc/jsonrpc"
	"fmt"
)

//定义服务
type Article struct {
}

//定义数据体
//请求
type Param struct {
	Title string
}
//响应
type ArticleInfo struct {
	Content string
}

//定义服务方法
func (this *Article) GetInfo(req Param, rsp *ArticleInfo) error {
	if req.Title == "t1" {
		rsp.Content = "this is t1 content"
	} else {
		rsp.Content = "this is other content"
	}
	return nil
}

func main() {
	//注册服务
	err := rpc.RegisterName("Article", new(Article))
	if err != nil {
		log.Fatal("register server fails:", err)
	}
	//监听端口
	lister, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal("listen 8000 port fail:", err)
	}
	fmt.Println("*********** server listen 8000 port ***********")
	//建立链接
	for {
		con, err := lister.Accept()
		if err != nil {
			continue
		}
		//使用json进行处理
		jsonrpc.ServeConn(con)
	}
}
