/*
ws 客户端
*/
package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
	"github.com/gorilla/websocket"
)

//ws地址
var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	//解析命令行参数
	flag.Parse()
	log.SetFlags(0)

	//中断信号
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	//请求的url
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
	log.Printf("connecting to %s", u.String())

	//ws拨号装置
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	//延迟函数
	defer c.Close()
	//完成数
	done := make(chan struct{})

	//go程
	go func() {
		//关闭通道
		defer close(done)
		//循环监听读取消息
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()
	// 定义打点器
	ticker := time.NewTicker(time.Second)
	//结束打点
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
			//开始执行打点器
		case t := <-ticker.C:
			//写入数据
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
