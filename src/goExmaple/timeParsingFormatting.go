package main

import (
	"fmt"
	"time"
)

/*
时间格式化：
Go 支持通过基于描述模板的时间格式化和解析。
Format 和 Parse 使用基于例子的形式来决定日期格式，
一般你只要使用 time 包中提供的模式常量就行了，但是你也可以实现自定义模式。
*/
func main() {
	p := fmt.Println
	t := time.Now()

	p(t.Format(time.RFC3339))
	t1, e := time.Parse(time.RFC3339, "2015-10-10T22:10:30+00:00")
	fmt.Println(t1, e)

	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d +0000 +0000\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}
