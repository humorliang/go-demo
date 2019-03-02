package main

import "fmt"

/*
通道方向：
当使用通道作为函数的参数时，你可以指定这个通道是不是只用来发送或者接收值。
这个特性提升了程序的类型安全性。
*/

func ping(pings chan<- string, msg string) {
	//只传消息，消息传入通道
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	// 从一个通道取值，传入另外一个通道，由它进行发送
	msg := <-pings //接收值
	pongs <- msg   //传入另一个通道
}
func main() {
	pings := make(chan string, 1) //定义通道并设置缓存数
	pongs := make(chan string, 1)
	ping(pings, "hello")
	pong(pings, pongs)
	//通道取值
	fmt.Println(<-pongs) //从pongs通道取值 //hello
}
