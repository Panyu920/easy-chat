// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package group

import (
	"context"

	"easy-chat/apps/social/api/internal/svc"
	"easy-chat/apps/social/api/internal/types"
	"easy-chat/apps/social/rpc/socialservice"
	"easy-chat/pkg/constant"
	"easy-chat/pkg/token"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type Get_group_member_online_listLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 群成员在线列表
func NewGet_group_member_online_listLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Get_group_member_online_listLogic {
	return &Get_group_member_online_listLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Get_group_member_online_listLogic) Get_group_member_online_list(req *types.GroupMemberOnlineListReq) (resp *types.GroupMemberOnlineListResp, err error) {
	// todo: add your logic here and delete this line
	resp = &types.GroupMemberOnlineListResp{}
	userId, err := token.GetUserID(l.ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// 1. 获取群成员列表
	groupMemberList, err := l.svcCtx.SocialService.GroupMemberList(l.ctx, &socialservice.GroupMemberListRequest{
		GroupId: req.GroupId,
		UserId:  userId,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// 2. 获取在线用户列表
	onlineUserList, err := l.svcCtx.RedisxClient.HgetallCtx(l.ctx, constant.USER_ONLINE_KEY)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// 3. 过滤群成员在线列表
	var groupMemberOnlineList = make([]types.UserOnline, 0, len(groupMemberList.GroupMembers))

	for _, member := range groupMemberList.GroupMembers {
		if _, ok := onlineUserList[member.UserId]; ok {
			groupMemberOnlineList = append(groupMemberOnlineList, types.UserOnline{
				UserId: member.UserId,
				Online: true,
			})
		}
	}
	resp = &types.GroupMemberOnlineListResp{}
	resp.GroupMemberOnlineList = groupMemberOnlineList
	return
}
