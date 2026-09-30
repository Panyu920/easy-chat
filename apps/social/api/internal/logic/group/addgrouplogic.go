// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package group

import (
	"context"
	"time"

	"easy-chat/apps/social/api/internal/svc"
	"easy-chat/apps/social/api/internal/types"
	"easy-chat/apps/social/rpc/socialservice"
	"easy-chat/pkg/token"
	"easy-chat/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type Add_groupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var (
	ErrGroupIdEmpty = xerr.New(xerr.REQUEST_PARAM_ERROR, "群ID不能为空")
	ErrReqMsgEmpty  = xerr.New(xerr.REQUEST_PARAM_ERROR, "申请消息不能为空")
)

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

	// 1. 获取token中的用户ID
	userId, err := token.GetUserID(l.ctx)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewInternalError(), "get user id failed")
	}

	// 2. 校验参数
	if req.GroupId == "" {
		return nil, errors.WithStack(ErrGroupIdEmpty)
	}
	if req.ReqMsg == "" {
		return nil, errors.WithStack(ErrReqMsgEmpty)
	}

	// 3. 调用服务层
	gp, err := l.svcCtx.SocialService.GroupAdd(l.ctx, &socialservice.GroupAddRequest{
		UserId:        userId,
		GroupId:       req.GroupId,
		ReqMsg:        req.ReqMsg,
		ReqTime:       time.Now().Unix(),
		JoinSource:    0,
		InviterUserId: req.InviterUserId,
	})
	if err != nil {
		return nil, err
	}

	// 4. 返回结果
	resp = &types.GroupAddResp{
		GroupId: gp.GroupId,
	}

	return
}
