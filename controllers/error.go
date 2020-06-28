package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.Data["json"] = RequestResponse(404, "您的请求不合法", nil)
	c.ServeJSON()
}

func (c *ErrorController) Error500() {
	c.Data["json"] = RequestResponse(500, "服务器异常，请联系平台处理", nil)
	c.ServeJSON()
}
