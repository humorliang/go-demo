package main

import "fmt"

/*
channel：通道
使用 make(chan val-type) 创建一个新的通道。通道类型就是他们需要传递值的类型。
使用 channel <- 语法 发送 一个新的值到通道中。
使用 <-channel 语法从通道中 接收 一个值。
默认发送和接收操作是阻塞的，直到发送方和接收方都准备完毕。
*/
func main() {
	cha := make(chan string)
	go func() { cha <- "ping" }()
	msg := <-cha
	fmt.Println(msg)
}
