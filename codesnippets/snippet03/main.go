package main

import "log"

/*
Disable Log Output During Tests
测试时候不输出日志
*/
func main() {
	DoAndLogSomething()
}

func DoAndLogSomething() int {
	log.Println("this is log info ")
	return 10
}
