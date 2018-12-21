package controllers

import (
	"github.com/kataras/iris"
	"fmt"
)

func HomeIndexHandler(ctx iris.Context)  {
	fmt.Println()
	ctx.WriteString("index page")
}