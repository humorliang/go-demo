package main

import (
	"os"
	"fmt"
)

/*
Check If a File Exists Before Using It
检查文件是否存在
*/

func main() {
	filePath := "test/"

	//返回文件的描述信息
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Printf("file or directory is not exists:%s \n", err)
	} else {
		fmt.Println("file or directory is exists")
	}
}
