package handler

import (
	"easy-chat/apps/im/ws/internal/handler/user"
	"easy-chat/apps/im/ws/internal/svc"
	"easy-chat/apps/im/ws/websocketx"
)

func RegisterRoutes(server *websocketx.Server, ctx *svc.ServiceContext) {
	server.RegisterRoutes(
		[]*websocketx.Route{
			{
				Method:  "user.online",
				Handler: user.SendOnline(ctx),
			},
		},
	)
}
