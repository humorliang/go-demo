package middler

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"blog/comm/file"
)

func LoggerConfig() gin.LoggerConfig {
	f, err := file.OpenCreateAppend("info.log")
	if err != nil {
		fmt.Println(err)
	}
	return gin.LoggerConfig{
		Output: f,
	}
}
