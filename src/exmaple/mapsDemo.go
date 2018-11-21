package main

import "fmt"

func main() {
	/*
	关联数组：map类型
	使用典型的 make[key] = val 语法来设置键值对。
	*/
	//定义map: make(map[key-type]val-type)
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println("map", m) //map map[one:1 two:2]
	fmt.Println("len", len(m))

	// map 添加
	m2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(m2)

	//访问
	vOne := m["one"]
	fmt.Println(vOne)
	_, prs := m["one2"]           // map可以返回两个值  第一个为键值(若存在) 第二个为bool 键是否存在
	fmt.Println("key four:", prs) //false

	//删除 delete(map,key)
	delete(m, "one") //
	fmt.Println("delete", m)
}
