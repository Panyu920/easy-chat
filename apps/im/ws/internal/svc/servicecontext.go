package svc

import (
	immodel "easy-chat/apps/im/im_model"
	"easy-chat/apps/im/ws/internal/config"
	"easy-chat/apps/task/mq/mqclient"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config

	immodel.ChatLogModel
	MqClient     mqclient.IMqClient
	RedisxClient *redis.Redis
}

func NewServiceContext(c *config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       *c,
		ChatLogModel: immodel.NewChatLogModel(c.Mongo.Url, c.Mongo.Db, c.Mongo.Collection),
		MqClient:     mqclient.NewMqClient([]string{c.MqTransfer.Host}, c.MqTransfer.Topic),
		RedisxClient: redis.MustNewRedis(c.Redisx),
	}
}
