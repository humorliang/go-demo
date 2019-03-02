package main

/*
HTTP Get with Timeout
请求超时
*/
import (
	"net/http"
	"time"
	"log"
)

func main() {
	//设置链接并设置超时时间
	client := http.Client{
		Timeout: time.Duration(5 * time.Millisecond),
	}
	_, err := client.Get("https://google.com")
	if err != nil {
		log.Fatal(err)
	}
}
