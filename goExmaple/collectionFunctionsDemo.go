package main

import "fmt"

/*
组合函数：
我们经常需要程序在数据集上执行操作，比如选择满足给定条件的所有项，
或者将所有的项通过一个自定义函数映射到一个新的集合上
*/

//查找元素字符索引
func index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

//查找是否包含元素
func include(vs []string, t string) bool {
	return index(vs, t) >= 0
}

//有任意一个元素满足返回true
func any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

//对数据进行过滤返回符合条件集合
func filter(vs []string, f func(string) bool) []string {
	//创建一个容器
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

//对数据进行遍历
func maps(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {
	var strs = []string{"banana", "apple", "plum", "pear"}
	fmt.Println(index(strs, "pear"))
}
