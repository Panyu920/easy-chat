package logic

import (
	"context"

	"easy-chat/apps/social/rpc/internal/svc"
	"easy-chat/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendAddLogic {
	return &FriendAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 好友添加服务
func (l *FriendAddLogic) FriendAdd(in *social.FriendAddRequest) (*social.FriendAddResponse, error) {
	// todo: add your logic here and delete this line

	return &social.FriendAddResponse{}, nil
}
