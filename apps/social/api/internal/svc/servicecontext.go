// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package svc

import (
	"easy-chat/apps/social/api/internal/config"
	"easy-chat/apps/social/rpc/socialservice"
	"easy-chat/apps/user/rpc/userservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	userservice.UserService
	socialservice.SocialService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserService:   userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
		SocialService: socialservice.NewSocialService(zrpc.MustNewClient(c.SocialRpc)),
	}
}
