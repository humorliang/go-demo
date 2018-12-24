/*
全局变量包
*/
package comm

import (
	"log"
	"irisCms/utils"
)

const (
	errorLogDir   = "./data/log/error/"
	infoLogDir    = "./data/log/info/"
	debugLogDir   = "./data/log/debug/"
)
func NewLogger(prefix string) (l *log.Logger) {

	switch prefix {
	case "debug":
		f, err := utils.CreateAppendWriteFile(debugLogDir, todayFileName(), nil)
		if err != nil {
			log.Panic("debug log file fail!")
		}
		l = log.New(f, "[debug]", log.Llongfile)
		return
	case "error":
		f, err := utils.CreateAppendWriteFile(errorLogDir, todayFileName(), nil)
		if err != nil {
			log.Panic("error log file fail!")
		}
		l = log.New(f, "[error]", log.Llongfile)
		return
	case "info":
		f, err := utils.CreateAppendWriteFile(infoLogDir, todayFileName(), nil)
		if err != nil {
			log.Panic("info log file fail!")
		}
		l = log.New(f, "[info]", log.Llongfile)
		return
	}
}
