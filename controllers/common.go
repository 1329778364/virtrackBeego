package controllers

import (
	"crypto/md5"
	"encoding/hex"
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
func RequestResponse(code int, msg interface{}, data interface{}) (json *JsonStruct) {
	if data == nil {
		data = []string{}
	}
	json = &JsonStruct{Code: code, Msg: msg, Data: data}
	return
}

/* MD5加密 */
func MD5V(password string) string {
	h := md5.New()
	h.Write([]byte(password + beego.AppConfig.String("md5code")))
	return hex.EncodeToString(h.Sum(nil))
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
		fmt.Println("pw wrong")
		return false
	} else {
		fmt.Println("pw ok")
		return true
	}
}
