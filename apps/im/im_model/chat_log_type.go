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

	ConversationId string            `bson:"conversationId"`
	SendId         string            `bson:"sendId"`
	RecvId         string            `bson:"recvId"`
	MsgFrom        int               `bson:"msgFrom"`
	ChatType       constant.ChatType `bson:"chatType"`
	MsgType        constant.MsgType  `bson:"msgType"`
	MsgContent     string            `bson:"msgContent"`
	SendTime       int64             `bson:"sendTime"`
	Status         int               `bson:"status"`

	// TODO: Fill your own fields
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
