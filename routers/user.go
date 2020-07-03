package routers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"gobeetestpro/controllers/v1/user"
	"gobeetestpro/utils/auth"
	_ "gobeetestpro/utils/auth"
	"strings"
)

func init() {
	ns :=
		beego.NewNamespace("/v1",
			//CRUD Create(创建)、Read(读取)、Update(更新)和Delete(删除)
			beego.NSBefore(func(ctx *context.Context) { //curd之前的验证token
				split := strings.Split(ctx.Request.RequestURI, "?")
				if split[0] == "/v1/user/uploadcontactlist" {
					token := ctx.Request.Header.Get("token") //获取前端提交过来的token 参数和前端进行约定
					fmt.Print(token)
					if token == "" {
						fmt.Print("请登录获取token")
						//RequestResponse(3003, "请登录获取token", [])
					}
					err := auth.ValidateToken(token)
					if err != nil {
						fmt.Print("token验证失败")
						//ctx.WriteString(auth.GenSimpleRespString(0,"token验证失败"))
					}
				}
			}),

			beego.NSNamespace("/user",
				//beego.InsertFilter("/uploadcontactlist", beego.BeforeRouter)

				beego.NSRouter("/login", &user.UserController{}, "*:Login"),
				beego.NSRouter("/register", &user.UserController{}, "*:Register"),
				beego.NSRouter("/uploadcontactlist", &user.UserController{}, "*:UploadContactList"),
			),
		)
	beego.AddNamespace(ns)
}
