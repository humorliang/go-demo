package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
随机数：
Go 的 math/rand 包提供了伪随机数生成器（英）
*/
func main() {
	//Intn() => [0,n)
	fmt.Println("random int:",rand.Intn(100))
	//Float64 a pseudo-random number in [0.0,1.0)
	fmt.Println("random float64:",rand.Float64())

	fmt.Println(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100))
}
