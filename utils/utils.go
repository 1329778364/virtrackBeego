package utils

import (
	"crypto/md5"
	"encoding/hex"
	_ "github.com/astaxie/beego"
	uuid "github.com/satori/go.uuid"
	_ "gobeetestpro/utils/consts"
	"io"
	"mime/multipart"
)

/*-----------获取UUID--------------------*/
func GetUUID(phone string) string {
	newV5 := uuid.NewV5(uuid.NamespaceDNS, phone)
	return newV5.String()
}

//func ShowSuccess(this *user.UserController) (interface{}, interface{}) {
//	this.Data["json"] = map[string]interface{}{
//		"code": consts.SUCCECC,
//		"msg":  "登出成功",
//	}
//	this.ServeJSON()
//	return nil, nil
//}

func Str2Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func GetFileMd5(file multipart.File) (md5Str string) {
	h := md5.New()
	if _, err := io.Copy(h, file); err != nil {
		//fmt.Printf("Get file md5 error: %v", err)
	}
	md5Str = hex.EncodeToString(h.Sum(nil))
	return md5Str
}
