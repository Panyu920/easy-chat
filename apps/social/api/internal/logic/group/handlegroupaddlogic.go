// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package group

import (
	"context"

	"easy-chat/apps/social/api/internal/svc"
	"easy-chat/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Handle_group_addLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 处理群申请
func NewHandle_group_addLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Handle_group_addLogic {
	return &Handle_group_addLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Handle_group_addLogic) Handle_group_add(req *types.GroupAddHandlerReq) (resp *types.GroupAddHandlerResp, err error) {
	// todo: add your logic here and delete this line

	return
}
