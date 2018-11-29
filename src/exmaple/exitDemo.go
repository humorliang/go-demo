package main

import (
	"fmt"
	"os"
)

/*
退出
当使用 os.Exit 时 defer 将不会 执行，所以这里的
*/
func main() {
	defer fmt.Println("!")

	os.Exit(3)
}
