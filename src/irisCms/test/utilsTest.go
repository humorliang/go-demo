package main

import "irisCms/utils"

func main() {
	utils.CreateAppendWriteFile("/Users/datatist001/gopath/data/","test.txt",[]byte("hello world\n"))
}
