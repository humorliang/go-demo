package comm

import (
	"log"
	"ginCms/utils"
)

const (
	debugDir = "../../data/log/debug"
	infoDir  = "../../data/log/info"
	errorDir = "../../data/log/error"
)

//全局的Log对象
func Log(prefix string) (l *log.Logger) {
	//f, _ := os.OpenFile("../data/log/test.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	//log.SetOutput(f)                                    //设置输出流
	//log.SetPrefix("[Test]")                             //日志前缀
	//log.SetFlags(log.Llongfile | log.Ldate | log.Ltime) //日志输出样式
	var filePath string
	todayFileName := utils.TodayFileName()
	switch prefix {
	case "debug":
		filePath = debugDir + "/" + todayFileName
	case "info":
		filePath = infoDir + "/" + todayFileName
	case "error":
		filePath = errorDir + "/" + todayFileName
	}
	l = log.New(utils.OpenFile(filePath), "["+prefix+"]", log.Llongfile|log.Ldate|log.Ltime)
	return l
}
