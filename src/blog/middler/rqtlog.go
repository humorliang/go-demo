package middler

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"blog/comm/file"
)

func LoggerConfig() gin.LoggerConfig {
	//请求日志
	today := file.TodayLogFileName("Request")
	runtimeDir := file.GetRuntimeLogPath()
	f, err := file.MustOpenSrc(today, runtimeDir)
	if err != nil {
		fmt.Println(err)
	}
	return gin.LoggerConfig{
		Output: f,
	}
}
