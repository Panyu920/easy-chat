package mqclient

import (
	"context"
	"easy-chat/apps/task/mq/mqtype"
	"encoding/json"

	"github.com/zeromicro/go-queue/kq"
)

type IMqClient interface {
	Push(data *mqtype.MqChatType) error
}

type mqClientImpl struct {
	pusher *kq.Pusher
}

func NewMqClient(addr []string, topic string, options ...kq.PushOption) IMqClient {
	return &mqClientImpl{
		pusher: kq.NewPusher(addr, topic, options...),
	}
}

func (m *mqClientImpl) Push(data *mqtype.MqChatType) error {
	dataJ, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	return m.pusher.Push(context.Background(), string(dataJ))
}
