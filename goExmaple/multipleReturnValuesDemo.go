package main

import "fmt"

func reVals(x int, y int) (int, int) {
	//多值返回函数
	return x, y
}
func main() {
	a, b := reVals(2, 3)
	fmt.Println("a:", a, "b:", b)
}
