package main

import "fmt"

/*
切片：slice类型
slice并不是真正意义上的动态数组，而是一个引用类型。slice总是指向一个底层array，slice的声明也可以像array一样，只是不需要长度
*/
func main() {
	//slice 定义 和数组相同但是没有长度
	var sl = []int{11, 21, 31, 14, 10}
	fmt.Println(sl) //[1 2 4]

	//简短声明
	sl2 := []byte{'a', 'b', 'c'}
	fmt.Println(sl2)

	//使用make([]string,length)定义
	sl3 := make([]string, 4)
	fmt.Println(sl3)

	sl2[0] = 10
	sl2[1] = 20
	fmt.Println("set", sl2)
	fmt.Println("get", sl2[0])
	fmt.Println("len", len(sl2))
	sl2 = append(sl2, 30) //返回包含一个或多个ele的slice
	fmt.Println("append", sl2)

	//复制  创建一个同等长度的slice  进行复制
	c := make([]int, len(sl))
	copy(c, sl) //将sl 复制给 c
	fmt.Println("copy", c)

	//切片操作
	l := sl[1:3]
	fmt.Println("sl1", l)
	l = sl[1:]
	fmt.Println("sl2", l)
	l = sl[:3]
	fmt.Println("sl3", l)

	//多维数据结构
	twoDic := make([][]int, 3)
	for i := 0; i < 3; i++ {
		lenStr := i + 1
		twoDic[i] = make([]int, lenStr)
		for j := 0; j < lenStr; j++ {
			twoDic[i][j] = i + j
		}
	}
	fmt.Println("twoDic", twoDic) //twoDic [[0] [1 2] [2 3 4]]
}
