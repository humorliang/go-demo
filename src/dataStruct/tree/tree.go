package tree

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
	Layer       int
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

//节点层级
func NodeLevel(n *Node) {
	//利用队列进行实现
	tq := []*Node{}
	//结果容器
	//var res [][]int
	if n != nil {
		tq = append(tq, n)
	}
	for len(tq) > 0 {
		//取一个元素
		p := tq[0]
		//var r []int
		//r = append(r, p.Value)
		//res = append(res, r)
		fmt.Println(p.Value)
		//更新队列
		tq = tq[1:]
		//
		if n.Left != nil {
			tq = append(tq, n.Left)
		}
		if n.Right != nil {
			tq = append(tq, n.Right)
		}

	}
}

//广度优先
func BreadthFirst(n *Node) ([]interface{}) {
	//广度优先利用队列  

	//结果容器
	var result []interface{}
	//节点容器
	var nodes = []Node{*n}
	//遍历循环判断
	for len(nodes) > 0 {
		//取节点
		node := nodes[0]
		//更新节点容器
		nodes = nodes[1:]
		//插入节点值
		result = append(result, node.Value)

		//下一层节点判断并插入节点容器
		if node.Left != nil {
			nodes = append(nodes, *node.Left)
		}
		if node.Right != nil {
			nodes = append(nodes, *node.Right)
		}
	}
	return result
}

//二叉树的深度
func NodeDepth(n *Node) int {
	//深度计算
	depth := 0
	//当节点为空时结束递归
	if n != nil {
		if NodeDepth(n.Left) > NodeDepth(n.Right) {
			depth = NodeDepth(n.Left) + 1
		} else {
			depth = NodeDepth(n.Right) + 1
		}
	}
	return depth
}
