package main

import "os"

/*
panic 意味着有些出乎意料的错误发生。
通常我们用它来表示程序正常运行中不应该出现的，或者我们没有处理好的错误。
运行程序将会引起 panic，输出一个错误消息和 Go 运行时栈信息，并且返回一个非零的状态码。
Process finished with exit code 2
*/

func main() {
	//panic("a problem")

	//
	_, err := os.Create("./test.txt")
	if err != nil {
		panic(err)
	}
}
