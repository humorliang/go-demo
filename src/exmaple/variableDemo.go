package main

import "fmt"

/*
1. var 声明一个或多个变量
2. 变量没有初始值，会默认为0
3. Go 自动推断初始化变量类型
4. := 初始化变量简写
*/
func main() {
	var a string = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short var"
	fmt.Println(f)
}
