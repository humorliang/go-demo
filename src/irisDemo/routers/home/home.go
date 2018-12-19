package home

import (
	"github.com/kataras/iris"
)

//index
func IndexHandler(ctx iris.Context)  {
	ctx.HTML("<h1>index<h1>")
}