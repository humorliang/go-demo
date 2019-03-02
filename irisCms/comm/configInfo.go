package comm

import (
	"irisCms/utils"
	"os"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"fmt"
)

type DbInfo struct {
	url      string `yaml:"url"`
	port     int    `yaml:"port"`
	username string `yaml:"username"`
	password string `yaml:"password"`
	charset  string `yaml:"charset"`
}

func DbConfigRead(path string, file string) *DbInfo {
	db := DbInfo{}
	if utils.CheckFileDirIsNotExits(path + file) {
		f, err := os.OpenFile(path+file, os.O_RDONLY, 0666)
		fmt.Println("读取配置文件！")
		if err != nil {
			fmt.Println("打开文件错误")
			NewLogger("error", err)
		}
		data, _ := ioutil.ReadAll(f)
		err = yaml.Unmarshal(data, &db)
		if err != nil {
			fmt.Println("yaml错误")
			NewLogger("error", err)
		}
		return &db
	} else {
		fmt.Println("配置文件不存在！")
		NewLogger("info", "配置文件不存在！")
		return &db
	}
}
