package main

import (
	"strconv"
	"fmt"
)

/*
Convert uint64 to a String
uint64整形转换成字符串
*/
func main() {
	var myNum uint64
	myNum = 18446744073709551615

	//格式转换
	//str0:=strconv.Itoa()参数为int 不能是uint64
	str := strconv.FormatUint(myNum, 10)

	fmt.Println("my number:" + str)
}
