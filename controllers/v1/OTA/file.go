package OTA

import (
	"context"
	"fmt"
	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	. "gobeetestpro/controllers"
	"gobeetestpro/utils"
	. "gobeetestpro/utils/DBUtils"
)

type APKversion struct {
	Version     string `json:"version"`
	Date        string `json:"date"`
	VersionCode string `json:"version_code"`
}

// FileController operations for File
type FileController struct {
	CommonController
}

var apkVersionColl *mongo.Collection

func init() {
	apkVersionColl = GetMongoCollection("ApkVersion")
}

func (this *FileController) Upload() {

	//获取 请求的参数
	apkCode := this.GetString("check_code")
	date := this.GetString("date")
	version := this.GetString("version")
	apkFile, _, errMissFile := this.GetFile("file")

	TokenKey := beego.AppConfig.String("TokenKey")
	checkCode := utils.Str2Md5(version + date + TokenKey)

	if errMissFile != nil || apkCode != checkCode {
		//this.RequestResponse(400,"文件为空或校验码错误",nil)
	}

	versionCode := utils.GetFileMd5(apkFile)

	/* 检查当前版本是否存在 */
	var apkver APKversion
	filter := bson.D{{"text", "hello"}}
	_ = apkVersionColl.FindOne(context.TODO(), filter).Decode(&apkver)

	if apkver.Version != "" {
		this.RequestResponse(126, "版本号已存在", apkver)
	}

	/*在数据库中记录版本号*/
	apkVersion := APKversion{}
	apkVersion.Version = version
	apkVersion.Date = date
	apkVersion.VersionCode = versionCode

	_, err := apkVersionColl.InsertOne(context.TODO(), apkVersion)
	if err != nil {
		this.RequestResponse(1000, "存储版本号错误", err)
	}

	path := "./upload/Android/" + version
	_ = apkFile.Close()
	_ = this.SaveToFile("file", path) //存文件

	this.RequestResponse(200, "上传APP成功", "")
}

func (this *FileController) Download() {
	version := this.Ctx.Input.Param(":version")
	fmt.Println(version)
	if version != "" {
		path := "./upload/Android/" + version
		this.Ctx.Output.Download(path)
	}
	this.RequestResponse(200, "找到更新包", version)
}

/*
检查版本更新 输入版本号
*/
func (this *FileController) CheckUpdate() {
	version := this.GetString("version")
	if version != "" {
		newApkVer := APKversion{}
		cursor, err := apkVersionColl.Find(context.TODO(), bson.M{}, &options.FindOptions{
			Sort: bson.M{
				"date": -1,
			},
		})

		fmt.Print(cursor, err)

		update := make(map[string]interface{})
		if newApkVer.Version != version {
			update["is_update"] = 1
			update["update_url"] = "/OTA/download/" + newApkVer.Version
			update["version_code"] = newApkVer.VersionCode
			this.RequestResponse(200, "检测到更新版本："+newApkVer.Version, update)
		} else {
			this.RequestResponse(304, "不需要更新，最新版为:"+newApkVer.Version+"当前为："+version, newApkVer.Version)
		}
	}
}

func getLatestVersion() {
	options.Find()
}
