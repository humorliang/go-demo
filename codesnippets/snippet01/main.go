package main

import (
	"log"
	"os"
	"fmt"
)

/*
Convert Interface to Type: Type Assertion
接口类型转换：断言
*/
func main() {
	res := resultDatas()
	//类型断言  返回值  和 bool值
	v, ok := res.(int)
	if !ok {
		log.Printf("get data of type %T don't wanted int", res)
		os.Exit(1)
	}
	fmt.Println(v)
}

func resultDatas() interface{} {
	return "i"
}
