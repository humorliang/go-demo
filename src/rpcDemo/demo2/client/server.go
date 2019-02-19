package main

import (
	"net/rpc/jsonrpc"
	"log"
	"fmt"
)

//请求参数
type Param struct {
	Title string
}

func main() {
	fmt.Println("********* client calll *********")
	//连接RPC服务
	client, err := jsonrpc.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal("client connnet fail:", client)
	}
	args := Param{Title: "t1"}
	var reply interface{}
	//
	err = client.Call("Article.GetInfo", args, &reply)
	if err != nil {
		log.Fatal("call Fail:", err)
	}
	fmt.Println(reply)
}
