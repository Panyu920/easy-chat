package logic

import (
	"context"
	immodel "easy-chat/apps/im/im_model"
	"easy-chat/apps/im/ws/internal/svc"
	msgtype "easy-chat/apps/im/ws/msg_type"
	"easy-chat/apps/im/ws/websocketx"
	"easy-chat/pkg/wuid"
	"sync"
	"time"
)

type ConversationLogic struct {
	ctx context.Context
	svc *svc.ServiceContext
	srv *websocketx.Server
}

var singleToneConversationLogic ConversationLogic
var once sync.Once

func NewConversationLogic(ctx context.Context, svc *svc.ServiceContext, srv *websocketx.Server) *ConversationLogic {
	once.Do(func() {
		singleToneConversationLogic = ConversationLogic{
			ctx: ctx,
			svc: svc,
			srv: srv,
		}
	})
	return &singleToneConversationLogic
}

func (c *ConversationLogic) SingleChat(data *msgtype.Chat, userId string) error {
	if data.ConversationId == "" {
		data.ConversationId = wuid.CombineUserID(userId, data.RecvId)
	}

	err := c.svc.ChatLogModel.Insert(c.ctx, &immodel.ChatLog{
		ConversationId: data.ConversationId,
		SendId:         userId,
		RecvId:         data.RecvId,
		ChatType:       data.ChatType,
		MsgType:        data.MsgType,
		MsgContent:     data.Content,
		SendTime:       time.Now().Unix(),
		MsgFrom:        0,
		Status:         0,
	})
	if err != nil {
		c.srv.Logger.Errorf("Insert chat log failed: %v", err)
		return err
	}

	return nil
}
