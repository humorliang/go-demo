package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/humorliang/go-demo/blog/comm/setting"
	"github.com/humorliang/go-demo/blog/middler"
	"github.com/humorliang/go-demo/blog/routers"
)

//运行模式
var runMode = flag.String("mode", "dev", "this is run mode args default dev")

func init()  {
	//命令行解析
	flag.Parse()

	//设置gin环境  gin的init函数初始化设置模式
	gin.SetMode(setting.ServerCfg.RunMode)

	//配置文件初始化
	setting.ConfigInit(runMode)

	fmt.Println("main init")
}


func main() {


	//创建router路由
	router := gin.New()
	//f, err := os.Open("info.log")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//router.Use(gin.LoggerWithWriter(io.MultiWriter(f)))
	router.Use(gin.LoggerWithConfig(middler.LoggerConfig()))
	router.Use(gin.Recovery())
	routers.Register(router)
	router.Run(":" + setting.ServerCfg.HttpPort)
}
