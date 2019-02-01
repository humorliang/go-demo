package main

import (
	"fmt"
	"strings"
	"github.com/dustin/go-humanize"
	"os"
	"net/http"
	"io"
	"log"
)

/*
Download Large Files with Progress Reports
下载大文件的进程管理
*/
func main() {
	fmt.Println("Download Start")
	fileUrl := "https://upload.wikimedia.org/wikipedia/commons/d/d6/Wp-w4-big.jpg"
	err := DownloadFile("map.jpg", fileUrl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Download Finished")
}

//定义一个写入对象的统计类型
type WriteCounter struct {
	Total uint64
}

//定义输出下载状态
func (wc WriteCounter) OutProgress() {
	//输出35个空格进行分割
	fmt.Printf("\r%s", strings.Repeat(" ", 40))

	//输出当前的下载状态  humanize将字节转为MB
	fmt.Printf("\r Downloading...%s complete", humanize.Bytes(wc.Total))
}

//数据写入大小 实现writer接口的Write()方法
func (wc *WriteCounter) Write(p []byte) (int, error) {
	//写入大小
	n := len(p)
	wc.Total += uint64(n)
	wc.OutProgress()
	return n, nil
}

//下载文件操作
func DownloadFile(saveFilePath string, url string) error {
	//创建下载临时文件，
	out, err := os.Create(saveFilePath + ".tmp")
	if err != nil {
		return err
	}
	defer out.Close()

	//获取数据
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	fmt.Println(resp.ContentLength)
	defer resp.Body.Close()

	//在写入的过程中写入信息
	counter := &WriteCounter{}
	// TeeReader(r,w)TeeReader返回一个将其从r读取的数据写入w的Reader接口。
	// 所有通过该接口对r的读取都会执行对应的对w的写入。
	// 没有内部的缓冲：写入必须在读取完成前完成。写入时遇到的任何错误都会作为读取错误返回。
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		return err
	}
	fmt.Print("\n")

	//将临时文件重命名
	err = os.Rename(saveFilePath+".tmp", saveFilePath)
	if err != nil {
		return err
	}
	return nil
}
