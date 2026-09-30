// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package friend

import (
	"context"

	"easy-chat/apps/social/api/internal/svc"
	"easy-chat/apps/social/api/internal/types"
	"easy-chat/apps/social/rpc/socialservice"

	"easy-chat/pkg/token"
	"easy-chat/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrParam = xerr.New(xerr.REQUEST_PARAM_ERROR, "参数错误")
)

type Add_friendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加好友申请
func NewAdd_friendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Add_friendLogic {
	return &Add_friendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Add_friendLogic) Add_friend(req *types.FriendAddReq) (*types.FriendAddResp, error) {
	// todo: add your logic here and delete this line
	// 1.从token中获取用户ID
	println("FriendId:", req.FriendId, 11111111111111)
	userId, err := token.GetUserID(l.ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	println("userId:", userId)

	// 2. 参数校验
	if req.FriendId == userId || req.FriendId == "" {
		return nil, errors.WithStack(ErrParam)

	}
	if len(req.ReqMsg) == 0 {
		return nil, errors.WithStack(ErrParam)
	}
	if req.ReqTime == 0 {
		return nil, errors.WithStack(ErrParam)
	}
	_, err = l.svcCtx.SocialService.FriendAdd(l.ctx, &socialservice.FriendAddRequest{
		UserId:   userId,
		FriendId: req.FriendId,
		ReqMsg:   req.ReqMsg,
		ReqTime:  req.ReqTime,
	})
	if err != nil {
		return nil, err
	}
	return &types.FriendAddResp{}, nil
}
