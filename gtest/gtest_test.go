package g

import (
	"testing"
	"fmt"
)

func TestDemo(t *testing.T) {
	a := [5]int{1, 2, 3, 4, 5}
	b := a[0:1]
	c := b[0:2]
	//fmt.Println("a:",a)
	fmt.Println("c:",c)
	c = append(c, 1)
	fmt.Println("a:",a)
	fmt.Println("b:",b)
	fmt.Println("c:",c)
	/*
	a: [1 1 3 4 5]
	b: [1]
	c: [1 1]
	*/
}
