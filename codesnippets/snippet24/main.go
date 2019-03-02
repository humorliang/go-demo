package main

import (
	"os"
	"net/http"
	"log"
	"fmt"
	"strings"
)

/*
How Detect Content Type of a File
判断文件类型
*/

func main() {
	//打开文件
	f, err := os.Open("test.pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//获取内容
	contentType, err := GetFileContentType(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File content Type:", contentType)
}

//传入文件句柄，返回文件类型
func GetFileContentType(out *os.File) (string, error) {

	//读取前 512byte 去辨别内容类型
	buffer := make([]byte, 512)

	//读取文件内容到buffer
	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	//利用http的DetectContentType函数判断文件类型
	contentType := http.DetectContentType(buffer)
	strList := strings.Split(contentType, "/")
	return strList[1], nil
}
