package main

import (
	"time"
	"fmt"
)

/*
我们常常需要在后面一个时刻运行 Go 代码，或者在某段时间间隔内重复运行。
Go 的内置 定时器 和 打点器 特性让这些很容易实现。
你告诉定时器需要等待的时间，然后它将提供一个用于通知的通道。
*/

func main() {
	// 定义一个定时器通知通道  2s后执行 返回时间通道
	time1 := time.NewTimer(time.Second * 2)

	<-time1.C //此处主线程阻塞
	fmt.Println("timer1 expired")

	//定义一个定时器通知通道
	timer2 := time.NewTimer(time.Second * 3)
	go func() {
		<-timer2.C
		fmt.Println("timer2 expired")
	}()
	//在定时器没有开启前取消定时器
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer2 stop")
	}

}
