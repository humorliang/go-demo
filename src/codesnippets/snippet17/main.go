package main

import (
	"net/http"
	"io"
	"log"
	"fmt"
)

/*
Apply Middleware to Your Route Handlers
应用一个中间件，在你的路由中
*/

const ApiKey = "auth_key"

func main() {
	//创建一个节点路由
	http.Handle("/", Middleware(
		http.HandlerFunc(ExampleHandler),
		AuthMiddleware,
	))
	fmt.Println("***** listen post 8080 *****")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

//定义中间件list
func Middleware(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	//对于中间件进行嵌套封装
	for _, mw := range middleware {
		h = mw(h)
	}
	return h
}

//认证中间件
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqKey := r.URL.Query().Get("key")
		if len(reqKey) == 0 || reqKey != ApiKey {
			w.Header().Add("Content-Type", "application/json")
			//写入状态响应状态码
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, `{"error":"invalid_key"}`)
			return
		}
		//执行下一个请求句柄
		next.ServeHTTP(w, r)
	})
}

//路径响应函数
func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{"status":"ok"}`)
}
