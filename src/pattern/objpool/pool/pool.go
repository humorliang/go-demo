package pool

import "fmt"

/*
对象池：创建一个简单的对象池模式
*/
//定义一个对象
type DbCon struct {
	name string
}

func (db *DbCon) index(index int) {
	fmt.Println("这是对象：", index)
}

//定义一个对象通道（对象池）
type Pool chan *DbCon

//对象池函数
func New(total int) *Pool {
	//初始化对象池
	p := make(Pool, total)
	//往对象池传递可用对象
	for i := 0; i < total; i++ {
		p <- new(DbCon)
	}
	//返回对象池
	return &p
}
