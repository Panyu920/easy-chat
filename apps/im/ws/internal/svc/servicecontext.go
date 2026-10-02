package svc

import "easy-chat/apps/im/ws/internal/config"

type ServiceContext struct {
	Config config.Config

	JWTAuth struct {
		Secret string
	}
}

func NewServiceContext(c *config.Config) *ServiceContext {
	return &ServiceContext{
		Config: *c,
	}
}
