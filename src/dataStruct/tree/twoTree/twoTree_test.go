package twoTree

import (
	"testing"
	"fmt"
)

func TestTwoTree(t *testing.T) {
	root := Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Right = Create(2)
	root.Left.Right.Left = Create(4)
	root.Right.Left = Create(6)

	fmt.Printf("前序遍历：")
	root.PreTraver()
	fmt.Printf("\n中序遍历：")
	root.MiddleTraver()
	fmt.Printf("\n后序遍历：")
	root.AfterTraver()
}
