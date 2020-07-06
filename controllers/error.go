package controllers

type ErrorController struct {
	CommonController
}

func (c *ErrorController) Error404() {
	c.RequestResponse(404, "您的请求不合法", nil)
}

func (c *ErrorController) Error500() {
	c.RequestResponse(500, "服务器异常，请联系平台处理", nil)
}
