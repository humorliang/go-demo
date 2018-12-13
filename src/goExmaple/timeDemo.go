package main

import (
	"fmt"
	"time"
)

/*
时间time模块：
*/
func main() {
	p := fmt.Println

	now := time.Now()
	p(now) //2018-11-29 10:43:37.792306 +0800 CST m=+0.000823377

	then := time.Date(2016, 1, 28, 12, 20, 10, 1000, time.UTC)
	p(then) //2016-01-28 12:20:10.000001 +0000 UTC

	p(then.Year())    //2016
	p(then.Month())   //January
	p(then.Second())  //10
	p(then.Weekday()) //Thursday

	p(then.Before(now)) //true
	p(then.After(now))  //false

	//时间间隔
	diff := now.Sub(then)
	fmt.Println(diff) //24854h34m50.913372s

	//add时间间隔
	p(then.Add(diff))
}
