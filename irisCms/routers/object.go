package routers

/*
routers包的公共数据类型
*/
import "github.com/kataras/iris"

//注册路由
func New(app *iris.Application) {
	homeRouter(app)
}
