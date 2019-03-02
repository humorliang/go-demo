package main

import "fmt"

/*
1. var 声明一个或多个变量
2. 变量没有初始值，会默认为0
3. Go 自动推断初始化变量类型
4. := 初始化变量简写


格式：

//定义一个名称为“variableName”，类型为"type"的变量
var variableName type

//定义三个类型都是“type”的变量
var vname1, vname2, vname3 type

//初始化“variableName”的变量为“value”值，类型是“type”
var variableName type = value


//	定义三个类型都是"type"的变量,并且分别初始化为相应的值 vname1为v1，vname2为v2，vname3为v3
var vname1, vname2, vname3 type= v1, v2, v3


// 定义三个变量，它们分别初始化为相应的值vname1为v1，vname2为v2，vname3为v3 然后Go会根据其相应值的类型来帮你初始化它们
var vname1, vname2, vname3 = v1, v2, v3


//	定义三个变量，它们分别初始化为相应的值vname1为v1，vname2为v2，vname3为v3 编译器会根据初始化的值自动推导出相应的类型
vname1, vname2, vname3 := v1, v2, v3



// _（下划线）是个特殊的变量名，任何赋予他的值都被抛弃
_, b := 34, 35

** 命名未使用的变量，编译时会报错


*/
func main() {
	var a string = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short var"
	fmt.Println(f)

	_, g := 2, 3
	fmt.Println(g)
}
