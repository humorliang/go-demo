package controllers

import (
	"github.com/kataras/iris"
)

func HomeIndexHandler(ctx iris.Context)  {
	ctx.WriteString("index page")
}