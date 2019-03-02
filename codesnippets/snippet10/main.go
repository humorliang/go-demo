package main

import (
	"io/ioutil"
	"os"
	"log"
	"fmt"
)

/*
Creating & Writing to Temp Files
创建并且写入一个临时文件
*/
func main() {
	//使用ioutil.TempDir()接收一个文件夹，和前缀
	tmpFile, err := ioutil.TempFile(os.TempDir(), "prefixTest-")
	if err != nil {
		log.Fatal("don't create temp file", err)
	}
	//文件使用完进行清理
	defer os.Remove(tmpFile.Name())

	fmt.Println("Created File:", tmpFile.Name())
	//文件写入内容
	data := []byte("this is tmp data ")
	if _, err := tmpFile.Write(data); err != nil {
		log.Fatal("tmp file write data err:", err)
	}
	if err := tmpFile.Close(); err != nil {
		log.Fatal("close file fail:", err)
	}
}
