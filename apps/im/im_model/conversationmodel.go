package immodel

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var _ ConversationModel = (*customConversationModel)(nil)

type (
	// ConversationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customConversationModel.
	ConversationModel interface {
		conversationModel
		ListByConversationIds(ctx context.Context, ids []string) ([]*Conversation, error)
		FindByConversationId(ctx context.Context, id string) (*Conversation, error)
		WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
	}

	customConversationModel struct {
		*defaultConversationModel
	}
)

// NewConversationModel returns a model for the mongo.
func NewConversationModel(url, db, collection string) ConversationModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customConversationModel{
		defaultConversationModel: newDefaultConversationModel(conn),
	}
}

func (m *customConversationModel) ListByConversationIds(ctx context.Context, ids []string) ([]*Conversation, error) {
	var conversations []*Conversation
	if len(ids) == 0 {
		return conversations, nil
	}

	err := m.conn.Find(ctx, &conversations, bson.M{"conversation_id": bson.M{"$in": ids}})

	switch err {
	case nil:
		return conversations, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customConversationModel) FindByConversationId(ctx context.Context, id string) (*Conversation, error) {
	var conversation Conversation

	err := m.conn.FindOne(ctx, &conversation, bson.M{"conversation_id": id})

	switch err {
	case nil:
		return &conversation, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// WithTransaction with transaction.
func (m *customConversationModel) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	session, err := m.conn.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx,
		func(ctx context.Context) (interface{}, error) {
			return nil, fn(ctx)
		},
	)
	if err != nil {
		return err
	}
	return nil
}
