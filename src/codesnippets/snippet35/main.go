package main

import (
	"fmt"
	"os"
)

/*
Get and Set Environment Variables
获取和设置一个环境变量
*/
func main() {
	//获取环境变量
	fmt.Println(os.Getenv("GOROOT"))

	//设置环境变量
	os.Setenv("Test", "test go")
	fmt.Println(os.Getenv("Test"))

	//环境变量不存在则为空
	fmt.Println(os.Getenv("LOVE"))

}
