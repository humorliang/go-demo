/*
日志中间件
*/
package middleware

import (
	"strings"
	"time"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"irisCms/utils"
	"log"
)

//日志文件目录
const (
	requestLogDir = "./data/log/request/"
	errorLogDir   = "./data/log/error/"
	infoLogDir    = "./data/log/info/"
	debugLogDir   = "./data/log/debug/"
)

//获取当天的日志文件名
//@return string
func todayFileName() (fileName string) {
	return strings.Split(time.Now().String(), " ")[0] + ".log"
}

//获取一个http请求的logger，不能代替全局的log
func NewRequestLogger() (h iris.Handler) {
	//排除在外处理的扩展
	var excludeExtensions = [...]string{
		".js",
		".css",
		".jpg",
		".png",
		".ico",
		".svg",
	}

	cfg := logger.Config{
		Status:  true,
		IP:      true,
		Method:  true,
		Path:    true,
		Columns: true,
	}
	cfg.LogFunc = func(now time.Time, latency time.Duration, status, ip, method, path string, message interface{}, headerMessage interface{}) {
		//组装信息
		output := logger.Columnize(now.String(), latency, status, ip, method, path, message, headerMessage)
		//写入数据
		_, err := utils.CreateAppendWriteFile(requestLogDir, todayFileName(), []byte(output))
		if err != nil {

		}
	}
	//跳过静态资源的请求
	cfg.AddSkipper(func(ctx iris.Context) bool {
		//获取path
		path := ctx.Path()

		//遍历跳过扩展
		for _, ext := range excludeExtensions {
			if strings.HasSuffix(path, ext) {
				return true
			}
		}
		return false
	})
}

func NewLogger(prefix string) (l *log.Logger) {

	switch prefix {
	case "debug":
		f, err := utils.CreateAppendWriteFile(debugLogDir, todayFileName(), nil)
		if err != nil {
			log.Panic("debug log file fail!")
		}
		l = log.New(f,"[debug]",log.Llongfile)
		return
	case "error":
		f, err := utils.CreateAppendWriteFile(errorLogDir, todayFileName(), nil)
		if err != nil {
			log.Panic("error log file fail!")
		}
		l = log.New(f,"[error]",log.Llongfile)
		return
	case "info":
		f, err := utils.CreateAppendWriteFile(infoLogDir, todayFileName(), nil)
		if err != nil {
			log.Panic("info log file fail!")
		}
		l = log.New(f,"[info]",log.Llongfile)
		return
	}
}
