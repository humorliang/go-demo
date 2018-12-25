/*
工具类：

*/
package utils

import (
	"log"
	"os"
	"strings"
	"time"
	"fmt"
)

//写入文件
//文件不存在则添加，存在直接追加
//@param path string ,filename string
//@return (f *os.File, e error)
func CreateAppendWriteFile(path string, filename string, content []byte) (*os.File, error) {
	//------------- 方法一 ------------------
	//检验文件是否存在
	//文件存在直接追加
	//if CheckFileDirIsNotExits(path, filename) {
	//	//此处以写和追加的方式
	//	file, err := os.OpenFile(path+filename, os.O_WRONLY|os.O_APPEND, 0666)
	//	if err != nil {
	//		log.Fatal("append write file error! ", err)
	//	}
	//	//写入文件
	//	_, err = file.Write(content)
	//	if err != nil {
	//		log.Fatal("write byte content error! ", err)
	//	}
	//	file.Close()
	//} else {
	//	//不存在时创建并写入
	//	err := ioutil.WriteFile(path+filename, content, 0666)
	//	if err != nil {
	//		log.Fatal("create write file content error! ", err)
	//	}
	//}
	//--------------- 方法二 -----------------
	//fmt.Println("开始判断文件夹")
	//判断文件夹是否存在
	if !CheckFileDirIsNotExits(path) {
		//fmt.Println("开始创建文件夹")
		//Mkdir 创建单个文件夹
		//MkdirAll 创建多个文件夹
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			//log.Fatal(err) 出错直接调用系统接口os.exit(1)直接退出不会执行defer
			fmt.Println("创建文件夹失败！")
			log.Panic(err) //志内容刷到标准错误后调用 panic 函数
		}else {
			//fmt.Println("创建成功")
			log.Println("创建成功")
		}
	}
	// 打开文件以创建和写和追加的方式
	f, err := os.OpenFile(path+filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Panic(err)
	}
	if content != nil {
		f.Write(content)
	}
	return f, err
}

//检查文件或文件夹是否存在
//@return bool
func CheckFileDirIsNotExits(pathOrFileName string) bool {
	// 文件或文件夹不存在则返回error
	_, err := os.Stat(pathOrFileName)
	if err != nil {
		if os.IsNotExist(err) {
			//log.Fatal("File or Path does not exist.")
			log.Println("File or Path does not exist")
			return false
		}
	}
	return true
}

//获取当天的日志文件名
//@return string
func TodayFileName() (fileName string) {
	return strings.Split(time.Now().String(), " ")[0] + ".log"
}