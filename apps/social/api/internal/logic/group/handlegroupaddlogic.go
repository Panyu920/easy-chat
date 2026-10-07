// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package group

import (
	"context"

	"easy-chat/apps/im/rpc/im"
	"easy-chat/apps/social/api/internal/svc"
	"easy-chat/apps/social/api/internal/types"
	"easy-chat/apps/social/rpc/socialservice"
	"easy-chat/pkg/constant"
	"easy-chat/pkg/token"
	"easy-chat/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

// 处理群申请
var (
	ErrReqStstusInvalid = xerr.New(xerr.REQUEST_PARAM_ERROR, "req status is invalid")
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
	// 校验参数
	if "" == req.GroupId {
		return nil, errors.WithStack(ErrGroupIdEmpty)
	}
	if req.ReqStatus != int32(constant.Pass) && req.ReqStatus != int32(constant.Refuse) {
		return nil, errors.WithStack(ErrReqStstusInvalid)
	}
	userId, err := token.GetUserID(l.ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	res, err := l.svcCtx.SocialService.GroupAddHandler(l.ctx, &socialservice.GroupAddHandlerRequest{
		GroupId:       req.GroupId,
		ReqStatus:     req.ReqStatus,
		HandlerUserId: userId,
		Id:            req.Id,
	})
	if err != nil {
		return nil, err
	}
	// if pass, then create group conversation
	if req.ReqStatus == int32(constant.Pass) {
		_, err = l.svcCtx.ImClient.JoinGroupConversation(l.ctx, &im.JoinGroupConversationReq{
			GroupId: req.GroupId,
			UserId:  res.ReqUserId,
		})
		if err != nil {
			return nil, err
		}
	}
	resp = &types.GroupAddHandlerResp{
		GroupId: req.GroupId,
	}
	return
}
