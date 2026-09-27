package logic

import (
	"context"

	"easy-chat/apps/social/rpc/internal/svc"
	"easy-chat/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupAddListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupAddListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupAddListLogic {
	return &GroupAddListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 群加入列表服务
func (l *GroupAddListLogic) GroupAddList(in *social.GroupAddListRequest) (*social.GroupAddListResponse, error) {
	// todo: add your logic here and delete this line

	return &social.GroupAddListResponse{}, nil
}
