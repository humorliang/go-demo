package main

import (
	"time"
	"fmt"
)

/*
有时候会出现goroutine阻塞的情况，那么我们如何避免整个程序进入阻塞的情况呢？
我们可以利用select来设置超时，通过如下的方式实现：
case <-time.After(time.Second*1)
*/
func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()
	//case一个时间通道作为超时返回
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}
	//fmt.Println("c1:",<-c1) //c1: result 1

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}
