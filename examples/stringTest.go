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

	//删除指定的键数据。结果是int64类型,参数可以是多个键名 返回值:删除的成功个数
	/*delReult, delErr := myredis.Del("mytest3").Result()
	if delErr != nil {
		fmt.Println(delErr)
	} else {
		fmt.Println(delReult)
	}*/

	//判断指定的键名是否存在。结果是int64类型,参数可以是多个键名 返回值:1
	/*checkReult, checkErr := myredis.Exists("mytest2", "mytest3").Result()
	if checkErr != nil {
		fmt.Println(checkErr)
	} else {
		fmt.Println(checkReult)
	}*/

	//设置指定的键的有效时间。设置的结果是bool类型,参数是键名跟time.Duration类型的时间 返回值:True
	/*boolResult, boolErr := myredis.Expire("mytest2", 60*time.Second).Result()
	if boolErr != nil {
		fmt.Println(boolErr)
	} else {
		fmt.Println(boolResult)
	}*/

	//设置指定的键的有效时间。设置的结果是bool类型,参数是键名跟time.Time类型的时间 返回值:True
	/*boolAtResult, booAtlErr := myredis.ExpireAt("mytest3", time.Unix(time.Now().Unix()+360, 0)).Result()
	if booAtlErr != nil {
		fmt.Println(booAtlErr)
	} else {
		fmt.Println(boolAtResult)
	}*/

	/*var hmap = map[string]interface{}{"age": 31, "name": "tym", "job": "php"}
	//设置的结果是个字符串,参数是map类型 返回值:OK
	hmsetResult, hmsetErr := myredis.HMSet("mytest3", hmap).Result()
	if hmsetErr != nil {
		fmt.Println(hmsetErr)
	} else {
		fmt.Println(hmsetResult)
	}*/

	//获取的结果是slice,参数是string类型，例如键名，字段名 返回值:[tym 31 php]
	/*hmgetResult, hmgetErr := myredis.HMGet("mytest3", "name", "age", "job").Result()
	if hmgetErr != nil {
		fmt.Println(hmgetErr)
	} else {
		fmt.Println(hmgetResult)
	}*/

	//设置的结果是数字,参数是string类型键名 返回值:当前键值
	/*increResult, increErr := myredis.Incr("iplimit").Result()
	if increErr != nil {
		fmt.Println(increErr)
	} else {
		fmt.Println(increResult)
	}*/
}
