package main

import (
	"flag"
	"fmt"
)

/*
Passing Arguments into your Application
解析命令行参数
*/

//返回命令行指针
var runTaskA = flag.Bool("taska", false, "this is run task a")

func main() {
	//解析命令行
	flag.Parse()

	if *runTaskA {
		taskA()
	} else {
		taskB()
	}
}

func taskA() {
	fmt.Println("运行taskA任务")
}

func taskB() {
	fmt.Println("运行taskB任务")
}
