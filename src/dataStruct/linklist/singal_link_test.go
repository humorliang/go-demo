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

	//fmt.Println("list number:", list.Size)
	//err := list.Remove(2)
	//fmt.Println("remove list node error:", err)
	//fmt.Println("list number:", list.Size)
	//err = list.Remove(3)
	//fmt.Println("remove list node error:", err)
	//err = list.Remove(3)
	//fmt.Println("remove list node error:", err)
	//fmt.Println("list number:", list.Size)
	//bl := list.SearchData(2)
	//fmt.Println(" 2 is exsit:", bl)
	//bl = list.SearchData(21)
	//fmt.Println(" 21 is exsit:", bl)
	//err := list.Delete(4)
	//fmt.Println("Delete List error:", err)
	//err = list.Delete(-1)
	//fmt.Println("Delete List error:", err)
	//err = list.Delete(8)
	//fmt.Println("Delete List error:", err)
	list.Insert(2, "1.1")
	fmt.Println("list size:", list.Size)
	fmt.Println("list show:", list.Show())
}
