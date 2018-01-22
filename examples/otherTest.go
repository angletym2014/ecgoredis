package main

import (
	"fmt"
	"os"
	"time"

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

	//删除指定的键数据。结果是int64类型,参数可以是多个键名 返回值:删除的成功个数
	delReult, delErr := myredis.Del("mytest3").Result()
	if delErr != nil {
		panic(delErr)
	} else {
		fmt.Println(delReult)
	}

	//判断指定的键名是否存在。结果是int64类型,参数可以是多个键名 返回值:1
	checkReult, checkErr := myredis.Exists("mytest2", "mytest3").Result()
	if checkErr != nil {
		panic(checkErr)
	} else {
		fmt.Println(checkReult)
	}

	//设置指定的键的有效时间。设置的结果是bool类型,参数是键名跟time.Duration类型的时间 返回值:True
	boolResult, boolErr := myredis.Expire("mytest2", 60*time.Second).Result()
	if boolErr != nil {
		panic(boolErr)
	} else {
		fmt.Println(boolResult)
	}

	//设置指定的键的有效时间。设置的结果是bool类型,参数是键名跟time.Time类型的时间 返回值:True
	boolAtResult, booAtlErr := myredis.ExpireAt("mytest3", time.Unix(time.Now().Unix()+360, 0)).Result()
	if booAtlErr != nil {
		panic(booAtlErr)
	} else {
		fmt.Println(boolAtResult)
	}
}
