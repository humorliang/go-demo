package main

import (
	"net/http"
	"log"
	"fmt"
)

/*
HTTP Response Status Codes
获取请求响应码
*/
func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("http response status:", resp.Status)
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		fmt.Println("http response status is 2xx ")
	} else {
		fmt.Println("request fail")
	}
}
