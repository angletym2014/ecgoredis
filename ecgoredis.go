package ecgoredis

import (
	"time"

	"github.com/go-redis/redis"
)

var  EcGoRedisConnMap = make(map[string]*redis.Client)

// GetServiceClient 获取redis链接实例
func GetRedisServiceClient(key string) *redis.Client{
	if _,ok := EcGoRedisConnMap[key]; !ok{
		EcGoRedisConnMap[key] = getRedisClient(key)
	}

	//fmt.Println(EcGoRedisConnMap[key].PoolStats().FreeConns)

	/*if EcGoRedisConnMap[key].PoolStats().FreeConns == 0 {
		//fmt.Println("pool is empty,new init")
		EcGoRedisConnMap[key] = getRedisClient(key)
	}*/

	return EcGoRedisConnMap[key]
}


// 初始化redis链接
func getRedisClient(name string) *redis.Client {
	var redisOption = GetRedisConfig(name)
	//通过redis配置参数 获取redis 相关参数
	client := redis.NewClient(&redis.Options{
		Addr:        redisOption.addr,
		Password:    redisOption.password,
		DB:          redisOption.db,
		DialTimeout: time.Duration(redisOption.timeout) * time.Second,
	})

	//fmt.Println("init client")
	_, pongErr := client.Ping().Result()
	if pongErr != nil {
		panic(pongErr)
	}

	return client
}
