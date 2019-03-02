package main

import (
	"time"
	"log"
)

/*
How Long Does a Function Take: Measuring Execution Time
做一个函数的执行时间的量化统计
*/
func main() {
	LongRunningFunction()
}

func LongRunningFunction() {
	defer TimeTaken(time.Now(), "LongRunningFunction")
	time.Sleep(3 * time.Second)
}

func TimeTaken(t time.Time, name string) {
	//间隔时间
	elapsed := time.Since(t)
	log.Printf("TIME:%s took %s \n", name, elapsed)
}
