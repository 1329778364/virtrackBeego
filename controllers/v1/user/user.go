package user

import (
	"encoding/json"
	. "gobeetestpro/controllers"
	. "gobeetestpro/models"
	_ "gobeetestpro/utils"
	"gobeetestpro/utils/auth"
	"gobeetestpro/utils/cache"
	_ "gobeetestpro/utils/cache"
	"io/ioutil"
	"regexp"
	"time"
	_ "time"
)

// UserController operations for User
type UserController struct {
	CommonController
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
		user.RequestResponse(4001, "手机号不能为空", nil)
	}

	isorno, _ := regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, phone)
	if !isorno {
		user.RequestResponse(4002, "手机号码不正确", nil)
	}

	if password == "" {
		user.RequestResponse(4003, "密码不能为空", nil)
	}

	// 判断手机号是否已经注册
	status := IsUserMobile(phone)
	if status {
		user.RequestResponse(4004, "此手机号已经注册", nil)
	} else {
		err := SaveUserInfo(phone, Crypto(password))
		if err == nil {
			user.RequestResponse(0, "注册成功", nil)
		} else {
			user.RequestResponse(5000, "注册失败", nil)
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
		this.RequestResponse(4001, "手机号不能为空", nil)
	}

	isorno, _ := regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, phone)
	if !isorno {
		this.RequestResponse(4002, "手机号码不正确", nil)
	}

	if password == "" {
		this.RequestResponse(4003, "密码不能为空", nil)
	}

	user := FindByUserInfo(phone)
	if user.Id == 0 {
		this.RequestResponse(4005, "您还没有注册", nil)
	}

	if !ValidatePassword(user.Password, password) {
		this.RequestResponse(4005, "输入密码错误", nil)
	}

	/* 检查是否是重复登录 */
	result, _ := cache.Get(string(user.Id))
	//TODO 优化 登录token的保存
	/* 没有登录 则生成token并进行token 和用户信息缓存 */
	token := ""
	if result == "" {
		token = auth.GenerateToken(100, user) // 默认的token过期时间
		err := cache.Set(string(user.Id), user.Phone, time.Second*time.Duration(100))
		if err != nil {
			this.RequestResponse(4005, "存储用户登录信息失败!", err)
		}
	} else {
		this.RequestResponse(4006, "请勿重复登录", nil)
	}
	this.RequestResponse(0, "登录成功", token)
}

func (this *UserController) UploadContactList() {

	var contactorder ContactOrder
	claims := this.Ctx.Input.GetData("claims").(*auth.MyCustomClaims)

	body, _ := ioutil.ReadAll(this.Ctx.Request.Body)
	_ = json.Unmarshal(body, &contactorder)

	AddContacts(contactorder, claims.User.Id)

	this.RequestResponse(200, "提交成功", nil)
}

func (this *UserController) Logout() {

}
