package immodel

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type ChatLogT struct {
	ID bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	// TODO: Fill your own fields
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
