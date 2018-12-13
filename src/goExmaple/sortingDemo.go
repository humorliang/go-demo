package main

import (
	"sort"
	"fmt"
)

/*
排序方法是正对内置数据类型的；这里是一个字符串的例子。注意排序是原地更新的，所以他会改变给定的序列并且不返回一个新值。
*/
func main() {
	//定义数据
	strs := []string{"a", "c", "b", "e"}
	ints := []int{2, 3, 1, 4}
	//fmt.Printf("类型为：%T 类型\n", str)
	fmt.Println("str:", strs)
	fmt.Println("ints:", ints)
	//排序
	sort.Strings(strs)
	sort.Ints(ints)
	fmt.Println("str排序后：", strs)
	fmt.Println("ints排序后：", ints)
	//查看序列是否排序
	fmt.Println("ints是否排序：", sort.IntsAreSorted(ints)) //true
}
