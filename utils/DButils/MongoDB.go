package DButils

import (
	"context"
	"fmt"
	"github.com/astaxie/beego"
	. "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	database *Database
)

func init() {
	//mongoUsername := beego.AppConfig.String("mongo_username")
	//mongoPassword := beego.AppConfig.String("mongo_pwd")
	mongoHost := beego.AppConfig.String("mongo_host")
	mongoPort := beego.AppConfig.String("mongo_port")
	monogoDb := beego.AppConfig.String("monogo_db")

	url := "mongodb://" + mongoHost + ":" + mongoPort

	clientOptions := options.Client().ApplyURI(url)

	// Connect to MongoDB
	client, err := Connect(context.TODO(), clientOptions)

	database = client.Database(monogoDb)

	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB !")

}

/*
	获取数据库中的集合 根据集合的名字
*/

func GetMongoCollection(collection string) *Collection {
	return database.Collection(collection)
}

func DisConnectCient(client Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
