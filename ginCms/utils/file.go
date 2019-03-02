package utils

import (
	"os"
	"log"
	"time"
	"strings"
)

//打开一个文件 如果文件夹不存在
//@return os.File
func OpenFile(filePath string) *os.File {
	//分割文件目录和文件
	filePathSlice := strings.Split(filePath, "/")
	//单个文件直接写
	if len(filePathSlice) <= 1 {
		f, err := os.OpenFile(filePathSlice[0], os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
		if err != nil {
			log.Panic(err)
		}
		return f
	} else {
		//详细文件路径
		//判断文件夹是否存在 不存在先创建 再写入
		path := filePathSlice[:len(filePathSlice)-1]
		if CheckFileDirIsNotExits(strings.Join(path, "/")) {
			f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
			if err != nil {
				log.Panic(err)
			}
			return f
		} else {
			//创建文件夹
			err := os.MkdirAll(strings.Join(path, "/"), os.ModePerm)
			if err != nil {
				log.Panic(err)
			}
			f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
			if err != nil {
				log.Panic(err)
			}
			return f
		}
	}

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

//检查文件夹，不存在则创建
func CheckCreateDir(path string) (bool, error) {
	if CheckFileDirIsNotExits(path) {
		return true, nil
	} else {
		//创建文件夹
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Panic(err)
			return false, err
		}
		return true, nil
	}
}
