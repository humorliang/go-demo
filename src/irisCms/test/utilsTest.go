package main

import "irisCms/utils"

func main() {
	utils.CreateAppendWriteFile("/Users/datatist001/gopath/car/","tests.txt",[]byte("hello world\n"))
}
