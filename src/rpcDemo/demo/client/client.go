package main

import (
	"net/rpc"
	"github.com/smallnest/rpcx/log"
	"rpcDemo/demo/gprotobuf"
	"fmt"
)

func main() {
	//1. 建立连接
	client, err := rpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("建立连接失败")
	}
	//2. 定义请求和回复数据类型
	var reply = gprotobuf.Hello{}
	var param = gprotobuf.Hello{
		Value: "params",
	}


	fmt.Println("******* client dail 127.0.0.1:8080 *********")
	//3. 开始调用
	err = client.Call("Hello.Do", param, &reply)
	if err != nil {
		log.Fatal("Hello.Do server Call server Fail:", err)
	}
	fmt.Println(reply.Value)
}
