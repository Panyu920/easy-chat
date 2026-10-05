package immodel

import (
	"easy-chat/pkg/constant"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Conversation struct {
	ID bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`

	ConversationId string            `bson:"conversation_id,omitempty"`
	ChatType       constant.ChatType `bson:"chat_type,omitempty"`
	// UserId         string            `bson:"user_id,omitempty"`
	IsShow bool `bson:"is_show,omitempty"`
	// Total          int               `bson:"total,omitempty"`
	Seq int64 `bson:"seq"`
	// Msg            *ChatLog          `bson:"msg,omitempty"`

	UpdateAt time.Time `bson:"update_at,omitempty" json:"update_at,omitempty"`
	CreateAt time.Time `bson:"create_at,omitempty" json:"create_at,omitempty"`
}
