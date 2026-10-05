package immodel

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var _ ChatLogModel = (*customChatLogModel)(nil)

const (
	DefaultLimit int64 = 30
	MaxLimit     int64 = 100
)

type (
	// ChatLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChatLogModel.
	ChatLogModel interface {
		chatLogModel
		ListChatLogBySendTime(ctx context.Context, conversationId string, startTime, endTime, limit int64) ([]*ChatLog, error)
	}

	customChatLogModel struct {
		*defaultChatLogModel
	}
)

// NewChatLogModel returns a model for the mongo.
func NewChatLogModel(url, db, collection string) ChatLogModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customChatLogModel{
		defaultChatLogModel: newDefaultChatLogModel(conn),
	}
}

func (m *customChatLogModel) ListChatLogBySendTime(ctx context.Context, conversationId string, startTime, endTime, limit int64) ([]*ChatLog, error) {
	if limit > MaxLimit {
		limit = MaxLimit
	} else if limit <= 0 {
		limit = DefaultLimit
	}
	opts := options.Find().
		SetLimit(limit).
		SetSort(bson.M{"send_time": -1})

	filter := bson.M{
		"conversation_id": conversationId,
	}

	if endTime > 0 {
		filter["send_time"] = bson.M{
			"$lte": endTime,
			"$gte": startTime,
		}
	} else {
		filter["send_time"] = bson.M{
			"$lte": startTime,
		}
	}

	var data []*ChatLog

	err := m.conn.Find(ctx, &data, filter, opts)

	switch err {
	case nil:
		return data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
