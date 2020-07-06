package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
)

// Operations about Users
type CommonController struct {
	beego.Controller
}

type JsonStruct struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

/* 请求结果封装返回. */
func (c *CommonController) RequestResponse(code int, msg string, data interface{}) {
	if data == nil {
		data = []string{}
	}
	json := &JsonStruct{Code: code, Msg: msg, Data: data}
	c.Data["json"] = json
	c.ServeJSON()
	c.StopRun()
}

// 加密密码
func Crypto(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	encodePW := string(hash) // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	return encodePW
}

func ValidatePassword(password string, truepwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(truepwd))
	if err != nil {
		fmt.Println("password wrong")
		return false
	} else {
		fmt.Println("password ok")
		return true
	}
}
