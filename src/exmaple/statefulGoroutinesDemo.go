package main

import (
	"math/rand"
	"sync/atomic"
	"time"
	"fmt"
)

/*
go
另一个选择是使用内置的 Go协程和通道的的同步特性来达到同样的效果。
这个基于通道的方法和 Go 通过通信以及 每个 Go 协程间通过通讯来共享内存，
确保每块数据有单独的 Go 协程所有的思路是一致的。

state 将被一个单独的 Go 协程拥有。这就能够保证数据在并行读取时不会混乱。
*/

/*
理解：
自定义数据类型：readOp ,writeOp

通道（读写任务）：reads ,writes

状态管理器： 读写任务中取任务，将信息返回给任务数据中的resp通道

多个读写任务：创建读写数据体，放入相关通道，从数据体中的resp通道读取数据
*/


//定义两个读写类型
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	//定义计数器
	var ops int64
	//定义读写类型通道
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	//创建拥有stat的go协程的状态管理
	go func() {
		//
		var stat = make(map[int]int)
		for true {
			//从读写通道获取读写元素，
			select {
			//读的通道获取一个readOp类型值，写入读请求的resp通道
			case read := <-reads:
				read.resp <- stat[read.key]
				//写的通道获取一个writeOp类型值，
			case write := <-writes:
				stat[write.key] = write.val
				write.resp <- true
			}
		}
	}()
	//创建读任务
	for r := 0; r < 100; r++ {
		go func() {
			for true {
				//readOp类型的读数据
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				//read数据发送到reads通道
				reads <- read
				//从reads通道里resp进行读数据
				<-read.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}
	//创建写任务
	for w := 0; w < 10; w++ {
		go func() {
			for true {
				//创建writeOp类型数据
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				//write数据放入writes通道
				writes <- write
				//从read的resp通道中取数据
				<-write.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	//
	time.Sleep(time.Second)
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("opsFinal:", opsFinal)
}
