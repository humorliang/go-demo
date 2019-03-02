package main

import "fmt"

/*
递归：不停的重复，直到有条件结束这种重复
在某个条件之前一只调用本身
f(n)=f(n-1)+f(n-2)
f(1)=1
f(0)=0
*/
func fib(a int) (int) {
	// 1 1 2 3 5 8 13
	//终止条件
	if a == 0 {
		return 0
	}
	if a == 1 {
		return 1
	}
	// 进行递归运算
	return fib(a-1) + fib(a-2)
}
func main() {
	// 递归运算
	num := fib(5)
	fmt.Println(num)
}
