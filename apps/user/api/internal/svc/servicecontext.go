// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package svc

import (
	"easy-chat/apps/user/api/internal/config"
	"easy-chat/apps/user/rpc/userservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		UserService: userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
	}
}
