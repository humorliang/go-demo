package main

import "fmt"

/*
for 和 range为基本的数据结构提供了迭代的功能。我们也可以使用这个语法来遍历从通道中取得的值。
*/

func main() {
	queue := make(chan string, 2) //这里要设置通道缓存
	queue <- "one"
	queue <- "two"
	close(queue) //要结束通道否则 range会一直等待接受通道值
	for v := range queue {
		fmt.Println(v)
	}
}
