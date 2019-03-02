package main

import (
	"os"
	"archive/zip"
	"io"
	"log"
	"fmt"
)

/*
Create Zip Files in Go
创建压缩文件
*/

func main() {
	//压缩文件列表
	files := []string{"test.txt", "a.jpg"}
	outFile := "test.zip"
	if err := ZipFiles(outFile, files); err != nil {
		log.Fatal(err)
	}
	fmt.Println("zipped file:", outFile)

}

//压缩文件
func ZipFiles(filename string, files []string) error {
	//创建一个压缩文件
	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	//创建一个压缩文件写对象
	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	//把文件添加到压缩文件
	for _, file := range files {
		if err := AddFileToZipFile(zipWriter, file); err != nil {
			return err
		}
	}
	return nil
}

//添加文件到压缩文件
func AddFileToZipFile(zw *zip.Writer, filename string) error {
	//打开文件
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	//文件信息
	fInfo, err := file.Stat()
	if err != nil {
		return err
	}

	//压缩文件信息头 基础信息
	header, err := zip.FileInfoHeader(fInfo)
	if err != nil {
		return err
	}
	//添加详细的文件头信息
	header.Name = filename
	//压缩方法
	header.Method = zip.Deflate

	//写入压缩文件
	writer, err := zw.CreateHeader(header)
	if err != nil {
		return err
	}
	//复制文件内容
	_, err = io.Copy(writer, file)
	if err != nil {
		return err
	}
	return nil
}
