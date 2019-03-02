package main

import (
	"pattern/memento/mem"
	"container/list"
	"fmt"
)

func main() {
	//创建备忘录管理器 创建一个新的链表
	storage := &mem.Storage{list.New()}
	//创建文本
	text := &mem.Text{Value: "hello world"}
	storage.PushBack(text.SaveToMemento())
	//写入对象
	text.Write("this is first commit ")
	//创建备备份
	storage.PushBack(text.SaveToMemento())
	fmt.Println(text.Read())

	text.Write("this is second commit ")
	//创建备备份
	storage.PushBack(text.SaveToMemento())
	fmt.Println(text.Read())

	text.Write("this is three commit ")
	//创建备备份
	storage.PushBack(text.SaveToMemento())
	fmt.Println(text.Read())

	//进行回滚(调度者)
	mediator := storage.RightPop()
	if mediator != nil {
		//mediator.value获取当前对象的值  并进行类型断言
		text.RestoreFromMemento(mediator.Value.(*mem.Memento))
	}
	fmt.Println(text.Read())

	//进行回滚(调度者)
	mediator = storage.RightPop()
	if mediator != nil {
		//mediator.value获取当前对象的值  并进行类型断言
		text.RestoreFromMemento(mediator.Value.(*mem.Memento))
	}
	fmt.Println(text.Read())

	//进行回滚(调度者)
	mediator = storage.RightPop()
	if mediator != nil {
		//mediator.value获取当前对象的值  并进行类型断言
		text.RestoreFromMemento(mediator.Value.(*mem.Memento))
	}
	fmt.Println(text.Read())

	//进行回滚(调度者)
	mediator = storage.RightPop()
	if mediator != nil {
		//mediator.value获取当前对象的值  并进行类型断言
		text.RestoreFromMemento(mediator.Value.(*mem.Memento))
	}
	fmt.Println(text.Read())

}
