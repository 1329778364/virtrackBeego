package cache

import (
	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
	"time"
)

var (
	client *redis.Client
	err    error
)

func init() {

	client = redis.NewClient(&redis.Options{
		Addr:     beego.AppConfig.String("host") + ":" + beego.AppConfig.String("redisport"),
		Password: beego.AppConfig.String("redispassword"), // no password set
		DB:       0,                                       // use default DB
	})
}

// 存
func Set(key, value string, expiration time.Duration) error {
	err := client.Set(key, value, expiration).Err()
	return err
}

// 查
func Get(key string) (string, error) {
	result, err := client.Get(key).Result()
	return result, err
}

// 删
func Del(key string) error {
	err := client.Del(key).Err()
	return err
}
