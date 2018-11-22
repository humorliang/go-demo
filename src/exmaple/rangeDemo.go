package main

import "fmt"

func main() {
	/*
	range 迭代各种数据结构 返回索引 和值
	*/
	// 迭代slice类型
	slice := []int{2, 3, 4, 5}
	sum := 0
	for index, value := range slice {
		fmt.Println("index: ", index)
		sum += value
	}
	fmt.Println("sum", sum)
	for _, values := range slice {
		//不需要索引可以使用 _ 进行丢弃
		fmt.Print(values, " ")
	}
	fmt.Println()
	fmt.Println("------map-----")
	// 迭代map类型
	kvs := map[string]string{"key1": "values1", "key2": "values2"}
	for k,v:=range kvs{
		fmt.Println(k+":"+v)
	}
	// 字符串
	for i,str:=range "golang"{
		fmt.Println(i,":",str,"-->",string(str)) //返回ascII值
	}
}
