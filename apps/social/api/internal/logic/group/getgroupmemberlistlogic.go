// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package group

import (
	"context"

	"easy-chat/apps/social/api/internal/svc"
	"easy-chat/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Get_group_member_listLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 群成员列表
func NewGet_group_member_listLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Get_group_member_listLogic {
	return &Get_group_member_listLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Get_group_member_listLogic) Get_group_member_list(req *types.GroupMemberListReq) (resp *types.GroupMemberListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
