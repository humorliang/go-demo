package main

import "fmt"

func main() {
	//数组定义：var arr [n]type

	// 第一种
	var a [5]int      // 定义一个a 数组
	a[0] = 10         //添加一个元素
	fmt.Println(a[3]) //数组访问  0  未初始化的变量默认值为0

	// 第二种
	b := [5]int{1, 2, 3, 4, 5} // 声明一个长度为5的数组
	fmt.Println(b[0])

	// 第三种
	c := [...]int{2, 3, 5, 6} // '...'自动根据元素个数计算长度
	fmt.Println(c[0])

	// 二维数组
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	fmt.Println(doubleArray) //[[1 2 3 4] [5 6 7 8]]

	a[2] = 100
	fmt.Println("set", a)
	fmt.Println("get", a[2])
	fmt.Println("len", len(a))

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD) //[[0 1 2] [1 2 3]]

}
