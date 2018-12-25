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
	"irisCms/comm"
)

//日志文件目录
const requestLogDir = "./data/log/request/"

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
		_, err := utils.CreateAppendWriteFile(requestLogDir, utils.TodayFileName(), []byte(output))
		if err != nil {
			comm.NewLogger("error","请求日志创建失败！")
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
	//请求的日志对象
	h = logger.New(cfg)
	return
}
