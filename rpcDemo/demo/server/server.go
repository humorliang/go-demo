package main

import (
	"rpcDemo/demo/gprotobuf"
	"net/rpc"
	"github.com/smallnest/rpcx/log"
	"net"
	"fmt"
)

//定义服务
type TestServer struct {
}

//定义服务方法
//*gprotobuf.Hello  使用proto数据体
func (this *TestServer) Do(request *gprotobuf.Hello, reply *gprotobuf.Hello) error {
	//获取proto数据的Value值
	reply.Value = "this is proto data:" + request.GetValue()
	return nil
}

func main() {
	//注册服务
	err := rpc.RegisterName("Hello", new(TestServer))
	if err != nil {
		log.Fatal("Hello Server Register Fail:", err)
	}
	//建立监听 创建socket
	listenr, err := net.Listen("tcp", ":8080")
	fmt.Println("******* hello server listen tcp:8080 *********")
	for {
		//接受连接
		con, err := listenr.Accept()
		if err != nil {
			log.Fatal("Con Fail:", err)
		}
		//RPC进行循环Read and Writer
		//并且通过rpc.ServeConn函数在该TCP链接上为对方提供RPC服务
		rpc.ServeConn(con)
	}
}
