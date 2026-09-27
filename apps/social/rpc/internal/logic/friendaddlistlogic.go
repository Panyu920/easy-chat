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

type FriendAddListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendAddListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendAddListLogic {
	return &FriendAddListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 好友添加列表服务
func (l *FriendAddListLogic) FriendAddList(in *social.FriendAddListRequest) (*social.FriendAddListResponse, error) {
	// todo: add your logic here and delete this line
	// 1. 查询好友添加列表
	friends, err := l.svcCtx.FriendsRequestsModel.ListFriendRequests(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBError(), "查询好友添加列表失败 err %v, id %v", err, in.UserId)
	}
	// 2. 转换为social.FriendRequest
	friendReqs := make([]*social.FriendRequest, 0, len(friends))
	for _, friend := range friends {
		if friend.ReqStatus == 0 {
			friendReqs = append(friendReqs, ConvertDBModelToSocialModel(friend))
		}
	}
	return &social.FriendAddListResponse{
		FriendRequests: friendReqs,
	}, nil

}

func ConvertDBModelToSocialModel(friends *socialmodels.FriendsRequests) *social.FriendRequest {
	return &social.FriendRequest{
		Id:        friends.Id,
		UserId:    friends.UserId,
		FriendId:  friends.FriendId,
		ReqMsg:    friends.ReqMsg.String,
		ReqTime:   friends.ReqTime.Unix(),
		ReqStatus: int32(friends.ReqStatus),
	}
}
