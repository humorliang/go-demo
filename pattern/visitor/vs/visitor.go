package vs

import (
	"fmt"
	"math"
)

/*
 访问者模式（Visitor Pattern） 主要将数据结构与数据操作分离

 	1、对象结构中对象对应的类很少改变，但经常需要在此对象结构上定义新的操作。
	2、需要对一个对象结构中的对象进行很多不同的并且不相关的操作，
		而需要避免让这些操作"污染"这些对象的类，也不希望在增加新操作时修改这些类。

角色组成：
	1. 抽象访问者

	2. 访问者

	3. 抽象元素类

	4. 元素类

	5. 结构容器: (非必须) 保存元素列表，可以放置访问者
*/
//抽象访问者
type IVisitor interface {
	Visit(DataStruct)
}

/*抽象元素类*/
type DataStruct interface {
	Accept(IVisitor)
}

//元素类
type Data struct {
	A int
	B int
}

//元素类接受一个访问者
func (d *Data) Accept(vi IVisitor) {
	vi.Visit(d)
}

//定义加法访问者
type AddVisitor struct {
}

func (self *AddVisitor) Visit(ds DataStruct) {
	//数据类型断言
	data := ds.(*Data)
	sum := data.A + data.B
	fmt.Println("A + B =", sum)
}

//定义减法访问者
type SubVisitor struct {
}

func (self *SubVisitor) Visit(ds DataStruct) {
	//数据类型断言
	data := ds.(*Data)
	sub := data.A - data.B
	fmt.Println("abs(A-B)=", math.Abs(float64(sub)))
}
