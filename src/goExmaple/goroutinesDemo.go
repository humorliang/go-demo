package main

import "fmt"

/*
使用 go f(s) 在一个 Go 协程中调用这个函数。这个新的 Go 协程将会并行的执行这个函数调用。
你也可以为匿名函数启动一个 Go 协程。
*/

func fStr(from string) {
	for i := 0; i < 13; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	//
	fStr("direct")

	go fStr("goroutine1")

	go func(from string) {
		for i := 0; i < 13; i++ {
			fmt.Println(from, ":", i)
		}
	}("goroutine2")

	var input string
	fmt.Scanln(&input) //输入内容进行停止
	fmt.Println("done")
}
