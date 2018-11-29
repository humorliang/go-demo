package main

import (
	"time"
	"fmt"
)

/*
时间戳：
一般程序会有获取 Unix 时间的秒数，毫秒数，或者微秒数的需要。
*/
func main() {
	now := time.Now()
	secd := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)
	fmt.Println(secd)
	fmt.Println(nanos)
}
