package immodel

import (
	"context"
	"errors"
	"time"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ ConversationsModel = (*customConversationsModel)(nil)

type (
	// ConversationsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customConversationsModel.
	ConversationsModel interface {
		conversationsModel
		FindByUserId(ctx context.Context, userId string) (*Conversations, error)
		InsertMany(ctx context.Context, data ...*Conversations) error
		UpdateOneConversation(ctx context.Context, data *Conversations, conversationId string) (*mongo.UpdateResult, error)
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

func (m *defaultConversationsModel) UpdateOneConversation(ctx context.Context, data *Conversations, conversationId string) (*mongo.UpdateResult, error) {
	data.UpdateAt = time.Now()

	conv, ok := data.ConversationList[conversationId]
	if !ok {
		return nil, errors.New("conversation not found in data")
	}

	res, err := m.conn.UpdateOne(ctx,
		bson.M{"user_id": data.UserId},
		bson.M{"$set": bson.M{
			"conversation_list." + conversationId: conv, // ← 点号 + key
			"update_at":                           data.UpdateAt,
		}},
	)
	if err != nil {
		return nil, err
	}
	if res.MatchedCount == 0 {
		return res, ErrNotFound
	}
	return res, nil
}
