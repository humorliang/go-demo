package main

import "fmt"

/*
格式：
参数类型  输出类型(无返回值可省略)
func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
	//这里是处理逻辑代码
	//返回多个值
	return value1, value2
}
*/
//函数定义
func add(x, y int) (string, int) {
	/*
	参数类型定义   返回值类型定义
	x,y：形参
	return 返回多值
	*/
	return "total：", x + y
}

//变参函数
func changeArgs(args ...int) {
	//函数接受不定数量 int类型参数  arg为slice类型变量
	for _, v := range args {
		fmt.Println("args:", v)
	}
}

// 传值传指针
func sendArgsIsValue(x int) {
	//参数被函数调用时实际传递的是 值得copy 无法修改值
	fmt.Println(x)
	x += 1
}
func sendArgsIsPointer(x *int) {
	//传递指针类型时 实际传递的是一个地址  可以进行修改
	fmt.Println("x:", x)
	fmt.Println("*x:", *x)
	*x = *x + 1
}

func main() {
	//实参
	fmt.Println(add(1, 2))

	x := 10
	sendArgsIsValue(x)
	fmt.Println("传值：", x)
	fmt.Println("-----------")
	// 这里的实参需要传递指针即变量地址
	sendArgsIsPointer(&x)
	fmt.Println("传指针：", x)
}
