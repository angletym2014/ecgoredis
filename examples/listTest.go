package main

import (
	"ecgoredis"
	"fmt"
)

func main() {
	//获取redis实例,GetRedisClien方法
	myredis, myreidserr := ecgoredis.GetRedisClient("corp.slave.one")
	if myreidserr != nil {
		fmt.Println(myreidserr)
	}

	//批量设置数据。结果是数字 返回值:3
	listResult, listErr := myredis.LPush("listkey", "listone", "listtwo", "listthree").Result()
	if listErr != nil {
		fmt.Println(listErr)
	} else {
		fmt.Println(listResult)
	}

	//获取链式表的长度。结果是数字 返回值:3
	lenResult, lenErr := myredis.LLen("listkey").Result()
	if lenErr != nil {
		fmt.Println(lenErr)
	} else {
		fmt.Println(lenResult)
	}
}
