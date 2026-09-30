// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package friend

import (
	"context"

	"easy-chat/apps/social/api/internal/svc"
	"easy-chat/apps/social/api/internal/types"
	"easy-chat/apps/social/rpc/socialservice"
	"easy-chat/pkg/token"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type Handle_friend_addLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 处理好友申请（同意/拒绝）
func NewHandle_friend_addLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Handle_friend_addLogic {
	return &Handle_friend_addLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Handle_friend_addLogic) Handle_friend_add(req *types.FriendAddHandlerReq) (resp *types.FriendAddHandlerResp, err error) {
	// todo: add your logic here and delete this line
	// 1.从token中获取用户ID
	userId, err := token.GetUserID(l.ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	_, err = l.svcCtx.SocialService.FriendAddHandler(l.ctx, &socialservice.FriendAddHandlerRequest{
		UserId:    userId,
		Id:        req.Id,
		ReqStatus: req.ReqStatus,
	})
	if err != nil {
		return nil, err
	}

	return
}
