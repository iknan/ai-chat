package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	Mysql struct {
		Host   string
		Port   int
		User   string
		Pass   string
		DbName string
	}
	Redis struct {
		Host string
		Type string
		User string
		Pass string
		Tls  bool
		Db   int
	}
	LogConf struct {
		Level string
		Path  string
		Mode  string
	}
	AliYun struct {
		AccessKeyId     string
		AccessKeySecret string
		Region          string
	}
}
