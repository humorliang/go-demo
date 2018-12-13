package main

import (
	"sync/atomic"
	"time"
	"fmt"
	"runtime"
)

/*
原子计数：
我们在工作池的例子中碰到过，但是还是有一些其他的方法来管理状态的。
这里我们将看看如何使用 sync/atomic包在多个 Go 协程中进行 原子计数 。
*/
func main() {
	//定义一个计数器
	var ops uint64 = 0

	//开启50个协程进行操作
	for i := 0; i < 50; i++ {
		go func() {
			//进行循环计数加一
			for {
				//使计数器自动增加
				atomic.AddUint64(&ops, 1)
				//允许其他go协程执行
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Second)
	//在计数器还在被其它 Go 协程更新时，安全的使用它，我们通过 LoadUint64 将当前值的拷贝提取到 opsFinal中。
	opsFinal := atomic.LoadUint64(&ops) //opsFinal: 4657298
	fmt.Println("opsFinal:", opsFinal)

}
