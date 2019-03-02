package linklist

import "fmt"

/*
单链表：
下面这副图是我们单链表运煤车队。
每节运煤车就是单链表里的元素，每节车厢里的煤炭就是元素中保存的数据。前后车通过锁链相连，
作为单链表运煤车，从1号车厢开始，每节车厢都知道后面拉着哪一节车厢，却不知道前面是哪节车厢拉的自己。
第一节车厢没有任何车厢拉它，我们就叫它车头，第五节车厢后面拉其他车厢，我们称为车尾。
作为单链表它最大的特点就是能随意增加车队的长度，也能随意减少车队的长度。这是比数组公交车最大的优点。

1. 节点
2. 链表
3. 方法
*/

//节点内容对象  存储内容
type Object interface{}

//节点
type Node struct {
	Data Object //储存内容
	Next *Node  //下一个节点
}

//链表
type List struct {
	Size int
	Head *Node
	Tail *Node
}

//链表方法
func NewList() *List {
	return &List{
		Size: 0,
		Head: nil,
		Tail: nil,
	}
}

//尾部添加节点
func (list *List) Append(n *Node) {
	//空链表直接
	if list.Size == 0 {
		list.Head = n
		list.Tail = n
	} else {
		//取出当前链表的 尾节点
		lastList := list.Tail
		//将尾节点的下一个节点变为添加的节点
		lastList.Next = n
		//将链表的尾节点变为 添加的节点
		list.Tail = n
	}
	list.Size++
}

//头部添加节点
func (list *List) Add(n *Node) {
	//当前节点的下一个节点关联
	n.Next = list.Head
	//更新链表
	list.Head = n
	list.Size++
}

//删除指定节点
func (list *List) Remove(data interface{}) error {
	//取出链表的头节点
	pre := list.Head
	lsize := list.Size
	//判断
	if pre.Data == data {
		//将链表的头节点变为 头结点的下一个节点
		list.Head = pre.Next
		list.Size--
	} else {
		//节点遍历  利用递归方式遍历
		for pre.Next != nil {
			//头节点的下一个节点  判断
			if pre.Next.Data == data {
				pre.Next = pre.Next.Next
				list.Size--
			} else {
				pre = pre.Next
			}
		}
		if list.Size == lsize {
			return fmt.Errorf("%v", "Data is not Exsit")
		}
	}
	return nil
}

//链表的遍历
func (list *List) SearchData(data interface{}) bool {
	//获取链表的头结点
	cur := list.Head
	//遍历链表
	for cur != nil {
		if cur.Data == data {
			return true
		} else {
			//将下一个进行遍历
			cur = cur.Next
		}
	}
	return false
}

//删除指定位置节点
func (list *List) Delete(index int) error {
	//取头节点
	cur := list.Head
	if index <= 0 {
		cur = cur.Next
	} else if list.Size <= index {
		return fmt.Errorf("%v", "index is out range")
	} else {
		count := 0 //遍历位置统计  当前首节点
		//while循环 遍历
		for count != (index-1) && cur.Next != nil {
			count++
			//当前节点转化
			cur = cur.Next
		}
		//当前节点的下一个节点
		//count代表的当前节点
		cur.Next = cur.Next.Next
	}
	return nil
}

//指定位置插入节点
func (list *List) Insert(index int, data interface{}) {
	//获取首节点
	cur := list.Head
	//索引为负数时默认从头部添加
	if index < 0 {
		list.Add(&Node{Data: data})
	} else if list.Size < index {
		list.Append(&Node{Data: data})
	} else {
		//位置统计
		count := 0
		//进行位置统计
		for count < (index - 1) {
			cur = cur.Next
			count++
		}
		//循环退出cur在index-1位置  取代的是cur.next位置
		node := &Node{Data: data}
		//取代位置
		node.Next = cur.Next
		//更新位置
		cur.Next = node
		list.Size++
	}
}

//判断链表是否为空
func (list *List) IsEmpty() bool {
	if list.Size == 0 {
		return true
	}
	return false
}

//显示链表
func (list *List) Show() []interface{} {
	if list.Size == 0 {
		return nil
	} else {
		var res []interface{}
		cur := list.Head
		res = append(res, cur.Data)
		for cur.Next != nil {
			//先添加在遍历
			res = append(res, cur.Next.Data)
			cur = cur.Next
		}
		return res
	}
}



