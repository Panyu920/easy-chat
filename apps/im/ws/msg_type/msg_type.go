package msgtype

import "easy-chat/pkg/constant"

type (
	Msg struct {
		constant.MsgType `mapstructure:"msgType"`
		Content          string `mapstructure:"content"`
	}

	Chat struct {
		ConversationId    string `mapstructure:"conversationId"`
		constant.ChatType `mapstructure:"chatType"`
		SendId            string `mapstructure:"sendId"`
		RecvId            string `mapstructure:"recvId"`
		SendTime          int64  `mapstructure:"sendTime"`
		Msg               `mapstructure:"msg"`
	}
)
