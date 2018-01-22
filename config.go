package ecgoredis

import (
	"github.com/tttlkkkl/go-config"
)

//Ecredisconfig  redis详细配置信息
type Ecredisconfig struct {
	addr     string
	password string
	db       int
	timeout  int64
}

//GetRedisConfig 通过配置文件获取相关信息
func GetRedisConfig(name string) *Ecredisconfig {
	//ecredis  是包中conf目录下 ecredis.toml 文件的名字
	x := conf.C("ecredis")
	var tmpconfig = &Ecredisconfig{
		addr:     x.Get(name + ".addr").String(),
		password: x.Get(name + ".password").String(),
		db:       int(x.Get(name + ".db").Int()),
		timeout:  x.Get(name + ".timeout").Int(),
	}
	return tmpconfig
}
