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

	"easy-chat/apps/user/rpc/user"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type Get_friend_add_listLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取好友申请列表
func NewGet_friend_add_listLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Get_friend_add_listLogic {
	return &Get_friend_add_listLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Get_friend_add_listLogic) Get_friend_add_list(req *types.FriendAddListReq) (resp *types.FriendAddListResp, err error) {
	// todo: add your logic here and delete this line
	// 1.从token中获取用户ID
	userId, err := token.GetUserID(l.ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// 2. 调用好友服务获取好友申请列表
	res, err := l.svcCtx.SocialService.FriendAddList(l.ctx, &socialservice.FriendAddListRequest{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	ids := make([]string, 0, len(res.FriendRequests))
	for _, req := range res.FriendRequests {
		ids = append(ids, req.FriendId)
	}

	users, err := l.svcCtx.UserService.FindUser(l.ctx, &userservice.FindUserRequest{
		Ids: ids,
	})
	if err != nil {
		return nil, err
	}
	userMap := make(map[string]*user.User)
	for _, user := range users.Users {
		userMap[user.Id] = user
	}

	resp = &types.FriendAddListResp{
		FriendRequests: make([]types.FriendRequest, 0, len(res.FriendRequests)),
	}

	for _, req := range res.FriendRequests {
		resp.FriendRequests = append(resp.FriendRequests, types.FriendRequest{
			Id:         req.Id,
			UserId:     req.UserId,
			FriendId:   req.FriendId,
			ReqMsg:     req.ReqMsg,
			ReqTime:    req.ReqTime,
			ReqStatus:  req.ReqStatus,
			UserAvatar: userMap[req.FriendId].Avatar,
			UserName:   userMap[req.FriendId].Nickname,
		})
	}

	return
}
