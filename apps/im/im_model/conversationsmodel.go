package immodel

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var _ ConversationsModel = (*customConversationsModel)(nil)

type (
	// ConversationsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customConversationsModel.
	ConversationsModel interface {
		conversationsModel
		FindByUserId(ctx context.Context, userId string) (*Conversations, error)
		InsertMany(ctx context.Context, data ...*Conversations) error
	}

	customConversationsModel struct {
		*defaultConversationsModel
	}
)

// NewConversationsModel returns a model for the mongo.
func NewConversationsModel(url, db, collection string) ConversationsModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customConversationsModel{
		defaultConversationsModel: newDefaultConversationsModel(conn),
	}
}

func (m *customConversationsModel) FindByUserId(ctx context.Context, userId string) (*Conversations, error) {
	var conversations Conversations

	err := m.conn.FindOne(ctx, &conversations, bson.M{"user_id": userId})

	switch err {
	case nil:
		return &conversations, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customConversationsModel) InsertMany(ctx context.Context, data ...*Conversations) error {
	for _, item := range data {
		if item.ID.IsZero() {
			item.ID = bson.NewObjectID()
			item.CreateAt = time.Now()
			item.UpdateAt = time.Now()
		}
	}
	docs := make([]any, 0, len(data))
	for _, v := range data {
		docs = append(docs, v)
	}

	_, err := m.conn.InsertMany(ctx, docs)
	if err != nil {
		return err
	}
	return nil
}
