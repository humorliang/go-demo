package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"bufio"
)

/*
读文件:
io 包提供了一些可以帮助我们进行文件读取的函数。
bufio 包实现了带缓冲的读取，
这不仅对有很多小的读取操作的 能提升性能，也提供了很多 附加的读取函数。
*/
//错误检查
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	//读取文件到内存
	dat, err := ioutil.ReadFile("./test.txt")
	check(err)
	fmt.Println(string(dat))

	//打开文件
	f, err := os.Open("./test.txt")
	check(err)
	//fmt.Println(f)
	//读取大小
	b1 := make([]byte, 5)
	//将内容读取到b1中返回读取大小
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	//设置读取位置
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	//带有缓冲的读取
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	//关闭文件
	f.Close()
}
