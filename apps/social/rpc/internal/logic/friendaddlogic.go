package logic

import (
	"context"
	"database/sql"
	"time"

	"easy-chat/apps/social/rpc/internal/svc"
	"easy-chat/apps/social/rpc/social"
	"easy-chat/apps/social/socialmodels"

	"easy-chat/pkg/constant"
	xerr "easy-chat/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type FriendAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var (
	ErrFriendAlreadyExist = xerr.New(xerr.REQUEST_PARAM_ERROR, "已经是好友")
	ErrFriendApplyExist   = xerr.New(xerr.REQUEST_PARAM_ERROR, "正在申请好友")
)

func NewFriendAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendAddLogic {
	return &FriendAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 好友添加服务
func (l *FriendAddLogic) FriendAdd(in *social.FriendAddRequest) (*social.FriendAddResponse, error) {
	// todo: add your logic here and delete this line
	// 1. 检查是否已经是好友
	friend, err := l.svcCtx.FriendsModel.FindOneByUserIdFriendId(l.ctx, in.UserId, in.FriendId)
	if err != nil && err != socialmodels.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewDBError(), "查询好友关系失败 err %v, user id %v, friend id %v", err, in.UserId, in.FriendId)
	}
	if friend != nil {
		return nil, errors.WithStack(ErrFriendAlreadyExist)
	}
	// 2. 是否已经申请
	friendApply, err := l.svcCtx.FriendsRequestsModel.FindOneByUserIdFriendId(l.ctx, in.UserId, in.FriendId)
	if err != nil && err != socialmodels.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewDBError(), "查询好友申请失败 err %v, user id %v, friend id %v", err, in.UserId, in.FriendId)
	}
	if friendApply != nil && friendApply.ReqStatus == int64(constant.NoHandle) {
		return nil, errors.WithStack(ErrFriendApplyExist)
	}
	if friendApply != nil && friendApply.ReqStatus == int64(constant.Refuse) {
		// 已拒绝，重新申请
		friendApply.ReqStatus = 0
		friendApply.ReqMsg = sql.NullString{Valid: true, String: in.ReqMsg}
		friendApply.ReqTime = time.Unix(in.ReqTime, 0)
		friendApply.UpdateAt = time.Now()
		l.svcCtx.FriendsRequestsModel.Update(l.ctx, friendApply)
		return &social.FriendAddResponse{}, nil
	}
	// 3. 创建好友申请记录
	friendApply = &socialmodels.FriendsRequests{
		UserId:    in.UserId,
		FriendId:  in.FriendId,
		ReqMsg:    sql.NullString{Valid: true, String: in.ReqMsg},
		ReqTime:   time.Unix(in.ReqTime, 0),
		ReqStatus: 0,
	}
	if _, err := l.svcCtx.FriendsRequestsModel.Insert(l.ctx, friendApply); err != nil {
		return nil, errors.Wrapf(xerr.NewDBError(), "创建好友申请失败 err %v, user id %v, friend id %v", err, in.UserId, in.FriendId)
	}

	return &social.FriendAddResponse{}, nil
}
