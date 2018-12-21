/*

AllMethods contains the valid http methods:
"GET", "POST", "PUT", "DELETE", "CONNECT", "HEAD",
"PATCH", "OPTIONS", "TRACE"
*/
package routers

import (
	"github.com/kataras/iris"
	"irisCms/controllers"
	"fmt"
)

func homeRouter(app *iris.Application) {

	fmt.Println(app.config)
	app.Handle("GET", "/", controllers.HomeIndexHandler)
}
