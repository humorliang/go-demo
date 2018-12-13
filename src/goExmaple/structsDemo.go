package main

import "fmt"

/*
结构体：struct类型（属于自定义类型）
	我们可以声明新的类型，作为其它类型的属性或字段的 容器
	可变的
格式：
type person struct {
	name string
	age int
}
*/
func main() {
	//定义一个struct结构体 一个新的person类型
	type person struct {
		name string
		age  int
	}
	p := person{"张三", 23} //声明p变量为person类型并初始化
	p2 := new(person)     //定义p2为person类型指针
	fmt.Println("p->age:", p.age)
	fmt.Println("p2:", p2)   //p2: &{ 0}
	fmt.Println("&p2:", &p2) //0xc000082018
	pPtr := &p               //将p地址给pPtr
	fmt.Println("pPtr->age", pPtr.age)	//23   //结构体指针使用 . 指针会被 自动解引用
	pPtr.age = 40
	fmt.Println("pPtr->age", pPtr.age) //p->age 40
	fmt.Println("-----------------------")

	//struct 匿名字段（嵌入字段）
	type human struct {
		name string
		age  int
	}
	type student struct {
		human //匿名字段 默认包含所有的human字段
		speciality string
	}
	//定义一个stu变量
	stu := student{human{"张三", 23}, "computer speciality"}
	fmt.Println("stu->name:", stu.name)
	fmt.Println("stu->age:", stu.age)
	fmt.Println("stu->speciality:", stu.speciality)


}
