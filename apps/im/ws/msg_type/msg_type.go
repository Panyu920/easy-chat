package msgtype

import "easy-chat/pkg/constant"

type (
	Msg struct {
		constant.MsgType `mapstructure:"msg_type" json:"msg_type"`
		Content          string `mapstructure:"content" json:"content"`
	}

	Chat struct {
		ConversationId string            `mapstructure:"conversation_id" json:"conversation_id"`
		ChatType       constant.ChatType `mapstructure:"chat_type" json:"chat_type"`
		SendId         string            `mapstructure:"send_id" json:"send_id"`
		RecvId         string            `mapstructure:"recv_id" json:"recv_id"`
		SendTime       int64             `mapstructure:"send_time" json:"send_time"`
		Seq            int64             `mapstructure:"seq" json:"seq"`
		Msg            `mapstructure:"msg" json:"msg"`
	}
	Push struct {
		ConversationId string            `mapstructure:"conversation_id" json:"conversation_id"`
		ChatType       constant.ChatType `mapstructure:"chat_type" json:"chat_type,omitempty"`
		FromID         string            `mapstructure:"from_id" json:"from_id,omitempty"`
		SendId         string            `mapstructure:"send_id" json:"send_id,omitempty"`
		RecvId         string            `mapstructure:"recv_id" json:"recv_id,omitempty"`
		RecvIds        []string          `mapstructure:"recv_ids" json:"recv_ids,omitempty"`
		SendTime       int64             `mapstructure:"send_time" json:"send_time,omitempty"`
		MsgType        constant.MsgType  `mapstructure:"msg_type" json:"msg_type,omitempty"`
		Content        string            `mapstructure:"content" json:"content,omitempty"`
	}
)
