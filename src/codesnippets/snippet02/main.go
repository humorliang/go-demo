package main

import (
	"sync"
	"fmt"
)

/*
Run Code Once on First Load (Concurrency Safe)
只加载一次代码(并发安全)：单例模式
*/
var doOnce sync.Once

func main() {
	doSomething()
	doSomething()
}

func doSomething() {
	doOnce.Do(func() {
		fmt.Println("Run Once - first run function loading...")
	})
	fmt.Println("Run this every time ")
}
