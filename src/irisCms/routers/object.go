package routers

/*
routers包的公共数据类型
*/
import "github.com/kataras/iris"

type BaseRequest struct {
	url     string                             //urlPath
	handler func(ctx iris.Context) interface{} //handler func
}

func New(app *iris.Application) {
	homeRouter(app)
}
