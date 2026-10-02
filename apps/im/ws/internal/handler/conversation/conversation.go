package conversation

import (
	"context"
	"easy-chat/apps/im/ws/internal/logic"
	"easy-chat/apps/im/ws/internal/svc"
	msgtype "easy-chat/apps/im/ws/msg_type"
	"easy-chat/apps/im/ws/websocketx"
	"easy-chat/pkg/constant"

	"github.com/mitchellh/mapstructure"
)

func Chat(svc *svc.ServiceContext) websocketx.HandlerFunc {
	return func(server *websocketx.Server, conn *websocketx.Connection, msg *websocketx.Message) {
		//
		var chatData msgtype.Chat
		if err := mapstructure.Decode(msg.Data, &chatData); err != nil {
			server.Logger.Error(err)
			server.Send("Decode chat data failed", conn)
			return
		}

		switch chatData.ChatType {
		case constant.ChatTypeSingle:
			// 单聊
			// 存储聊天记录
			err := logic.NewConversationLogic(context.Background(), svc, server).SingleChat(&chatData, conn.GetUserID())
			if err != nil {
				server.Logger.Error(err)
				server.Send("save chat log failed", conn)
				return
			}

			// 发送消息给接收者
			err = server.SendToUser(websocketx.NewMessage(websocketx.FrameTypeData, msg.Method, conn.GetUserID(), msg.Data), chatData.RecvId)
			server.Logger.Info("recv id", chatData.RecvId)
			if err != nil {
				server.Logger.Errorf("SendToUser failed err  %v ,from %s to %s,msgType  %v, msg  %s ", err, conn.GetUserID(), chatData.RecvId, chatData.MsgType, chatData.Content)
				server.Send("SendToUser failed", conn)
				return
			}

			return
		case constant.ChatTypeGroup:
			// 群聊
		default:
			// 其他类型
		}

	}
}
