package main

import (
	"fmt"
	"time"
)

/*
常规的通过通道发送和接收数据是阻塞的。
然而，我们可以使用带一个 default 子句的 select 来实现非阻塞 的发送、接收，甚至是非阻塞的多路 select。
在select里面还有default语法，select其实就是类似switch的功能，
default就是当监听的channel都没有准备好的时候，默认执行的（select不再阻塞等待channel）。
select {
case i := <-c:
	// use i
default:
	// 当c阻塞的时候执行这里
}
*/
func main() {
	messages := make(chan string, 2)
	signals := make(chan bool)

	go func() {
		for i := 0; i < 5; i++ {
			//传入数据
			time.Sleep(time.Second * 1)
			messages <- string(i)
			fmt.Println(i)
		}
		signals <- true
	}()

	// 读取通道数据
	select {
	case msg := <-messages:
		fmt.Println("received messages:", msg)
	default:
		fmt.Println("no message recevied")
	}

	//写入通道数据
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent messages:", msg)
	default:
		fmt.Println("no sent messages")
	}

	//读取通道消息
	select {
	case res := <-messages:
		fmt.Println("recived messages:", res)
	case sig := <-signals:
		fmt.Println("recived signals:", sig)
	default:
		fmt.Println("no activity")
	}
	/*
	no message recevied
	sent messages: hi
	recived messages: hi
	*/
}
