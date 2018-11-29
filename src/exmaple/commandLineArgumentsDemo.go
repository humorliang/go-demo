package main

import (
	"os"
	"fmt"
)

/*
命令行参数：
os.Args 提供原始命令行参数访问功能。
注意，切片中的第一个参数是该程序的路径，并且 os.Args[1:]保存所有程序的的参数。
*/
func main() {
	//获取命令行所有的参数[./commandLineArgumentsDemo sd asd das asd]
	argWithProg := os.Args
	//获取所有额外参数[sd asd das asd]
	argwithOutProg := os.Args[1:]
	//某个命令行参数asd
	arg := os.Args[2]
	fmt.Println(argWithProg)
	fmt.Println(argwithOutProg)
	fmt.Println(arg)
}
