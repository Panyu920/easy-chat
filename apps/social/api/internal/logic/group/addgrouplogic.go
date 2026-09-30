// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package group

import (
	"context"

	"easy-chat/apps/social/api/internal/svc"
	"easy-chat/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Add_groupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 申请加入群组
func NewAdd_groupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Add_groupLogic {
	return &Add_groupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Add_groupLogic) Add_group(req *types.GroupAddReq) (resp *types.GroupAddResp, err error) {
	// todo: add your logic here and delete this line

	return
}
