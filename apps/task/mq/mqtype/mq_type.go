package mqtype

import "easy-chat/pkg/constant"

type MqChatType struct {
	ConversationId string            `json:"conversation_id"`
	ChatType       constant.ChatType `json:"chat_type"`
	FromID         string            `json:"from_id"`
	SendId         string            `json:"send_id"`
	RecvId         string            `json:"recv_id"`
	RecvIds        []string          `json:"recv_ids"`
	SendTime       int64             `json:"send_time"`
	MsgType        constant.MsgType  `json:"msg_type"`
	Content        string            `json:"content"`
	Seq            int64             `json:"seq"`
}
