package logic

import (
	"context"

	immodel "easy-chat/apps/im/im_model"
	"easy-chat/apps/im/rpc/im"
	"easy-chat/apps/im/rpc/internal/svc"
	"easy-chat/pkg/constant"
	"easy-chat/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type PutConversationsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutConversationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutConversationsLogic {
	return &PutConversationsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新会话
func (l *PutConversationsLogic) PutConversations(in *im.PutConversationsReq) (*im.PutConversationsResp, error) {
	// todo: add your logic here and delete this line
	data, err := l.svcCtx.ConversationsModel.FindByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBError(), "ConversationsModel.FindByUserId err %v, req %v", err, in.UserId)
	}

	if data.ConversationList == nil {
		data.ConversationList = make(map[string]*immodel.Conversation, len(in.ConversationList))
	}

	for s, conversation := range in.ConversationList {
		data.ConversationList[s] = &immodel.Conversation{
			ConversationId: conversation.ConversationId,
			ChatType:       constant.ChatType(conversation.ChatType),
			IsShow:         conversation.IsShow,
			Seq:            conversation.Seq,
		}
	}

	_, err = l.svcCtx.ConversationsModel.Update(l.ctx, data)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBError(), "ConversationsModel.Update err %v, req %v", err, data)
	}

	return &im.PutConversationsResp{}, nil
}
