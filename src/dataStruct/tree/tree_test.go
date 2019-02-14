package tree

import (
	"testing"
	"fmt"
)

func TestTwoTree(t *testing.T) {
	root := Node{Value: 1}
	root.Left = &Node{ Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = Create(4)
	root.Left.Right = Create(5)
	root.Right.Left = Create(6)
	root.Right.Left.Right = Create(7)
	/*
			 1
		  2     3
		4  5  6
				7
	*/
	fmt.Printf("前序遍历：")
	root.PreTraver()
	fmt.Printf("\n中序遍历：")
	root.MiddleTraver()
	fmt.Printf("\n后序遍历：")
	root.AfterTraver()
	fmt.Printf("\n广度优先：")
	fmt.Println(BreadthFirst(&root))
	fmt.Println("\n节点深度：")
	fmt.Println(NodeDepth(&root))
	fmt.Println("\n节点层级：")
	NodeLevel(&root,)
}
