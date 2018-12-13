package main

import (
	"fmt"
	"time"
)

func main() {
	//选择语句
	i := 2
	fmt.Print("write", i, "as:")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	//最后执行default
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("this is weekend")
	default:
		fmt.Println("this is weekday")
	}

	//没有选择条件时类似于 if/else
	t := time.Now() //返回时间
	fmt.Println(t)  //2018-11-21 11:18:01.919006 +0800 CST m=+0.000674814
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}

}
