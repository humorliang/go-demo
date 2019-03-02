package main

import (
	"fmt"
	"math"
)

/*
Go 支持字符、字符串、布尔和数值 常量 。

const 关键字进行定义

*/

func main() {
	// 定义变量
	const s string = "const"

	const n = 500000

	const d = 3e20 / n
	const e int = 3          // 常量定义不使用不会报错
	fmt.Println(s)           // const
	fmt.Println(int64(d))    //600000000000000
	fmt.Println(math.Sin(n)) //0.17783120151825887
}
