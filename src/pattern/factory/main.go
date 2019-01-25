package main

import (
	"pattern/factory/car"
	"fmt"
)

/*
接口理解：
1. 当一个结构体实现了这个接口,我们可以通过接口变量来保存这个结构体变量
2. 但是接口变量不能调用接口中没有的方法,也不能调用结构体的属性,
要想调用接口以外的方法和变量,我们只能通过接口类型转换来实现
通过value,ok := 接口变量名称.(具体数据类型)

*/
func main() {

	//定义一个车辆工厂
	var fty car.CarFactory
	//指定一个车辆工厂
	//一个类型A只要实现了接口X所定义的全部方法，那么A类型的变量也是X类型的变量。
	fty = new(car.ShFactory) //对接口进行赋值
	//造一个大车
	c := fty.CreateBigCar()
	fmt.Println(c.Info())
	fty = new(car.SzFactory)
	c = fty.CreateBigCar()
	fmt.Println(c.Info())
}
