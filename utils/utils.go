package utils

import (
	_ "github.com/astaxie/beego"
	uuid "github.com/satori/go.uuid"
	_ "gobeetestpro/utils/consts"
)

/*-----------获取UUID--------------------*/
func GetUUID(phone string) string {
	uuid := uuid.NewV5(uuid.NamespaceDNS, phone)
	return uuid.String()
}

//func ShowSuccess(this *user.UserController) (interface{}, interface{}) {
//	this.Data["json"] = map[string]interface{}{
//		"code": consts.SUCCECC,
//		"msg":  "登出成功",
//	}
//	this.ServeJSON()
//	return nil, nil
//}
