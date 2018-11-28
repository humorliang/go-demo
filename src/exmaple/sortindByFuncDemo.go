package main

import (
	"fmt"
	"sort"
)

/*
照字母的长度而不是首字母顺序对字符串排序。
实现了 sort.Interface 的 Len，Less和 Swap 方法，
这样我们就可以使用 sort 包的通用Sort 方法了，Len 和 Swap 通常在各个类型中都差不多，
Less 将控制实际的自定义排序逻辑

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
*/

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"apple", "banana", "peer"}
	fmt.Println("friut:", fruits)
	//调用排序接口
	sort.Sort(ByLength(fruits))
	fmt.Println("frite sort:", fruits)
}
