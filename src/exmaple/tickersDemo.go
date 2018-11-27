package main

import (
	"time"
	"fmt"
)

/*
打点器 则是当你想要在固定的时间间隔重复执行准备的。
打点器和定时器的机制有点相似：一个通道用来发送数据。
这里我们在这个通道上使用内置的 range 来迭代值每隔500ms 发送一次的值。
*/
func main() {
	//定义一个打点器
	ticker := time.NewTicker(time.Second * 1) //每1s执行一次
	done := make(chan bool)
	//
	go func() {
		//时间通道进行读取
		for t := range ticker.C {
			fmt.Println("ticker run:", t) //答应时间节点
		}
		done <- true
	}()

	//打点器停止时间 主线程阻塞时间
	//time.Sleep(time.Millisecond*1600)

	//主线程一直阻塞
	if <-done {
		ticker.Stop() //停止
		fmt.Println("ticker stop")
	}
}
