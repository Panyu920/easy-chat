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

type CreateGroupConversationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateGroupConversationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupConversationLogic {
	return &CreateGroupConversationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建群聊会话
func (l *CreateGroupConversationLogic) CreateGroupConversation(in *im.CreateGroupConversationReq) (*im.CreateGroupConversationResp, error) {
	// todo: add your logic here and delete this line
	// 1.查询群聊会话是否存在
	conversation, err := l.svcCtx.ConversationModel.FindByConversationId(l.ctx, in.GroupId)
	if err != nil && err != immodel.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewDBError(), "ConversationModel.FindByConversationId failed err %v ,groupId %v", err, in.GroupId)
	}
	if conversation != nil {
		return nil, nil
	}

	// 2.创建群聊会话
	err = l.svcCtx.ConversationModel.WithTransaction(l.ctx, func(ctx context.Context) error {
		conversation := &immodel.Conversation{
			ConversationId: in.GroupId,
			ChatType:       constant.ChatTypeGroup,
			Seq:            0,
		}
		err := l.svcCtx.ConversationModel.Insert(ctx, conversation)
		if err != nil {
			return errors.Wrapf(xerr.NewDBError(), "ConversationModel.Insert failed err %v ,groupId %v", err, in.GroupId)
		}

		err = NewSetUpUserConversationLogic(ctx, l.svcCtx).SetUpUserOneConversation(in.GroupId, in.CreateId, conversation.ChatType, false)
		if err != nil {
			return errors.Wrapf(xerr.NewDBError(), "SetUpUserOneConversation failed err %v ,groupId %v", err, in.GroupId)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return &im.CreateGroupConversationResp{}, nil
}
