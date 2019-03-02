package main

import (
	"net/http"
	"io"
)

/*
Get the HTTP Method from a Request
获取请求方法
*/
func main() {
	/*
		istenAndServe使用指定的监听地址和处理器启动一个HTTP服务端。
		处理器参数通常是nil，这表示采用包变量DefaultServeMux作为处理器。
		Handle和HandleFunc函数可以向DefaultServeMux添加处理器。
	*/
	http.HandleFunc("/", http.HandlerFunc(ExampleHandler))
	http.ListenAndServe(":8080", nil)
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		io.WriteString(w, "this is a get request")
	} else if r.Method == http.MethodPost {
		io.WriteString(w, "this is a post request")
	} else {
		io.WriteString(w, "this is  a "+r.Method+" request")
	}
}
