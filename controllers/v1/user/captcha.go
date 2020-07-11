package user

import (
	"github.com/mojocn/base64Captcha"
	. "gobeetestpro/controllers"
	"image/color"
)

type CaptchaController struct {
	CommonController
}

// Captcha 图形验证码
type Captcha struct {
	ID     string
	B64s   string
	Answer string
}

var store = base64Captcha.DefaultMemStore

// GetCaptcha 获取验证码
func GetCaptcha() (id string, b64s string, err error) {
	bgcolor := color.RGBA{0, 0, 0, 0}
	fonts := []string{"wqy-microhei.ttc"}
	driver := base64Captcha.NewDriverMath(40, 102, 0, 0, &bgcolor, fonts)
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err = captcha.Generate()
	return id, b64s, err
}

// VerifyCaptcha 校验图形验证码
func VerifyCaptcha(id string, answer string) bool {
	return store.Verify(id, answer, true)
}

func (c *CaptchaController) Get() {
	id, b64s, err := GetCaptcha()
	if err != nil {
		c.RequestResponse(400, "failed", err.Error())
	} else {
		c.RequestResponse(200, "seccess", Captcha{ID: id, B64s: b64s})
	}
}

func (c *CaptchaController) Post() {
	var captcha Captcha
	captcha.ID = c.GetString("id")
	captcha.Answer = c.GetString("answer")
	if VerifyCaptcha(captcha.ID, captcha.Answer) {
		c.RequestResponse(300, "success", captcha)
	} else {
		c.RequestResponse(300, "fail", captcha)
	}
}
