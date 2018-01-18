package main

import (
	"fmt"
	"time"

	"github.com/angletym2014/ecgoredis"
)

func main() {
	//获取redis实例,GetRedisClien方法的参数暂时还没有用到
	myredis, myreidserr := ecgoredis.GetRedisClient("corp.slave.one")
	if myreidserr != nil {
		fmt.Println(myreidserr)
	}

	//单个设置数据。结果是string,参数可以是键名-键值-缓存时间 返回值:OK
	setResult, setErr := myredis.Set("mytest0", "tymtest", 120*time.Second).Result()
	if setErr != nil {
		fmt.Println(setErr)
	} else {
		fmt.Println(setResult)
	}

	//单个获取数据。结果是byte,参数可以是多个键名跟键值
	getVal, getErr := myredis.Get("mytest0").Result()
	if getErr != nil {
		fmt.Println(getErr)
	} else {
		fmt.Println(getVal)
	}

	//批量设置数据。结果是string,参数可以是多个键名跟键值
	msetErr := myredis.MSet("mytest1", "value1", "mytest2", "value2").Err()
	if msetErr != nil {
		fmt.Println(msetErr)
	}

	//批量获取数据。结果是slice,参数可以是多个键名 返回值:[<nil> <nil> <nil>]
	mgetval, mgetErr := myredis.MGet("mytest1", "mytest2").Result()
	if mgetErr != nil {
		fmt.Println(mgetErr)
	} else {
		fmt.Println(mgetval)
	}

	//设置的结果是数字,参数是string类型键名 返回值:当前键值
	increResult, increErr := myredis.Incr("iplimit").Result()
	if increErr != nil {
		fmt.Println(increErr)
	} else {
		fmt.Println(increResult)
	}
}
