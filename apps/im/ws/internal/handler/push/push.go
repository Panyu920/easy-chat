package push

import (
	"easy-chat/apps/im/ws/internal/svc"
	msgtype "easy-chat/apps/im/ws/msg_type"
	"easy-chat/apps/im/ws/websocketx"
	"easy-chat/pkg/constant"

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

		// }else {
		// 	// 用户在其它网关登录
		// }

		if data.ChatType == constant.ChatTypeSingle {
			singlePush(server, data, data.FromID, data.RecvId)
		} else if data.ChatType == constant.ChatTypeGroup {
			groupPush(server, data, data.FromID)
		}

	}
}

func singlePush(server *websocketx.Server, data msgtype.Push, fromID string, recvID string) {
	// 发送聊天记录
	recvConn, _ := server.GetConn(recvID)
	if recvConn == nil {
		// 用户未登录
		return
	}
	chatData := msgtype.Chat{
		ConversationId: data.ConversationId,
		ChatType:       data.ChatType,
		SendId:         data.SendId,
		RecvId:         recvID,
		SendTime:       data.SendTime,
		Seq:            data.Seq,
		Msg: msgtype.Msg{
			MsgType: data.MsgType,
			Content: data.Content,
		},
	}
	msg := websocketx.NewMessage(websocketx.FrameTypeData, "push", fromID, chatData)
	server.Send(msg, recvConn)
}

func groupPush(server *websocketx.Server, data msgtype.Push, fromID string) {
	// 发送群聊消息
	for _, id := range data.RecvIds {
		func(id string) {
			server.Scheduler.Schedule(func() {
				singlePush(server, data, fromID, id)
			})
		}(id)
	}
}
