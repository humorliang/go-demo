package main

import (
	"time"
	"fmt"
)

/*
通道选择器：
Go 的通道选择器 让你可以同时等待多个通道操作。Go 协程和通道以及选择器的结合是 Go 的一个强大特性
各个通道将在若干时间后接收一个值，这个用来模拟例如并行的 Go 协程中阻塞的 RPC 操作
*/
func main() {
	//创建通道
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	//创建结束

	//通道里面传值
	go func() {
		time.Sleep(time.Second * 3)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 4)
		c2 <- "two"
	}()
	// for循环配合select进行通道选择
	for i := 0; i < 2; i++ {
		fmt.Println("i:", i)
		select {
		case msg1 := <-c1:
			fmt.Println("received:", msg1)
		case msg2 := <-c2:
			fmt.Println("recevied:", msg2)
		}
	}

}
