package main

import (
	"archive/zip"
	"path/filepath"
	"strings"
	"os"
	"fmt"
	"io"
	"log"
	"time"
)

/*
Unzip Files in Go
解压文件
*/
func main() {
	start := time.Now()
	files, err := unZip("test.zip", "test")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Unziped done:\n " + strings.Join(files, "\n"))
	fmt.Printf("use time :%s\n", time.Since(start))
}

//解压到指定文件夹
func unZip(src string, dest string) ([]string, error) {
	//定义文件名集合载体
	var fileNames []string

	//读取压缩文件
	r, err := zip.OpenReader(src)
	if err != nil {
		return fileNames, err
	}
	defer r.Close()

	//对压缩包内的文件进行处理
	//r.File ==> []*File
	for _, f := range r.File {
		//返回文件内容信息
		fInfo, err := f.Open()
		if err != nil {
			return fileNames, err
		}
		defer fInfo.Close()

		//拼接文件路径
		fpath := filepath.Join(dest, f.Name)

		//检查压缩文件是否包含非法字符
		// os.PathSeparator = "/"
		// Clean函数通过单纯的词法操作返回和path代表同一地址的最短路径。
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fileNames, fmt.Errorf("%s: illegal is path")
		}

		//文件路径添加到容器
		fileNames = append(fileNames, fpath)

		//判断文件属性分别处理 文件夹 文件
		if f.FileInfo().IsDir() {
			//创建文件夹
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			//创建文件
			//创建文件的文件夹
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return fileNames, err
			}

			//创建文件
			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
			if err != nil {
				return fileNames, err
			}

			//数据拷贝
			_, err = io.Copy(outFile, fInfo)

			//关闭文件
			outFile.Close()
			if err != nil {
				return fileNames, err
			}
		}
	}
	return fileNames, nil
}
