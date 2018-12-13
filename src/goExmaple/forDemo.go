package main

import "fmt"

func main() {
	// 第一种for格式: 带个单个循环条件
	i := 1
	for i < 3 {
		fmt.Println(i)
		i += 1
	}

	// 第二种：初始化/条件/表达式
	for j := 5; j < 8; j++ {
		fmt.Println(j)
	}

	// 第三种：无条件  等待break ,return 结束
	for {
		fmt.Println("loop")
		break
	}
}
