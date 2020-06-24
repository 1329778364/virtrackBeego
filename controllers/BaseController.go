package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (c *BaseController) showSuccess(data interface{}, msg string, code int) {
	var response response
	response.Data = data
	response.Msg = msg
	response.Code = code

	c.Data["json"]= &response
	c.ServeJSON()
}

func (c *BaseController) showError(msg string, code int) {
	var response response
	response.Msg = msg
	response.Code = code
	c.Data["json"]= &response
	c.ServeJSON()
}











