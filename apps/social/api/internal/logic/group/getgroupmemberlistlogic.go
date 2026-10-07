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
	userId, err := token.GetUserID(l.ctx)
	if err != nil || "" == userId {
		return nil, err
	}

	if "" == req.GroupId {
		return nil, errors.WithStack(ErrGroupIdEmpty)
	}

	groupMembers, err := l.svcCtx.SocialService.GroupMemberList(l.ctx, &socialservice.GroupMemberListRequest{
		GroupId: req.GroupId,
		UserId:  userId,
	})
	if err != nil {
		return nil, err
	}

	userIds := make([]string, 0, len(groupMembers.GroupMembers))
	for _, member := range groupMembers.GroupMembers {
		userIds = append(userIds, member.UserId)
	}
	userInfoList, err := l.svcCtx.UserService.FindUser(l.ctx, &userservice.FindUserRequest{
		Ids: userIds,
	})
	if err != nil {
		return nil, err
	}
	userMap := make(map[string]*userservice.User, len(userInfoList.Users))
	for _, user := range userInfoList.Users {
		userMap[user.Id] = user
	}

	res := make([]types.GroupMember, 0, len(groupMembers.GroupMembers))
	for i, user := range groupMembers.GroupMembers {
		res = append(res, types.GroupMember{
			GroupId:    req.GroupId,
			UserId:     user.UserId,
			UserAvatar: userMap[user.UserId].Avatar,
			NickName:   user.Nickname,
			RoleLevel:  user.RoleLevel,
		})
		if "" == user.Nickname {
			res[i].NickName = userMap[user.UserId].Nickname
		}
	}
	resp = &types.GroupMemberListResp{
		GroupMembers: res,
	}

	return
}
