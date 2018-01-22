package main

import (
	"fmt"
	"os"

	"github.com/angletym2014/ecgoredis"
)

func main() {
	//获取redis实例,GetRedisClien方法的参数暂时还没有用到
	myredis := ecgoredis.GetRedisClient("corp.slave.one")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			myredis.Close()
			os.Exit(3)
		} else {
			myredis.Close()
		}
	}()

	var hmap = map[string]interface{}{"age": 31, "name": "tym", "job": "php"}
	//设置的结果是个字符串,参数是map类型 返回值:OK
	hmsetResult, hmsetErr := myredis.HMSet("mytest3", hmap).Result()
	if hmsetErr != nil {
		panic(hmsetErr)
	} else {
		fmt.Println(hmsetResult)
	}

	//获取的结果是slice,参数是string类型，例如键名，字段名 返回值:[tym 31 php]
	hmgetResult, hmgetErr := myredis.HMGet("mytest3", "name", "age", "job").Result()
	if hmgetErr != nil {
		panic(hmgetErr)
	} else {
		fmt.Println(hmgetResult)
	}

}
