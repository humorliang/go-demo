package main

import "fmt"

/*
Anonymous Functions (aka Closures)
匿名函数：（闭包）
*/
func main() {

	//匿名函数 闭包
	msg := "this is msg"
	func(m string) {
		fmt.Println(m)
	}(msg)
}
