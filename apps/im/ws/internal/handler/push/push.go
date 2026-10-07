package push

import (
	"easy-chat/apps/im/ws/internal/svc"
	msgtype "easy-chat/apps/im/ws/msg_type"
	"easy-chat/apps/im/ws/websocketx"

	"github.com/mitchellh/mapstructure"
)

func Push(svc *svc.ServiceContext) websocketx.HandlerFunc {
	return func(server *websocketx.Server, conn *websocketx.Connection, msg *websocketx.Message) {
		// 处理聊天记录
		var data msgtype.Push
		err := mapstructure.Decode(msg.Data, &data)
		if err != nil {
			server.Logger.Errorf("decode chat failed: %v", err)
			return
		}
		// 发送聊天记录
		recvConn, _ := server.GetConn(data.RecvId)
		if recvConn == nil {
			return
		}
		// }else {
		// 	// 用户在其它网关登录
		// }

		chatData := msgtype.Chat{
			ConversationId: data.ConversationId,
			ChatType:       data.ChatType,
			SendId:         data.SendId,
			RecvId:         data.RecvId,
			SendTime:       data.SendTime,
			Msg: msgtype.Msg{
				MsgType: data.MsgType,
				Content: data.Content,
			},
		}
		msg = websocketx.NewMessage(websocketx.FrameTypeData, "push", data.FromID, chatData)
		server.Send(msg, recvConn)
	}
}
