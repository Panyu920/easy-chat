// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package friend

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

type Get_friend_online_listLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 好友在线列表
func NewGet_friend_online_listLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Get_friend_online_listLogic {
	return &Get_friend_online_listLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Get_friend_online_listLogic) Get_friend_online_list(req *types.FriendOnlineListReq) (resp *types.FriendOnlineListResp, err error) {
	// todo: add your logic here and delete this line
	// 1. 获取用户ID
	userId, err := token.GetUserID(l.ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// 1. 获取用户好友列表
	friendList, err := l.svcCtx.SocialService.FriendList(l.ctx, &socialservice.FriendListRequest{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// 2. 获取在线用户列表
	onlineUserList, err := l.svcCtx.RedisxClient.HgetallCtx(l.ctx, constant.USER_ONLINE_KEY)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// 3. 过滤好友在线列表
	var friendOnlineList = make([]types.UserOnline, 0, len(friendList.Friends))

	for _, friend := range friendList.Friends {
		if _, ok := onlineUserList[friend.UserId]; ok {
			friendOnlineList = append(friendOnlineList, types.UserOnline{
				UserId: friend.UserId,
				Online: true,
			})
		}
	}
	resp = &types.FriendOnlineListResp{}
	resp.FriendOnlineList = friendOnlineList

	return
}
