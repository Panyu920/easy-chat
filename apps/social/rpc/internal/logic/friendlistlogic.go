package logic

import (
	"context"

	"easy-chat/apps/social/rpc/internal/svc"
	"easy-chat/apps/social/rpc/social"
	"easy-chat/apps/social/socialmodels"
	"easy-chat/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 好友列表服务
func (l *FriendListLogic) FriendList(in *social.FriendListRequest) (*social.FriendListResponse, error) {
	// todo: add your logic here and delete this line

	// 1.查询好友列表
	friends, err := l.svcCtx.FriendsModel.ListFriends(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBError(), "查询好友列表失败 err %v, id %v", err, in.UserId)
	}

	// 2. 转换为social.Friend
	socialFriends := make([]*social.Friend, 0, len(friends))
	for _, friend := range friends {
		socialFriends = append(socialFriends, ConvertDBFriendToSocialFriend(friend))
	}

	return &social.FriendListResponse{
		Friends: socialFriends,
	}, nil
}

func ConvertDBFriendToSocialFriend(dbFriend *socialmodels.Friends) *social.Friend {
	return &social.Friend{
		Id:        dbFriend.Id,
		UserId:    dbFriend.UserId,
		FriendId:  dbFriend.FriendId,
		AddSource: int32(dbFriend.AddSource),
		Remark:    dbFriend.Remark.String,
	}
}
