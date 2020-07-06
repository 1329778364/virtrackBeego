package utils

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
)

var RClient *redis.Client

// 连接redis
func RedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     beego.AppConfig.String("host") + ":" + beego.AppConfig.String("redisport"),
		Password: beego.AppConfig.String("redispassword"), // no password set
		DB:       0,                                       // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client
}
