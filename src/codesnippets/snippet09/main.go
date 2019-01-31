package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*
Sorting an Array of Numbered String Values
排序一个数字字符串
*/
func main() {
	//待排序数字字符切片
	myList := []string{"1", "10", "11", "21", "13", "4", "15", "6", "17", "80", "9"}

	fmt.Printf("Before:%v\n", myList)

	sort.Slice(myList, func(i, j int) bool {
		//这里的i,j 相当于 i和i+1 是位置索引
		a, _ := strconv.Atoi(myList[i])
		b, _ := strconv.Atoi(myList[j])
		//返回bool值，确定是否交换位置
		return a < b
	})
	fmt.Printf("Sort after:%v\n", myList)
}
