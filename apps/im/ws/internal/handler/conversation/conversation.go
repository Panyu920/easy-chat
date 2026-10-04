package conversation

import (
	"easy-chat/apps/im/ws/internal/svc"
	msgtype "easy-chat/apps/im/ws/msg_type"
	"easy-chat/apps/im/ws/websocketx"
	"easy-chat/apps/task/mq/mqtype"
	"easy-chat/pkg/constant"
	"fmt"

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

		if chatData.RecvId == "" {
			server.Logger.Error("recv id is empty")
			server.Send(websocketx.NewErrMessage("recv id is empty"), conn)
			return
		}
		userID := conn.GetUserID()
		combineKey := fmt.Sprintf("%s_%s", userID, msg.ClientMsgId)
		_, ok := server.MsgCache.Get(combineKey)
		if ok {
			// 消息已处理，直接返回成功
			server.Logger.Info("msg is cached", combineKey)
			server.Send(websocketx.NewAckMessage(websocketx.Result{
				ClientMsgId: msg.ClientMsgId,
				Status:      "success",
				Seq:         msg.Seq,
			}), conn)
			return
		} else {
			// cache 消息
			// 查数据库消息
		}

		switch chatData.ChatType {
		case constant.ChatTypeSingle:
			// 单聊
			// 发送消息到消息队列

			err := svc.MqClient.Push(&mqtype.MqChatType{
				ConversationId: chatData.ConversationId,
				FromID:         userID,
				SendId:         userID,
				RecvId:         chatData.RecvId,
				ChatType:       chatData.ChatType,
				MsgType:        chatData.MsgType,
				Content:        chatData.Msg.Content,
				SendTime:       chatData.SendTime,
			})

			if err != nil {
				server.Logger.Error(err)
				server.Send(websocketx.NewErrMessage("push chat to mq failed"), conn)
				return
			}

			// 缓存消息
			server.MsgCache.Put(combineKey, struct{}{})
			// 消息已处理，直接返回成功
			server.Logger.Info("msg is cached", combineKey)
			server.Send(websocketx.NewAckMessage(websocketx.Result{
				ClientMsgId: msg.ClientMsgId,
				Status:      "success",
				Seq:         msg.Seq,
			}), conn)
			// 存储聊天记录
			// err := logic.NewConversationLogic(context.Background(), svc, server).SingleChat(&chatData, conn.GetUserID())
			// if err != nil {
			// 	server.Logger.Error(err)
			// 	server.Send("save chat log failed", conn)
			// 	return
			// }

			// // 发送消息给接收者
			// err = server.SendToUser(websocketx.NewMessage(websocketx.FrameTypeData, msg.Method, conn.GetUserID(), msg.Data), chatData.RecvId)
			// server.Logger.Info("recv id", chatData.RecvId)
			// if err != nil {
			// 	server.Logger.Errorf("SendToUser failed err  %v ,from %s to %s,msgType  %v, msg  %s ", err, conn.GetUserID(), chatData.RecvId, chatData.MsgType, chatData.Content)
			// 	server.Send("SendToUser failed", conn)
			// 	return
			// }

			// return
		case constant.ChatTypeGroup:
			// 群聊
		default:
			// 其他类型
		}

	}
}
