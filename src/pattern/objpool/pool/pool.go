package pool

import (
	"fmt"
	"strconv"
)

/*
对象池：创建一个简单的对象池模式
对象池模式是一种创建型模式，根据需求来预测将要使用的对象，提前创建并保存在内存中。
*/
//定义一个对象
type Object struct {
	name string
}

func (obj *Object) Do(index int) {
	fmt.Println("object Do:", strconv.Itoa(index))
}

type Pool chan *Object

//对象池函数
func New(total int) *Pool {
	//初始化对象池
	p := make(chan *DbCon, total)
	//往对象池传递可用对象
	for i := 0; i < total; i++ {
		p <- new(Object)
	}
	//返回对象池
	return &p
}
