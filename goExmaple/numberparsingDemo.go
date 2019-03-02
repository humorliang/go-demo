package main

import (
	"fmt"
	"strconv"
)

/*
数字解析：
内置的 strconv 包提供了数字解析功能。
*/
func main() {
	//19.4 <nil>
	f, _ := strconv.ParseFloat("2.34", 64)
	fmt.Println(f)

	//自动推断字符的进制
	i, _ := strconv.ParseInt("23", 0, 64)
	fmt.Println(i)
}
