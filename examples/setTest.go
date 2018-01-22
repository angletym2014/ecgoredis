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

	//批量设置数据。结果是int,参数可以是键名-多个成员 返回值:3
	setResult, setErr := myredis.SAdd("setkey", "setmember1", "setmember2", "setmember3").Result()
	if setErr != nil {
		panic(setErr)
	} else {
		fmt.Println(setResult)
	}

	//获取集合所有成员。结果是字符串类型的slice,参数可以是键名-多个成员 返回值:[setmember3 setmember2 setmember1]
	setgetResult, setgetErr := myredis.SMembers("setkey").Result()
	if setgetErr != nil {
		panic(setgetErr)
	} else {
		fmt.Println(setgetResult)
	}
}
