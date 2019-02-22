package linklist

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

//删除指定节点
func (list *List) Remove(data interface{}) error {
	//取出链表的头节点
	pre := list.Head
	//判断
	if pre.Data == data {
		//将链表的头节点变为 头结点的下一个节点
		list.Head = pre.Next
	} else {
		//节点遍历  利用递归方式遍历
		for pre.Next != nil {
			//头节点的下一个节点  判断
			if pre.Next.Data == data {
				pre.Next = pre.Next.Next
			} else {
				pre = pre.Next
				if  {
					
				}
			}
		}
	}
	list.Size--
	return nil
}
