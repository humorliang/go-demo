package main

import (
	"os"
	"fmt"
)

/*
Defer 被用来确保一个函数调用在程序执行结束前执行。
同样用来执行一些清理工作。 defer 用在像其他语言中的ensure 和 finally用到的地方。

defer在函数结束前进行调用的函数
*/

//创建文件 文件路径 返回文件指针
//os.File代表一个打开的文件对象。
func createFile(p string) *os.File {
	fmt.Println("create file")
	f, err := os.Create(p)
	if err != nil {
		fmt.Println(err)
	}
	return f
}

//写入文件
/*
	Fprintln采用默认格式将其参数格式化并写入w。
	总是会在相邻参数的输出之间添加空格并在输出结束后添加换行符。
	返回写入的字节数和遇到的任何错误。
*/
func writeFile(f *os.File) {
	fmt.Println("write file")
	fmt.Fprint(f, "this is data")
}

//关闭文件
func closeFile(f *os.File) {
	fmt.Println("close file")
	f.Close()
}

func main() {
	f := createFile("./tests.txt")
	//延迟函数在函数要结束关闭前
	defer closeFile(f)
	writeFile(f)
}
