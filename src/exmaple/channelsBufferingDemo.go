package main

import "fmt"

/*
默认通道是 无缓冲 的，这意味着只有在对应的接收（<- chan）通道准备好接收时，
才允许进行发送（chan <-）。可缓存通道允许在没有对应接收方的情况下，缓存限定数量的值。
*/
func main() {
	msg := make(chan string, 2)

	msg <-"buffer"
	msg <-"channels"


	fmt.Println(<-msg)
	fmt.Println(<-msg)
}
