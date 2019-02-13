package twoTree

import (
	"fmt"
)

/*
二叉树：
二叉树是每个结点最多有两个子树的树结构。
通常子树被称作“左子树”（left subtree）和“右子树”（right subtree）。二叉树常被用于实现二叉查找树和二叉堆。
*/

//节点
type Node struct {
	Value       int
	Left, Right *Node //节点指针类型
}

//获取节点值
func (n *Node) GetValue() {
	fmt.Printf("%v ", n.Value)
}

//设置节点值
func (n *Node) SetValue(value int) {
	if n == nil {
		fmt.Println("node is exits ")
		return
	}
	n.Value = value
}

//创建节点
func Create(value int) *Node {
	return &Node{
		Value: value,
	}
}

/*
1、前序遍历（与树的前序遍历一样）
基本思想：先访问根结点，再先序遍历左子树，最后再先序遍历右子树即根—左—右。

2、中序遍历
基本思想：先中序遍历左子树，然后再访问根结点，最后再中序遍历右子树即左—根—右。

3、后序遍历
基本思想：先后序遍历左子树，然后再后序遍历右子树，最后再访问根结点即左—右—根。
*/

//前序遍历
func (n *Node) PreTraver() {
	if n == nil {
		return
	}
	n.GetValue()
	n.Left.PreTraver()
	n.Right.PreTraver()
}

//中序遍历
func (n *Node) MiddleTraver() {
	if n == nil {
		return
	}
	n.Left.MiddleTraver()
	n.GetValue()
	n.Right.MiddleTraver()
}

//后序遍历
func (n *Node) AfterTraver() {
	if n == nil {
		return
	}
	n.Left.AfterTraver()
	n.Right.AfterTraver()
	n.GetValue()
}

//广度优先(层级遍历)
func (n *Node) BreadthTraver() {
	if n == nil {
		return
	}

}
