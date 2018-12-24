/*
日志处理模块
*/

package utils

//package main

import (
	"time"
	"strings"
	"os"
	"log"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"fmt"
	"io/ioutil"
)

//异常退出删除日志文件
const deleteFileOnExit = false

//排除在外处理的扩展
var excludeExtensions = [...]string{
	".js",
	".css",
	".jpg",
	".png",
	".ico",
	".svg",
}
//日志文件名
//创建当天的日志文件
//@path 日子文件绝对路径
func todayFilename(path string) string {
	now := time.Now().String()
	today := strings.Split(now, " ")[0]
	//fmt.Println("todayFilename........")
	return path + today + ".txt"
}

//创建一个新的日志文件
func newLogFile() *os.File {
	//文件名
	filename := todayFilename("/Users/datatist001/gopath/data/log/")
	//创建文件
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("newLogFile........")
	return f
}

func newLogWriteFile(content []byte) {
	//文件名
	filename := todayFilename("/Users/datatist001/gopath/data/log/")
	ioutil.WriteFile(filename, content, 0666)
}

//请求日志
//返回 h请求手柄  err 错误
func NewRequestLogger() (h iris.Handler, close func() error) {
	//日志中间件的配置
	cfg := logger.Config{
		Status:  true,
		IP:      true,
		Method:  true,
		Path:    true,
		Columns: true,
	}
	//默认返回没有错误
	close = func() error {
		return nil
	}

	//fmt.Println("logger new reqError 1........")
	//log内容格式函数
	cfg.LogFunc = func(now time.Time, latency time.Duration, status, ip, method, path string, message interface{}, headerMessage interface{}) {
		nowStr := strings.Split(now.String(), " ")
		output := logger.Columnize(nowStr[0]+" "+nowStr[1], latency, status, ip, method, path, message, headerMessage)
		//内容写入文件
		fmt.Println("logger data in write........")
		fmt.Println(output)
		//fmt.Println(now)
		newLogWriteFile([]byte(output))
	}
	//fmt.Println("logger new reqError 2........")

	//添加跳过的日志的路由
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
	fmt.Println("logger new reqError........")
	//日志中间件配置
	h = logger.New(cfg)
	return
}
