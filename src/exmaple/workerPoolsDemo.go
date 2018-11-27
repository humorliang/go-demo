package main

import (
	"fmt"
	"time"
)

/*
利用go协程和通道实现一个工作池
Go中channel可以是只读、只写、同时可读写的。

//定义只读的channel
read_only := make (<-chan int)

//定义只写的channel
write_only := make (chan<- int)

//可同时读写
read_write := make (chan int)
*/

func workerPools(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("work id:", id, "process job", j)
		//模拟耗时任务
		time.Sleep(time.Second * 2)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	result := make(chan int, 100)

	//创建5个工作协程的工作池
	for w := 1; w < 5; w++ {
		//开启5个工作协程
		go workerPools(w, jobs, result)
	}

	//生产者
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)

	////消费者
	//for j := 0; j < 1; j++ {
	//	<-result
	//	//fmt.Println("result:", res)
	//}

	//阻塞主线程
	<-result

}
