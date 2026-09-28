package logic

import (
	"context"
	"database/sql"
	"time"

	"easy-chat/apps/social/rpc/internal/svc"
	"easy-chat/apps/social/rpc/social"
	"easy-chat/apps/social/socialmodels"
	"easy-chat/pkg/constant"
	"easy-chat/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrGroupNotFound      = xerr.NewError("用户群不存在")
	ErrUserAlreadyInGroup = xerr.NewError("用户已加入群聊")
	ErrUserAlreadyApply   = xerr.NewError("用户已申请加入群聊")
)

type GroupAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupAddLogic {
	return &GroupAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 群加入服务
func (l *GroupAddLogic) GroupAdd(in *social.GroupAddRequest) (*social.GroupAddResponse, error) {
	// todo: add your logic here and delete this line
	// 1.查找群是否存在
	group, err := l.svcCtx.GroupsModel.FindOne(l.ctx, in.GroupId)
	if err != nil {
		if err == socialmodels.ErrNotFound {
			return nil, errors.WithStack(ErrGroupNotFound)
		}
		return nil, errors.Wrapf(xerr.NewDBError(), "find group failed: %v, group_id: %s", err, in.GroupId)
	}

	// 2.用户是否已加入群聊
	groupMember, err := l.svcCtx.GroupMembersModel.FindOneByGroupIdUserId(l.ctx, in.GroupId, in.UserId)
	if err != nil && err != socialmodels.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewDBError(), "find group member failed: %v, group_id: %s, user_id: %s", err, in.GroupId, in.UserId)
	}
	if groupMember != nil {
		// 用户已加入群聊
		return nil, errors.WithStack(ErrUserAlreadyInGroup)
	}

	if group.IsVerify == 0 {
		// 普通群组，直接加入 todo
	}
	// 3.用户是否已经申请加入群聊
	groupApply, err := l.svcCtx.GroupRequestsModel.FindOneByGroupIdUserId(l.ctx, in.GroupId, in.UserId)
	if err != nil && err != socialmodels.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewDBError(), "find group apply failed: %v, group_id: %s, user_id: %s", err, in.GroupId, in.UserId)
	}
	if groupApply != nil {
		if groupApply.ReqStatus == int64(constant.NoHandle) {
			// 用户已申请加入群聊
			return nil, errors.WithStack(ErrUserAlreadyApply)
		} else if groupApply.ReqStatus == int64(constant.Refuse) {
			// 用户已加入群聊，但是申请状态为已拒绝
			// 重新申请加入群聊
			groupApply.ReqStatus = int64(constant.NoHandle)
			groupApply.ReqTime = time.Unix(in.ReqTime, 0)
			groupApply.UpdateAt = time.Now()
			groupApply.ReqMsg = sql.NullString{String: in.ReqMsg, Valid: true}
			err = l.svcCtx.GroupRequestsModel.Update(l.ctx, groupApply)
			if err != nil {
				return nil, errors.Wrapf(xerr.NewDBError(), "update group apply failed: %v, group_id: %s, user_id: %s", err, in.GroupId, in.UserId)
			} else {
				return &social.GroupAddResponse{GroupId: in.GroupId}, nil
			}
		} else {

		}

	}
	// 4.创建申请记录
	groupApply = &socialmodels.GroupRequests{
		GroupId:       in.GroupId,
		UserId:        in.UserId,
		ReqStatus:     int64(constant.NoHandle),
		ReqMsg:        sql.NullString{String: in.ReqMsg, Valid: true},
		ReqTime:       time.Unix(in.ReqTime, 0),
		JoinSource:    int64(in.JoinSource),
		InviterUserId: sql.NullString{String: in.InviterUserId},
	}
	_, err = l.svcCtx.GroupRequestsModel.Insert(l.ctx, groupApply)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBError(), "insert group apply failed: %v, group_id: %s, user_id: %s", err, in.GroupId, in.UserId)
	}
	return &social.GroupAddResponse{GroupId: in.GroupId}, nil
}
