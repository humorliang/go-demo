package log

import (
	"runtime"
	"fmt"
	"log"
	"os"
	"blog/comm/file"
)

var logger = log.New(os.Stderr, "", log.LstdFlags)

func Debug(v interface{}) {
	f, err := file.MustOpenSrc(file.TodayLogFileName("DEBUG"), file.GetRuntimeLogPath())
	if err != nil {
		log.Println("debug err:", err)
	}
	logger.SetOutput(f)
	logger.SetPrefix(prefix("DEBUG"))
	logger.Println(v)

}
func Error(v interface{}) {
	f, err := file.MustOpenSrc(file.TodayLogFileName("ERROR"), file.GetRuntimeLogPath())
	if err != nil {
		log.Println("error err:", err)
	}
	logger.SetOutput(f)
	logger.SetPrefix(prefix("ERROR"))
	logger.Println(v)
}
func Fatal(v interface{}) {
	f, err := file.MustOpenSrc(file.TodayLogFileName("FATAL"), file.GetRuntimeLogPath())
	if err != nil {
		log.Println("fatal err:", err)
	}
	logger.SetOutput(f)
	logger.SetPrefix(prefix("FATAL"))
	logger.Fatal(v)
}

func prefix(level string) string {
	//可以返回运行时正在执行的文件名和行号 0 当前函数 1 上一层函数 2 上两层
	var prefix string
	_, filePath, line, ok := runtime.Caller(2)
	if ok {
		prefix = fmt.Sprintf("[%s][%s:%d]", level, filePath, line)
	} else {
		prefix = fmt.Sprintf("[%s]", level)
	}
	return prefix
}
