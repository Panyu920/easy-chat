package logic

import (
	"context"

	"easy-chat/apps/im/rpc/im"
	"easy-chat/apps/im/rpc/internal/svc"
	"easy-chat/pkg/constant"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinGroupConversationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJoinGroupConversationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinGroupConversationLogic {
	return &JoinGroupConversationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 加入群聊会话
func (l *JoinGroupConversationLogic) JoinGroupConversation(in *im.JoinGroupConversationReq) (*im.JoinGroupConversationResp, error) {
	// todo: add your logic here and delete this line

	err := NewSetUpUserConversationLogic(l.ctx, l.svcCtx).SetUpUserOneConversation(in.GroupId, in.UserId, constant.ChatTypeGroup, false)
	if err != nil {
		return nil, err
	}

	return &im.JoinGroupConversationResp{
		GroupId: in.GroupId,
	}, nil
}
