package main

import (
	"fmt"
	"os"
)

/*
格式化字符串
*/
type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
	fmt.Printf("%T\n", p)
	fmt.Printf("%t\n", true)
	fmt.Printf("%d\n", 123)
	fmt.Printf("%b\n", 14)
	fmt.Printf("%c\n", 33)
	fmt.Printf("%x\n", 456)
	fmt.Printf("%f\n", 78.9)
	fmt.Printf("%p\n", &p)
	fmt.Printf("%s\n", "\"strings\"")
	fmt.Printf("%q\n", "\"strings\"")
	//Sprintf 则格式化并返回一个字符串而不带任何输出。
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)
	//Fprintf 来格式化并输出到 io.Writers而不是 os.Stdout
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
