package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	//comm.Log()
	//log.Print("test")
	//path := "/data/log/test.log"
	////path := "test.log"
	//pathSlice := strings.Split(path, "/")
	//pathStr:=strings.Join(pathSlice[:len(pathSlice)-1],"/")
	//fmt.Println(pathStr)
	//fmt.Println(pathSlice[len(pathSlice)-1])

	//utils.OpenFile("../data/log/test.log")
	//utils.OpenFile("test2.log")
	//fmt.Println(time.Now().String())
	fmt.Println(strings.Split(time.Now().String()," ")[0])
}
