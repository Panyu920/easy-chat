package user

import (
	"easy-chat/apps/im/ws/internal/svc"
	websocketx "easy-chat/apps/im/ws/websocketx"
)

func SendOnline(svc *svc.ServiceContext) websocketx.HandlerFunc {
	return func(server *websocketx.Server, conn *websocketx.Connection, msg *websocketx.Message) {
		userIds := server.GetUsers()

		currentUserIds := server.GetUsers(conn)

		err := server.SendToUser(websocketx.NewMessage(msg.MsgType, msg.Method, currentUserIds[0], msg.Data), userIds...)
		server.Logger.Error(err)
	}
}
