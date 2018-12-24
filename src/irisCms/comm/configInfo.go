package comm

import (
	"irisCms/utils"
	"os"
	"log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type DbInfo struct {
	url      string `json:"url"`
	port     int    `json:"port"`
	username string `json:"username"`
	password string `json:"password"`
	charset  string `json:"charset"`
}

func DbConfigRead(path string, file string) *DbInfo {
	if utils.CheckFileDirIsNotExits(path + file) {
		db := &DbInfo{}
		f, err := os.OpenFile(path+file, os.O_RDONLY, 0666)
		if err != nil {
			log.Fatal("配置文件不存在")
		}
		err := yaml.Unmarshal(ioutil.ReadAll(f))
		return db
	} else {
		log.Fatal("配置文件不存在！")
	}

}
