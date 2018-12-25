/*

AllMethods contains the valid http methods:
"GET", "POST", "PUT", "DELETE", "CONNECT", "HEAD",
"PATCH", "OPTIONS", "TRACE"
*/
package routers

import (
	"github.com/kataras/iris"
	"irisCms/controllers"
)

func homeRouter(app *iris.Application) {

	app.Handle("GET", "/", controllers.HomeIndexHandler)
}
