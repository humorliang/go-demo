package main

import (
	"fmt"
	"time"
)

/*
通道关闭： close(channels)
for i := range c能够不断的读取channel里面的数据，直到该channel被显式的关闭。
上面代码我们看到可以显式的关闭channel，生产者通过内置函数close关闭channel。
关闭channel之后就无法再发送任何数据了，在消费方可以通过语法v, ok := <-ch测试channel是否被关闭。
如果ok返回false，那么说明channel已经没有任何数据并且已经被关闭。
*/
func main() {
	//数据通道
	jobs := make(chan int)
	done := make(chan bool)

	//消费者
	go func() {
		for {
			//读数据 并判断通道是否关闭
			job, ok := <-jobs
			if ok {
				fmt.Println("recived job:", job)
			} else {
				fmt.Println("recived all jobs")
				done <- true
				return
			}
		}
	}()

	//生产者
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second * 1)
		jobs <- i
		fmt.Println("sent job:", i)
	}
	close(jobs) //关闭通道

	//主线阻塞
	<-done
}
