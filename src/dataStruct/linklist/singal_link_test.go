package linklist

import (
	"testing"
	"fmt"
)

func TestLinkList(t *testing.T) {
	//定义一个链表
	list := NewList()

	//添加节点
	list.Append(&Node{Data: 1})
	list.Append(&Node{Data: 2})
	list.Append(&Node{Data: 3})
	list.Append(&Node{Data: 4})
	list.Append(&Node{Data: 5})

	fmt.Println("list number:", list.Size)

	err := list.Remove(2)
	fmt.Println("remove list node error:", err)
	fmt.Println("list number:", list.Size)
	err = list.Remove(2)
	fmt.Println("remove list node error:", err)
	fmt.Println("list number:", list.Size)



}
