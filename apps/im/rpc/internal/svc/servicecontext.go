package svc

import (
	immodel "easy-chat/apps/im/im_model"
	"easy-chat/apps/im/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	immodel.ChatLogModel
	immodel.ConversationModel
	immodel.ConversationsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:             c,
		ChatLogModel:       immodel.NewChatLogModel(c.Mongo.Url, c.Mongo.Db, c.Mongo.Collection),
		ConversationModel:  immodel.NewConversationModel(c.Mongo.Url, c.Mongo.Db, c.Mongo.ConversationCollection),
		ConversationsModel: immodel.NewConversationsModel(c.Mongo.Url, c.Mongo.Db, c.Mongo.ConversationsCollection),
	}
}
