package svc

import (
	immodel "easy-chat/apps/im/im_model"
	"easy-chat/apps/im/ws/websocketx"
	"easy-chat/apps/task/mq/internal/config"
	"easy-chat/pkg/constant"
	"net/http"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config             *config.Config
	Redisx             *redis.Redis
	WsClient           websocketx.Client
	ChatLogModel       immodel.ChatLogModel
	ConversationModel  immodel.ConversationModel
	ConversationsModel immodel.ConversationsModel
}

func NewServiceContext(c *config.Config) *ServiceContext {
	svc := &ServiceContext{
		Config:             c,
		ChatLogModel:       immodel.NewChatLogModel(c.Mongo.Url, c.Mongo.Db, c.Mongo.ChatLogCollection),
		ConversationModel:  immodel.NewConversationModel(c.Mongo.Url, c.Mongo.Db, c.Mongo.ConversationCollection),
		ConversationsModel: immodel.NewConversationsModel(c.Mongo.Url, c.Mongo.Db, c.Mongo.ConversationsCollection),
		Redisx:             redis.MustNewRedis(c.Redisx),
	}

	token, err := svc.getSystemToken()
	if err != nil {
		panic(err)
	}
	header := http.Header{}
	header.Set("Authorization", token)
	svc.WsClient = websocketx.NewClient(c.Ws.Host, websocketx.WithClientHeader(header))
	return svc
}

func (s *ServiceContext) getSystemToken() (string, error) {
	return s.Redisx.Get(constant.REDIS_KEY_ROOT_SYSTEM)
}
