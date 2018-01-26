package main

import (
	"github.com/angletym2014/ecgoredis"
	"strconv"
	"time"
	"fmt"
)


func main()  {

	myredis := ecgoredis.GetRedisServiceClient("corp.slave.one")

	for i:=0; i<10; i++ {
		_, pongErr := myredis.Ping().Result()
		fmt.Println("ping:",pongErr)
		setResult, setErr := myredis.Set("mytest0"+strconv.Itoa(i), "myname"+strconv.Itoa(i), 120*time.Second).Result()
		if setErr != nil {
			fmt.Println("set:",setErr)
		} else {
			fmt.Println(setResult)
		}

		time.Sleep(10*time.Second)
	}


	/*fmt.Printf("%#v\n", ecgoredis.EcGoRedisConnMap)
	stats1 := myredis.PoolStats()
	fmt.Println("statues1:",stats1)

	var wg = sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer  wg.Done()
		myredis2 := ecgoredis.GetRedisServiceClient("corp.slave.one")
		fmt.Printf("%#v\n", ecgoredis.EcGoRedisConnMap)
		setResult2, setErr2 := myredis2.Set("mytest2", "myname2", 120*time.Second).Result()
		if setErr2 != nil {
			fmt.Println(setErr2)
		} else {
			fmt.Println("2:",setResult2)
		}
	}()

	go func() {
		defer  wg.Done()
		myredis3 := ecgoredis.GetRedisServiceClient("corp.slave.one")
		fmt.Printf("%#v\n", ecgoredis.EcGoRedisConnMap)
		setResult3, setErr3 := myredis3.Set("mytest3", "myname3", 120*time.Second).Result()
		if setErr3 != nil {
			fmt.Println(setErr3)
		} else {
			fmt.Println("3:",setResult3)
		}
	}()
	wg.Wait()

	stats2 := myredis.PoolStats()
	fmt.Println("statues1:",stats2)*/
}
