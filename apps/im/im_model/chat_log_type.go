package immodel

import (
	"easy-chat/pkg/constant"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

var DefaultChatLogLimit int64 = 100

type ChatLog struct {
	//ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ID bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`

	ConversationId string            `bson:"conversation_id"`
	SendId         string            `bson:"send_id"`
	RecvId         string            `bson:"recv_id"`
	FrameFrom      string            `bson:"frame_from"`
	ChatType       constant.ChatType `bson:"chat_type"`
	MsgType        constant.MsgType  `bson:"msg_type"`
	MsgContent     string            `bson:"msg_content"`
	SendTime       int64             `bson:"send_time"`
	Status         int               `bson:"status"`

	// TODO: Fill your own fields
	UpdateAt time.Time `bson:"update_at,omitempty" json:"update_at,omitempty"`
	CreateAt time.Time `bson:"create_at,omitempty" json:"create_at,omitempty"`
}
