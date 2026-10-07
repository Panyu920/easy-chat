package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	service.ServiceConf
	ListenOn string

	MsgChatTransfer kq.KqConf

	Redisx redis.RedisConf
	Mongo  struct {
		Url                     string
		Db                      string
		ChatLogCollection       string
		ConversationCollection  string
		ConversationsCollection string
	}

	Ws struct {
		Host string
	}
	SocialRpc zrpc.RpcClientConf
}
