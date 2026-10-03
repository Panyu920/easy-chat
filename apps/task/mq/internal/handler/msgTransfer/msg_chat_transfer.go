package msgtransfer

import (
	"context"
	"easy-chat/apps/task/mq/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MsgChatTransfer struct {
	svc *svc.ServiceContext
	logx.Logger
}

func NewMsgChatTransfer(svc *svc.ServiceContext) *MsgChatTransfer {
	return &MsgChatTransfer{
		svc:    svc,
		Logger: logx.WithContext(context.Background()),
	}
}

func (m *MsgChatTransfer) Consume(ctx context.Context, key, value string) error {
	m.Logger.Infof("MsgChatTransferSuccess key :%s , val :%s", key, value)

	return nil
}
