package main

import (
	"sync"
	"time"
	"fmt"
)

/*
Waiting for Goroutines to Finish with a WaitGroup
go程组等待：主程序等待所有go程 等待组
*/
func main() {

	//go程数
	total := 3
	//go程等待组 
	var wg sync.WaitGroup
	//添加三个等待组
	wg.Add(3)
	for i := 1; i <= total; i++ {
		//运行go程
		go longConcurrentSleepProcess(i, &wg)
	}
	//等待所有go程结束
	wg.Wait()
}
func longConcurrentSleepProcess(sleep int, wg *sync.WaitGroup) {
	//在函数结束时减少一个等待组，表示该线程完成,没有最后会出现：
	// fatal error: all goroutines are asleep - deadlock!
	defer wg.Done()
	time.Sleep(time.Duration(sleep) * time.Second)
	fmt.Println("sleep time ", sleep, " time second!")
}
