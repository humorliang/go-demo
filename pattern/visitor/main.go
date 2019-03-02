package main

import "pattern/visitor/vs"

func main() {
	//定义数据
	data := &vs.Data{A: 10, B: 22}

	//定义访问者
	add := &vs.AddVisitor{}
	sub := &vs.SubVisitor{}

	//数据添加访问者
	data.Accept(add)
	data.Accept(sub)
}
