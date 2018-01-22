package main

import (
	"fmt"
	"os"

	"github.com/angletym2014/ecgoredis"
)

func main() {
	//获取redis实例,GetRedisClien方法
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

	//批量设置数据。结果是数字 返回值:3
	listResult, listErr := myredis.LPush("listkey", "listone", "listtwo", "listthree").Result()
	if listErr != nil {
		panic(listErr)
	} else {
		fmt.Println(listResult)
	}

	//获取链式表的长度。结果是数字 返回值:3
	lenResult, lenErr := myredis.LLen("listkey").Result()
	if lenErr != nil {
		panic(lenErr)
	} else {
		fmt.Println(lenResult)
	}
}
