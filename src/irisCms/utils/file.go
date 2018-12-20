/*
工具类：

*/
package utils

import (
	"log"
	"os"
	"io/ioutil"
)

//写入文件
//文件不存在则添加，存在直接追加
func CreateAppendWriteFile(path string, filename string, content []byte) {
	//检验文件是否存在
	//文件存在直接追加
	if checkFileIsNotExits(path, filename) {
		//此处以写和追加的方式
		file, err := os.OpenFile(path+filename, os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal("append write file error! ", err)
		}
		//写入文件
		_, err = file.Write(content)
		if err != nil {
			log.Fatal("write byte content error! ", err)
		}
		file.Close()
	} else {
		//不存在时创建并写入
		err := ioutil.WriteFile(path+filename, content, 0666)
		if err != nil {
			log.Fatal("create write file content error! ", err)
		}
	}
}

//检查文件
func checkFileIsNotExits(path string, filename string) bool {
	// 文件不存在则返回error
	_, err := os.Stat(path + filename)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
			return false
		}
	}
	return true
}
