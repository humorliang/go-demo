package main

import "fmt"

/*
在条件语句之前可以有一个语句；任何在这里声明的变量都可以在当前所有的条件分支中使用。
*/
func main() {
	//只有if 没有else
	if 8%2 == 0 {
		fmt.Println("8可以被2整除")
	}

	// if中声明变量
	if a := 2; 6/a == 0 {
		fmt.Println("6可以整除a")
	} else if 5/a != 0 {
		// 之前定义的a可以在这里面使用
		fmt.Println(a)
		fmt.Println("5不能被a整除")
	}
}
