package main

import (
	"log"
	"errors"
)

/*
Creating Your Own Custom Errors
创建自定义错误
*/
func main() {
	//运行错误
	if err := DoSomeThing(); err != nil {
		log.Fatal(err)
	}
}

func DoSomeThing() error {
	return errors.New("this is define error , doSomething has error")
}
