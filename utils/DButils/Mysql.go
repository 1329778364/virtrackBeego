package DButils

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

/*
 全局初始化
*/
func init() {
	driverName := beego.AppConfig.String("driverName")
	orm.RegisterDriver(driverName, orm.DRMySQL)

	user := beego.AppConfig.String("mysqlUser")
	pwd := beego.AppConfig.String("mysqlPwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbName")

	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	err := orm.RegisterDataBase("default", driverName, dbConn)

	if err != nil {
		fmt.Printf("数据库错误, err：%v", err)
		return
	}
	fmt.Println("Connected to MySQL !")
}
