package main

/*
这是一个我们将要在 Go 协程中运行的函数。
done 通道将被用于通知其他 Go 协程这个函数已经工作完毕。
*/
import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func main() {
	done := make(chan bool, 1) //创建通道
	go worker(done)
	//等待通道传值，主程序进行阻塞
	if <-done {
		fmt.Println("worker 运行完毕")
	}

}
