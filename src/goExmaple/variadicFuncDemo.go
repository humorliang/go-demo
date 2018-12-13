package main

import "fmt"

func variableFunc(num ...int) (int) {
	// 不定数量参数
	sum := 0
	for _, v := range num {
		sum = sum + v
	}
	return sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	//注意传递不定参数时需要加 ...
	totle := variableFunc(nums...)
	fmt.Println(totle)
}
