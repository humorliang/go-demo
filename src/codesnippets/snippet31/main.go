package main

/*
Generating a Random Number
获取一个随机数
*/
import (
	"math/rand"
	"time"
	"fmt"
)

//获取区间随机数
func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	// Seed, unlike the Rand.Seed method, is safe for concurrent use.
	// 设置当前随机数产生的方法
	rand.Seed(time.Now().UnixNano())
	randomNum := random(1000, 2000)
	fmt.Println("random num:", randomNum)
}
