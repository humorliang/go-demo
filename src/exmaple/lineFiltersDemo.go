package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

/*
一个行过滤器 在读取标准输入流的输入，处理该输入，然后将得到一些的结果输出到标准输出的程序中是常见的一个功能
对 os.Stdin 使用一个带缓冲的 scanner，让我们可以直接使用方便的 Scan 方法来直接读取一行，
每次调用该方法可以让 scanner 读取下一行。
*/

func main() {
	//标准输入
	scanner := bufio.NewScanner(os.Stdin)

	//标准输出
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "error", err)
		os.Exit(1)
	}
}
