package svc

import (
	immodel "easy-chat/apps/im/im_model"
	"easy-chat/apps/im/ws/internal/config"
)

type ServiceContext struct {
	Config config.Config

	immodel.ChatLogModel
}

func NewServiceContext(c *config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       *c,
		ChatLogModel: immodel.NewChatLogModel(c.Mongo.Url, c.Mongo.Db, c.Mongo.Collection),
	}
}
