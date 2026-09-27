package logic

import (
	"context"

	"easy-chat/apps/social/rpc/internal/svc"
	"easy-chat/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendAddListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendAddListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendAddListLogic {
	return &FriendAddListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 好友添加列表服务
func (l *FriendAddListLogic) FriendAddList(in *social.FriendAddListRequest) (*social.FriendAddListResponse, error) {
	// todo: add your logic here and delete this line

	return &social.FriendAddListResponse{}, nil
}
