// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package group

import (
	"context"

	"easy-chat/apps/social/api/internal/svc"
	"easy-chat/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Get_group_listLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户所属群组列表
func NewGet_group_listLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Get_group_listLogic {
	return &Get_group_listLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Get_group_listLogic) Get_group_list(req *types.GroupListReq) (resp *types.GroupListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
