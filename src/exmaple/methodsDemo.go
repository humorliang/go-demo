package main

import (
	"math"
	"fmt"
)

/*
若把函数当成struct结构体的字段，函数带有接受者的函数，称为method(方法)

"A method is a function with an implicit first argument, called a receiver."
方法是带有隐式第一个参数的函数，称为接收者。

语法：func (r ReceiverType) funcName(parameters) (results)
*/

//定义两个struct类型
type reactAngle struct {
	with, height float64
}
type circle struct {
	radius float64
}

// 根据接受者的类型来判断函数
func (r reactAngle) area() float64 {
	return r.height * r.with
}

func (c circle) area() float64 {
	return 2 * math.Pi * c.radius * c.radius
}

// 当接受者为指针类型时可以修改接受者
func (c *circle) setRadius() float64 {
	c.radius = 10
	return 0
}

func main() {
	// 其他自定义类型的定义
	// 语法：type typeName typeLiteral
	type age int
	type money float32
	type months map[string]int

	m := months{"january": 1, "february": 2, "december": 10}
	r := reactAngle{10, 10}
	c := circle{2}

	fmt.Println("reactAngle area:", r.area())
	fmt.Println("circle area:", c.area())
	ptr := &c
	fmt.Println("circle radius: ", ptr.setRadius())
	fmt.Println(m)
	fmt.Println("circle radius: ", ptr.radius)
}
