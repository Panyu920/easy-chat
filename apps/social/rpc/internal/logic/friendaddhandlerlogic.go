package logic

import (
	"context"

	"easy-chat/apps/social/rpc/internal/svc"
	"easy-chat/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendAddHandlerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendAddHandlerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendAddHandlerLogic {
	return &FriendAddHandlerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 好友添加处理服务
func (l *FriendAddHandlerLogic) FriendAddHandler(in *social.FriendAddHandlerRequest) (*social.FriendAddHandlerResponse, error) {
	// todo: add your logic here and delete this line

	return &social.FriendAddHandlerResponse{}, nil
}
