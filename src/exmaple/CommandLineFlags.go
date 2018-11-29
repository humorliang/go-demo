package main

import (
	"flag"
	"fmt"
)

/*
命令行标志是命令行程序指定选项的常用方式。例如，在 wc -l 中，这个 -l 就是一个命令行标志。
Go 提供了一个 flag 包，支持基本的命令行标志解析。我们将用这个包来实现我们的命令行程序示例。

*/
func main() {
	//标志符 值  描述
	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 21, "a int")
	boolPtr := flag.Bool("fork", true, "a bool")

	//用已有变量定义 要取地址
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	//执行命令行解析
	flag.Parse()

	//
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("args:", flag.Args())

}
