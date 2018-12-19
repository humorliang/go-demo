/*
全局配置使用包，应用初始化配置
*/
package app

import (
	"github.com/kataras/iris"
	recover2 "github.com/kataras/iris/middleware/recover"
	"fmt"
	"irisDemo/utils"
	"irisDemo/controllers"
)

type BaseResponse struct {
	code int    `json:"code"`
	msg  string `json:"msg"`
	data interface{}
}

//工厂函数初始化
func InitApp() *iris.Application {
	//全局实例app对象
	var app = iris.New()
	//创建日志对象
	reqLog, close := utils.NewRequestLogger()
	defer close()

	//所有域的路由上 或使用`UseGlobal`注册一个将跨子域触发的中间件
	//注册日志中间件
	app.Use(reqLog)
	//异常处理机制
	app.Use(recover2.New())
	//路由之前处理函数
	app.Use(before)
	//路由之后处理函数
	app.Done(after)

	//注册路由
	controllers.InitRouter(app)

	//请求异常统一返回
	app.OnAnyErrorCode(reqLog, func(ctx iris.Context) {
		//错误组装
		restul := BaseResponse{
			code: ctx.GetStatusCode(),
			msg:  "fail",
			data: ctx.GetStatusCode(),
		}
		ctx.JSON(&restul)
	})
	//返回iris应用实例
	return app
}

//应用程序处理请求之前 执行的函数
func before(ctx iris.Context) {
	fmt.Println("请求处理函数之前处理函数")
}

//处理程序处理之后 执行的函数
func after(ctx iris.Context) {
	fmt.Println("请求处理函数之后处理函数")
}

//判断处理成功请求
func (br BaseResponse) Success() bool {
	return br.code == 0 && br.msg == "success"
}
