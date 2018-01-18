# ecgoredis
结合go-redis跟toml格式配置文件的简单封装

# 安装
 
  `go get -u github.com/angletym2014/ecgoredis`

# 代码示例 

进入当前包examples目录，使用go run命令执行即可看到效果：



- 字符串简单示例：stringTest.go


- 集合简单示例：setTest.go


- 链表简单示例：listTest.go


- 哈希简单示例：hashTest.go


- 其他命令简单示例： otherTest.go
 

**简单的调用demo:**

新建文件test.go，内容如下：

    package main
    
    import (
    	"fmt"
    	"time"
    
    	"github.com/angletym2014/ecgoredis"
    )
    
    func main() {
    
    	//获取redis实例,GetRedisClien方法
    	myredis, myreidserr := ecgoredis.GetRedisClient("corp.slave.one")
    	if myreidserr != nil {
    		fmt.Println(myreidserr)
    	}
    
    	//单个设置数据。结果是string,参数可以是键名-键值-缓存时间 返回值:OK
    	setResult, setErr := myredis.Set("mytest0", "myname", 120*time.Second).Result()
    	if setErr != nil {
    		fmt.Println(setErr)
    	} else {
    		fmt.Println(setResult)
    	}
    }

执行如下：

go run test.go  -conf="$GOPATH/src/github.com/angletym2014/ecgoredis/conf/ecredis.toml"  [其中$GOPATH为你的GO项目路径] 
完整的例子： go run test.go -conf="E:/goProject/src/github.com/angletym2014/ecgoredis/conf/ecredis.toml" 

输出  OK  【伴有有其他的错误暂时忽略^.^】

# 注意事项 #



1. 源码中conf/ecredis.toml 的配置文件，就是redis 链接配置信息，现在默认的ip跟端口为:127.0.0.1:6379


1. 检查本地环境redis的相关环境：ip跟端口号 是否正确


