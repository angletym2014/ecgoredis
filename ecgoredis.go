package ecgoredis

import (
	"time"

	"github.com/go-redis/redis"
)

//GetRedisClient 获取go-redis 客户端实例
func GetRedisClient(name string) *redis.Client {
	var redisOption = GetRedisConfig(name)
	//通过redis配置参数 获取redis 相关参数
	client := redis.NewClient(&redis.Options{
		Addr:        redisOption.addr,
		Password:    redisOption.password,
		DB:          redisOption.db,
		DialTimeout: time.Duration(redisOption.timeout) * time.Second,
	})

	_, pongErr := client.Ping().Result()
	if pongErr != nil {
		panic(pongErr)
	}

	return client
}
