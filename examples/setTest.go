package main

import (
	"ecgoredis"
	"fmt"
)

func main() {
	//获取redis实例,GetRedisClien方法的参数暂时还没有用到
	myredis, myreidserr := ecgoredis.GetRedisClient("corp.slave.one")
	if myreidserr != nil {
		fmt.Println(myreidserr)
	}

	//批量设置数据。结果是int,参数可以是键名-多个成员 返回值:3
	setResult, setErr := myredis.SAdd("setkey", "setmember1", "setmember2", "setmember3").Result()
	if setErr != nil {
		fmt.Println(setErr)
	} else {
		fmt.Println(setResult)
	}

	//获取集合所有成员。结果是字符串类型的slice,参数可以是键名-多个成员 返回值:[setmember3 setmember2 setmember1]
	setgetResult, setgetErr := myredis.SMembers("setkey").Result()
	if setgetErr != nil {
		fmt.Println(setgetErr)
	} else {
		fmt.Println(setgetResult)
	}
}
