package main

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"time"
)

/*
Handle Ctrl+C (Signal Interrupt) Close in the Terminal
添加一个终止信号的句柄
*/

const FILE_NAME = "go-test.txt"

func main() {

	//设置句柄
	SetUpCloseHandler()

	//创建文件
	CreateFile()

	for {
		fmt.Println("sleep")
		time.Sleep(4 * time.Second)
	}

}

//设置关闭句柄
func SetUpCloseHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	//开启一个go程去监听 信号
	go func() {
		<-c
		fmt.Println("\r Ctrl+C pressedin Terminal ")
		DeleteFiles()
		os.Exit(0)
	}()
}

//创建文件
func CreateFile() {
	fmt.Println("Create File!")
	file, _ := os.Create(FILE_NAME)
	defer file.Close()
}

//删除文件
func DeleteFiles() {
	fmt.Println("Delete File!")
	_ = os.Remove(FILE_NAME)
	fmt.Println("Done!")

}
