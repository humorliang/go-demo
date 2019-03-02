package main

import "fmt"

/*
闭包：闭包就是能够读取其他函数内部变量的函数。
好处：一个是前面提到的可以读取函数内部的变量，另一个就是让这些变量的值始终保持在内存中。
匿名函数：没有命名的函数 当成变量使用
*/

func intSeq() func() (int) {
	// 定义一个intSeq函数 返回值为一个匿名函数 匿名函数返回值为int类型
	// 返回给外界的func
	i := 1
	return func() int {
		i += 1
		return i
	}
}
func main() {
	gI := intSeq() // 返回一个含有变量i 环境的函数

	// 保存上下文环境
	fmt.Println("gI1:", gI()) //2
	fmt.Println("gI2:", gI()) //3
	fmt.Println("gI3:", gI()) //4

	// 新的变量会有新的上下文环境
	newgI := intSeq()
	fmt.Println("newgI:", newgI()) //2
}
