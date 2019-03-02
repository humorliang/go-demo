package main

import "fmt"

/*
channel：通道
使用 make(chan val-type) 创建一个新的通道。通道类型就是他们需要传递值的类型。
使用 channel <- 语法 发送 一个新的值到通道中。
使用 <-channel 语法从通道中 接收 一个值。
默认发送和接收操作是阻塞的，直到发送方和接收方都准备完毕。

channel接收和发送数据都是阻塞的，除非另一端已经准备好，这样就使得Goroutines同步变的更加的简单，
而不需要显式的lock。所谓阻塞，也就是如果读取（value := <-ch）它将会被阻塞，直到有数据接收。
其次，任何发送（ch<-5）将会被阻塞，直到数据被读出。
无缓冲channel是在多个goroutine之间同步很棒的工具。
*/
func main() {
	cha := make(chan string)
	go func() { cha <- "ping" }()
	msg := <-cha
	fmt.Println(msg)
}
