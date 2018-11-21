package main

import "fmt"

func main() {
	fmt.Println("hello World")
}

// go run helloworld.go  运行

// go build hellworld.go  编译成二进制文件

// ./helloworld
/*
Go程序是通过package来组织的，main包是一个独立可执行的包，编译后可生成可执行文件
Go支持utf-8编码

package <pkgName>

Go使用package（和Python的模块类似）来组织代码。main.main()函数(这个函数位于主包）是每一个独立的可运行程序的入口点。
Go使用UTF-8字符串和标识符(因为UTF-8的发明者也就是Go的发明者之一)，所以它天生支持多语言。

*/
