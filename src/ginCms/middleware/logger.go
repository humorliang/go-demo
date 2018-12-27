package middleware

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"strings"
)

//日志数据格式字段
type LogFormatterParams struct {
	Request *http.Request

	// TimeStamp shows the time after the server returns a response.
	TimeStamp time.Time
	// StatusCode is HTTP response code.
	StatusCode int
	// Latency is how much time the server cost to process a certain request.
	Latency time.Duration
	// ClientIP equals Context's ClientIP method.
	ClientIP string
	// Method is the HTTP method given to the request.
	Method string
	// Path is a path the client requests.
	Path string
	// ErrorMessage is set if error has occurred in processing the request.
	ErrorMessage string
	// IsTerm shows whether does gin's output descriptor refers to a terminal.
	IsTerm bool
}

func Logger() gin.HandlerFunc {
	//排除在外处理的扩展
	var excludeExtensions = [...]string{
		".js",
		".css",
		".jpg",
		".png",
		".ico",
		".svg",
	}

	return func(ctx *gin.Context) {
		// Start timer
		start := time.Now()
		path := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery

		// Process request
		ctx.Next()

		//遍历跳过扩展
		for _, ext := range excludeExtensions {
			//如果有后缀直接跳过
			if strings.HasSuffix(path, ext) {

			}else {

			}
		}

		//

	}
}
