package models

import (
	_ "fmt"
	_ "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"gobeetestpro/utils"

	_ "gobeetestpro/utils"
	//uuid "github.com/satori/go.uuid"
	"time"
)

func init() {
	orm.RegisterModel(new(User))
}

type User struct {
	Id    int64  `json:"id"`
	Phone string `json:"phone"`
	/** 返回的数据中移除密码字段 `json:"-"`.*/
	Password   string `json:"-"`
	Status     int    `json:"status"`
	CreateTime int64  `json:"create_time"`
	Useruuid   string `json:"user_uuid"`
}

type JwtToken struct {
	Token string `json:"token"`
}

/* 判断是否已经注册 */
func IsUserMobile(phone string) bool {
	o := orm.NewOrm()
	user := User{Phone: phone}
	err := o.Read(&user, "Phone")
	if err == orm.ErrNoRows {
		return false
	} else if err == orm.ErrMissPK {
		return false
	}
	return true
}

/* 存储用户注册信息 */
func SaveUserInfo(phone string, password string) error {
	o := orm.NewOrm()
	var user User
	user.Useruuid = utils.GetUUID(phone)
	user.Password = password
	user.Phone = phone
	user.Status = 1
	user.CreateTime = time.Now().Unix()
	_, err := o.Insert(&user)
	return err
}

/* 查找用户信息 */
func FindByUserInfo(phone string) User {
	o := orm.NewOrm()
	user := User{}
	o.QueryTable("user").Filter("phone", phone).One(&user)
	return user
}
