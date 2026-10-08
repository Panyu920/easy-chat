package config

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Config struct {
	service.ServiceConf

	ListenOn string

	JwtAuth struct {
		AccessSecret string
	}
	Mongo struct {
		Url        string
		Db         string
		Collection string
	}
	MqTransfer struct {
		Host  string
		Topic string
	}

	Redisx redis.RedisConf
}
