package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	jwt "github.com/dgrijalva/jwt-go"
	. "gobeetestpro/controllers"
	"gobeetestpro/controllers/v1/user"
	"gobeetestpro/utils/auth"
	_ "gobeetestpro/utils/auth"
)

func init() {
	ns :=
		beego.NewNamespace("/v1",
			//CRUD Create(创建)、Read(读取)、Update(更新)和Delete(删除)

			beego.NSNamespace("/user",
				beego.NSRouter("/login", &user.UserController{}, "*:Login"),
				beego.NSRouter("/logout", &user.UserController{}, "*:Logout"),
				beego.NSRouter("/register", &user.UserController{}, "*:Register"),
				beego.NSRouter("/uploadcontactlist", &user.UserController{}, "*:UploadContactList"),
			),
		)
	beego.InsertFilter("/v1/user/uploadcontactlist", beego.BeforeRouter, TokenFilter)
	beego.AddNamespace(ns)
}

/*
	验证是否登录以及token是否有效 同时将token 的数据进行存储 到请求头中
*/
var TokenFilter = func(ctx *context.Context) {
	tokenstr := ctx.Request.Header.Get("token") //获取前端提交过来的token 参数和前端进行约定
	/* 检查是否有token */
	if tokenstr == "" {
		ctx.Output.JSON(JsonStruct{Msg: "token为空", Data: "", Code: 4003}, false, false)
		ctx.Abort(4005, "终止")
	}
	/* 解析token */
	token, _ := jwt.ParseWithClaims(
		tokenstr,
		&auth.MyCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(auth.KEY), nil
		})

	claims, ok := token.Claims.(*auth.MyCustomClaims)
	if ok && token.Valid {
		ctx.Input.SetData("claims", claims)
	} else {
		ctx.Output.JSON(JsonStruct{Msg: "token验证失败", Data: claims, Code: 4003}, false, false)
		return
	}
}

var isExitFilter = func(ctx *context.Context) {
	ctx.Input.GetData("phone")

}
