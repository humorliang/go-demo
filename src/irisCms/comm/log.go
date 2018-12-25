/*
全局变量包
*/
package comm

import (
	"log"
	"irisCms/utils"
	"fmt"
)

const (
	errorLogDir = "./data/log/error/"
	infoLogDir  = "./data/log/info/"
	debugLogDir = "./data/log/debug/"
)

func NewLogger(prefix string, errContent interface{}) (l *log.Logger) {

	switch prefix {
	case "debug":
		f, err := utils.CreateAppendWriteFile(debugLogDir, utils.TodayFileName(), nil)
		if err != nil {
			log.Panic("debug log file fail!")
		}
		l = log.New(f, "[debug]", log.Llongfile)
		l.Println(errContent)
		return
	case "error":
		f, err := utils.CreateAppendWriteFile(errorLogDir, utils.TodayFileName(), nil)
		fmt.Println("error log info")
		if err != nil {
			log.Panic("error log file fail!")
		}
		l = log.New(f, "[error]", log.Llongfile)
		l.Panic(errContent)
		return
	case "info":
		f, err := utils.CreateAppendWriteFile(infoLogDir, utils.TodayFileName(), nil)
		if err != nil {
			log.Panic("info log file fail!")
		}
		l = log.New(f, "[info]", log.Llongfile)
		l.Panic(errContent)
		return
	}
	return
}
