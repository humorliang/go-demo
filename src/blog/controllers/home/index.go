package home

import "blog/controllers"

type IndexController struct {
	controllers.BaseController
}

func (c *IndexController) Get() {
	c.Ctx.WriteString("index")
}
