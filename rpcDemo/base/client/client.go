package main

import (
	"github.com/smallnest/rpcx/client"
	"rpcDemo/base/service/sv"
	"fmt"
	"context"
	"github.com/smallnest/rpcx/log"
)

/*
客户端：
*/
func main() {
	// 1. 客户端的服务发现连接方式
	fmt.Println("********* server discover 127.0.0.1:8989 ********** ")
	d := client.NewPeer2PeerDiscovery("tcp@127.0.0.1:8989", "")

	// 2. 创建xClient客户端 客户端的配置
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	// 3. 定义请求 参数
	args := &sv.Args{
		A: 10,
		B: 20,
	}

	// 4.定义响应对象 服务结果
	reply := &sv.Reply{}

	// 5.调用远程服务 并同步返回结果
	//err := xclient.Call(context.Background(), "Mul", args, reply)
	//if err != nil {
	//	log.Fatal("failed to call :%v", err)
	//}
	//fmt.Printf("sync call: %d * %d = %d", args.A, args.B, reply.C)

	// 5.1 调用远程服务 并异步返回结果
	call, err := xclient.Go(context.Background(), "Mul", args, reply, nil)
	if err != nil {
		log.Fatal("failed to call :%v", err)
	}
	replyCall := <-call.Done
	if replyCall.Error != nil {
		log.Fatal("failed to replyCall:%v", replyCall.Error)
	} else {
		fmt.Printf("async call: %d * %d = %d", args.A, args.B, reply.C)
	}
}
