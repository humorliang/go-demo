package main

import (
	"pattern/singleton/single"
	"fmt"
)

func main() {
	//创建实例
	s := single.New()
	s["this"] = "that"
	//创建实例
	s2 := single.New()
	fmt.Println("s:", &s, "s2:", &s2)
	fmt.Println("this is ", s2["this"])
}
