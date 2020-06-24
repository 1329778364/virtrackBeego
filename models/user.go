package models

import (
	"fmt"
	_ "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id int
	Name string
	Age int
}

func init()  {
	orm.RegisterModel(new(User))
}

func GetUser() User{
	o:=orm.NewOrm()
	user := User{Id: 1}
	err := o.Read(&user)
	if err == orm.ErrNoRows {
    fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(user)
	}
	return user
}

func Updateuser()  {
	u := User{Id: 1,Name: "王立强", Age: 102}
	o := orm.NewOrm()
	o.Update(&u)
	fmt.Println("更新成功")
}