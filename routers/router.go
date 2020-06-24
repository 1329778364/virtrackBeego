package routers

import (
	"gobeetestpro/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user/register", &controllers.UserController{},"post:Reister")
}
