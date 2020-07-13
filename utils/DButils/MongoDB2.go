package DButils

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
)

var MgoSession *mgo.Session

func MongoInit() {
	use_auth := beego.AppConfig.String("use_auth")
	if use_auth == "use" {
		MgoSession, _ = mgo.Dial(beego.AppConfig.DefaultString("HostNPort4Mongo_auth", "mongodb://127.0.0.1:27017"))
	} else {
		MgoSession, _ = mgo.Dial(beego.AppConfig.DefaultString("HostNPort4Mongo", "mongodb://127.0.0.1:27017"))
	}
	MgoSession.SetMode(mgo.Monotonic, true)
}

func GetMongoSession() *mgo.Session {
	return MgoSession.Clone() // 记得close
}
