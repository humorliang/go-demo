package main

import "irisCms/utils"

func main() {
	utils.CreateAppendWriteFile("/Users/datatist001/gopath/data/","tests.txt",[]byte("hello world\n"))
}
