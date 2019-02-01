package main

import (
	"net/http"
	"time"
	"io"
	"fmt"
)

/*
Set a HTTP Cookie Response Header
设计一个请求的cookie信息
*/

func main() {
	//使用默认的defaultServerHTTP
	// DefaultServeMux is the default ServeMux used by Serve.
	// var DefaultServeMux = &defaultServeMux
	http.HandleFunc("/", IndexHandler)
	fmt.Println("****** listen to :8080 port *****")
	http.ListenAndServe(":8080", nil)
}

func AddCookie(w http.ResponseWriter, name string, value string) {
	//过期时间
	exprie := time.Now().AddDate(0, 0, 1)
	//定义cookie信息
	cookie := http.Cookie{
		Name:    name,
		Value:   value,
		Expires: exprie,
		//Domain:
	}
	//设置cookie信息
	http.SetCookie(w, &cookie)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	AddCookie(w, "sessionId", "123456")
	io.WriteString(w, "set Cookie")
}
