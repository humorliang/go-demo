package main

import (
	"pattern/observer/obser"
	"time"
	"strconv"
)

func main() {
	//定义发布者
	sub := &obser.Subject{}
	//定义观察者
	aObs := &obser.ObserverObj{"a"}
	bObs := &obser.ObserverObj{"b"}

	//发布者添加观察成员
	sub.Attach(aObs, bObs)

	//发布消息
	for i := 0; i < 6; i++ {
		time.Sleep(1 * time.Second)
		sub.State("this is " + strconv.Itoa(i) + " views")
	}
}
