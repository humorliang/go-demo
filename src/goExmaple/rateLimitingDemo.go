package main

import (
	"time"
	"fmt"
)

/*

速率限制(英) 是一个重要的控制服务资源利用和质量的途径。
Go 通过 Go 协程、通道和打点器优美的支持了速率限制。
*/
func main() {
	//创建请求通道
	requests := make(chan int, 5)
	//向通道发送数据 当发送的数量大于缓存数时就会出现阻塞
	for i := 0; i < 5; i++ {
		requests <- i
	}
	//关闭请求
	close(requests)
	//定义一个速率限制管理器
	limiter := time.Tick(time.Millisecond * 200)
	// 迭代速率限制器
	for req := range requests {
		//管理器中读取 进行速率限制 进行阻塞接受
		<-limiter
		fmt.Println("request:", req, time.Now())
	}

	//进行临时速率限制
	//使用通道缓存实现
	burstyLimiter := make(chan time.Time, 3)
	//填充需要临时改变3次的值
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	//go协程定义一个整体的时间控制
	go func() {
		for t := range time.Tick(time.Millisecond * 300) {
			burstyLimiter <- t
		}
	}()
	//脉冲请求任务
	burstyRequests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		burstyRequests <- i
	}
	//关闭脉冲请求任务
	close(burstyRequests)
	//读取任务
	for r := range burstyRequests {
		//脉冲速率接受阻塞
		<-burstyLimiter
		fmt.Println()
		fmt.Println("burstRequest:", r, time.Now())
	}

}
