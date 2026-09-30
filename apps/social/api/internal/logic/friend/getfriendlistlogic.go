// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package friend

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

type Get_friend_listLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 好友列表
func NewGet_friend_listLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Get_friend_listLogic {
	return &Get_friend_listLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Get_friend_listLogic) Get_friend_list(req *types.FriendListReq) (resp *types.FriendListResp, err error) {
	// todo: add your logic here and delete this line
	// 1.获取user_id
	userId, err := token.GetUserID(l.ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// 2.获取好友列表
	friends, err := l.svcCtx.SocialService.FriendList(l.ctx, &socialservice.FriendListRequest{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	ids := make([]string, 0, len(friends.Friends))
	for _, friend := range friends.Friends {
		ids = append(ids, friend.FriendId)
	}

	users, err := l.svcCtx.UserService.FindUser(l.ctx, &userservice.FindUserRequest{
		Ids: ids,
	})
	if err != nil {
		return nil, err
	}

	userMap := make(map[string]*userservice.User)
	for _, user := range users.Users {
		userMap[user.Id] = user
	}

	resp = &types.FriendListResp{
		Friends: make([]types.Friend, 0, len(friends.Friends)),
	}

	for _, friend := range friends.Friends {
		resp.Friends = append(resp.Friends, types.Friend{
			FriendId:  friend.FriendId,
			Avatar:    userMap[friend.FriendId].Avatar,
			Remark:    friend.Remark,
			AddSource: friend.AddSource,
		})
	}

	return
}
