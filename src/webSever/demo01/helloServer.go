package main

import (
	"net/http"
	"log"
	"fmt"
)

//视图函数
// @*http.Request：请求包读取操作函数
// @http.ResponseWriter: 响应包操作函数
func sayHello(w http.ResponseWriter, r *http.Request) {
	//收到的请求首先必须要进行参数解析，否则后续的请求值都为空
	r.ParseForm()

	fmt.Println(r.PostForm==nil)
}

func main() {
	//设置访问路由
	http.HandleFunc("/", sayHello)
	//设置服务器监听端口
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		//错误日志信息
		log.Fatal("ListenAndServe:", err)
	}
}
