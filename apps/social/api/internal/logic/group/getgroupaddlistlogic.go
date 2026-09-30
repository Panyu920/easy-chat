// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package group

import (
	"context"

	"easy-chat/apps/social/api/internal/svc"
	"easy-chat/apps/social/api/internal/types"
	"easy-chat/apps/social/rpc/socialservice"
	"easy-chat/apps/user/rpc/userservice"
	"easy-chat/pkg/token"

	"github.com/pkg/errors"
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
	// // 1.获取userId
	userId, err := token.GetUserID(l.ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// // 2. 检查用户权限，是否是群主或群管理员
	// memberInfo ,err :=

	if "" == req.GroupId {
		return nil, errors.WithStack(ErrGroupIdEmpty)
	}
	temp_res, err := l.svcCtx.SocialService.GroupAddList(l.ctx, &socialservice.GroupAddListRequest{
		GroupId: req.GroupId,
		UserId:  userId,
	})
	if err != nil {
		return nil, err
	}
	if len(temp_res.GroupRequests) == 0 {
		return &types.GroupAddListResp{
			GroupRequests: []types.GroupRequest{},
		}, nil
	}
	userIds := make([]string, 0, len(temp_res.GroupRequests))
	for _, req := range temp_res.GroupRequests {
		userIds = append(userIds, req.UserId)
	}
	userInfos, err := l.svcCtx.UserService.FindUser(l.ctx, &userservice.FindUserRequest{
		Ids: userIds,
	})
	if err != nil {
		return nil, err
	}

	userMap := make(map[string]*userservice.User)
	for _, user := range userInfos.Users {
		userMap[user.Id] = user
	}

	groupRequests := make([]types.GroupRequest, 0, len(temp_res.GroupRequests))
	for _, req := range temp_res.GroupRequests {
		groupRequests = append(groupRequests, types.GroupRequest{
			Id:            req.Id,
			GroupId:       req.GroupId,
			UserId:        req.UserId,
			ReqMsg:        req.ReqMsg,
			ReqTime:       req.ReqTime,
			ReqStatus:     req.ReqStatus,
			JoinSource:    req.JoinSource,
			InviterUserId: req.InviterUserId,
			HandlerUserId: req.HandlerUserId,
			UserAvatar:    userMap[req.UserId].Avatar,
			UserName:      userMap[req.UserId].Nickname,
		})
	}
	resp = &types.GroupAddListResp{
		GroupRequests: groupRequests,
	}
	return
}
