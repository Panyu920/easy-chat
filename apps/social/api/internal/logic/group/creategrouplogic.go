// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package group

import (
	"context"

	"easy-chat/apps/social/api/internal/svc"
	"easy-chat/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Create_groupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建群组
func NewCreate_groupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Create_groupLogic {
	return &Create_groupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Create_groupLogic) Create_group(req *types.GroupCreateReq) (resp *types.GroupCreateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
