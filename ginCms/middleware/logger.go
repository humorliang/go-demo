package middleware

import (
	"time"
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"strings"
	"github.com/humorliang/go-demo/ginCms/utils"
)

const requestDir = "../../car/log/request"

//日志数据格式字段
type LogFormatterParams struct {
	Request *http.Request

	// TimeStamp shows the time after the server returns a response.
	TimeStamp string
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
	//IsTerm bool
	// judgement path
	IsSkipPath bool
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
		var path string
		var rqFilePath string
		// Start timer
		start := time.Now()
		urlPath := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery

		if raw != "" {
			path = urlPath + raw
		} else {
			path = urlPath
		}
		// Process request
		ctx.Next()

		//复制一个新的请求
		var newCtx = ctx.Copy()

		//开启一个go程处理
		go func() {
			//请求
			param := LogFormatterParams{
				Request: newCtx.Request,
			}
			//遍历跳过扩展
			for _, ext := range excludeExtensions {
				//如果有后缀直接跳过
				if strings.HasSuffix(path, ext) {
					//...
					param.IsSkipPath = true
				} else {
					// Stop timer
					param.IsSkipPath = false
				}
			}
			if param.IsSkipPath {
				//...
			} else {
				param.Latency = time.Now().Sub(start)
				param.ClientIP = newCtx.ClientIP()
				param.Method = newCtx.Request.Method
				param.StatusCode = newCtx.Writer.Status()
				param.ErrorMessage = newCtx.Errors.ByType(gin.ErrorTypePrivate).String()
				//写入到请求日志
				reqInfo := fmt.Sprintf("[StartTime] %v [IP] %s [Path] %s [Method] %s [LtcyTime] %v [StatusCode] %d [ErrorInfo] %s\n",
					start.Format("2006/01/02 - 15:04:05"), param.ClientIP,
					path, param.Method, param.Latency, param.StatusCode, param.ErrorMessage)
				//日志文件路径
				rqFilePath = requestDir + "/" + utils.TodayFileName()
				//打开日志文件
				f := utils.OpenFile(rqFilePath)
				//写入日志文件
				f.WriteString(reqInfo)
			}
		}()
	}
}
