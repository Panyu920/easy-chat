package logic

import (
	"context"

	"easy-chat/apps/social/rpc/internal/svc"
	"easy-chat/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupAddHandlerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupAddHandlerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupAddHandlerLogic {
	return &GroupAddHandlerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 群加入处理服务
func (l *GroupAddHandlerLogic) GroupAddHandler(in *social.GroupAddHandlerRequest) (*social.GroupAddHandlerResponse, error) {
	// todo: add your logic here and delete this line

	return &social.GroupAddHandlerResponse{}, nil
}
