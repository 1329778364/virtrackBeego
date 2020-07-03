package user

import (
	"fmt"
	"github.com/astaxie/beego"
	. "gobeetestpro/controllers"
	"gobeetestpro/models"
	"gobeetestpro/utils/auth"
	"regexp"
)

// UserController operations for User
type UserController struct {
	beego.Controller
}

// @Description 用户注册功能
// @Success 200 {object} models.User.SaveUserInfo
// @Param   mobile		formData	string   true   "手机号码"
// @Param   password   	formData    string   true 	"登录密码"
// @router /register [post]
func (user *UserController) Register() {

	phone := user.GetString("phone")
	password := user.GetString("password")
	if phone == "" {
		user.Data["json"] = RequestResponse(4001, "手机号不能为空", nil)
		user.ServeJSON()
	}

	isorno, _ := regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, phone)
	if !isorno {
		user.Data["json"] = RequestResponse(4002, "手机号码不正确", nil)
		user.ServeJSON()
		return
	}

	if password == "" {
		user.Data["json"] = RequestResponse(4003, "密码不能为空", nil)
		user.ServeJSON()
		return
	}

	// 判断手机号是否已经注册
	status := models.IsUserMobile(phone)
	if status {
		user.Data["json"] = RequestResponse(4004, "此手机号已经注册", nil)
		user.ServeJSON()
		return
	} else {
		err := models.SaveUserInfo(phone, Crypto(password))
		if err == nil {
			user.Data["json"] = RequestResponse(0, "注册成功", nil)
			user.ServeJSON()
			return
		} else {
			user.Data["json"] = RequestResponse(5000, "注册失败", nil)
			user.ServeJSON()
			return
		}
	}
}

// @Description 用户登录功能
// @Success 200 {object} models.User.FindByUserInfo
// @Param   phone		query	string   true   "手机号码"
// @Param   password	query	string   true 	"登录密码"
// @router /login [get]
func (this *UserController) Login() {

	phone := this.GetString("phone")
	password := this.GetString("password")

	if phone == "" {
		this.Data["json"] = RequestResponse(4001, "手机号不能为空", nil)
		this.ServeJSON()
		return
	}

	isorno, _ := regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, phone)
	if !isorno {
		this.Data["json"] = RequestResponse(4002, "手机号码不正确", nil)
		this.ServeJSON()
	}

	if password == "" {
		this.Data["json"] = RequestResponse(4003, "密码不能为空", nil)
		this.ServeJSON()
	}

	user := models.FindByUserInfo(phone)
	if user.Id == 0 {
		this.Data["json"] = RequestResponse(4005, "您还没有注册", nil)
		this.ServeJSON()
		return
	}

	if !ValidatePassword(user.Password, password) {
		this.Data["json"] = RequestResponse(4005, "输入密码错误", nil)
		this.ServeJSON()
		return
	}

	token := auth.GenerateToken(30 * 24 * 3600) //默认的token过期时间

	this.Data["json"] = RequestResponse(0, "登录成功", token)
	this.ServeJSON()
}

func (this *UserController) UploadContactList() {
	data := this.Ctx.Input.RequestBody
	fmt.Printf("%s", data)
	this.Data["json"] = RequestResponse(400, "提交成功", nil)
	this.ServeJSON()
}
