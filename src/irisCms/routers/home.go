package routers

import (
	"github.com/kataras/iris"
	"irisCms/controllers"
)

func homeRouter(app *iris.Application) {
	app.Handle("get", "/", controllers.HomeIndexHandler)
}
