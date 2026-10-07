// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package group

import (
	"context"

	"easy-chat/apps/im/rpc/im"
	"easy-chat/apps/social/api/internal/svc"
	"easy-chat/apps/social/api/internal/types"
	"easy-chat/apps/social/rpc/socialservice"
	"easy-chat/apps/user/rpc/userservice"
	"easy-chat/pkg/xerr"

	"easy-chat/pkg/token"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrGroupNameEmpty   = xerr.New(xerr.REQUEST_PARAM_ERROR, "群名称不能为空")
	ErrGroupAvatarEmpty = xerr.New(xerr.REQUEST_PARAM_ERROR, "群头像不能为空")
	ErrGroupTypeEmpty   = xerr.New(xerr.REQUEST_PARAM_ERROR, "群类型不能为空")
	ErrIsVerifyEmpty    = xerr.New(xerr.REQUEST_PARAM_ERROR, "是否需要验证不能为空")
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
	userId, err := token.GetUserID(l.ctx)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewInternalError(), "get user id failed")
	}
	// 校验参数
	if req.GroupName == "" {
		return nil, ErrGroupNameEmpty
	}
	if req.GroupAvatar == "" {
		return nil, ErrGroupAvatarEmpty
	}

	// 获取创建者信息
	userInfo, err := l.svcCtx.UserService.GetUserInfo(l.ctx, &userservice.GetUserInfoRequest{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}

	gp, err := l.svcCtx.SocialService.GroupCreate(l.ctx, &socialservice.GroupCreateRequest{
		UserId:      userId,
		GroupName:   req.GroupName,
		GroupAvatar: req.GroupAvatar,
		GroupType:   req.GroupType,
		IsVerify:    req.IsVerify,
		Nickname:    userInfo.User.Nickname,
	})
	if err != nil {
		return nil, err
	}

	// 创建群组会话
	_, err = l.svcCtx.ImClient.CreateGroupConversation(l.ctx, &im.CreateGroupConversationReq{
		GroupId:  gp.GroupId,
		CreateId: userId,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.GroupCreateResp{
		GroupInfo: types.Group{
			Id:           gp.GroupId,
			GroupName:    req.GroupName,
			CreateUserId: userId,
			GroupAvatar:  req.GroupAvatar,
			GroupType:    req.GroupType,
			IsVerify:     req.IsVerify,
		},
	}
	return resp, nil
}
