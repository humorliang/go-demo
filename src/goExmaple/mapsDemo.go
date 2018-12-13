package main

import "fmt"

func main() {
	/*
	关联数组：map类型 类似python 字典 map[keyType]valueType
	使用典型的 make[key] = val 语法来设置键值对。

	1. map是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必须通过key获取
	2. map的长度是不固定的，也就是和slice一样，也是一种引用类型
	3. 内置的len函数同样适用于map，返回map拥有的key的数量
	4. map的值可以很方便的修改，通过numbers["one"]=11可以很容易的把key为one的字典值改为11
	5. map和其他基本型别不同，它不是thread-safe，在多个go-routine存取时，必须使用mutex lock机制
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
