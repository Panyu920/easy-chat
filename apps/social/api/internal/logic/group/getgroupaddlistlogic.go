// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package group

import (
	"context"

	"easy-chat/apps/social/api/internal/svc"
	"easy-chat/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Get_group_add_listLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取群申请列表
func NewGet_group_add_listLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Get_group_add_listLogic {
	return &Get_group_add_listLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Get_group_add_listLogic) Get_group_add_list(req *types.GroupAddListReq) (resp *types.GroupAddListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
