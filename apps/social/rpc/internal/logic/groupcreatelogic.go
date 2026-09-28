package logic

import (
	"context"
	"database/sql"
	"time"

	"easy-chat/apps/social/rpc/internal/svc"
	"easy-chat/apps/social/rpc/social"
	"easy-chat/apps/social/socialmodels"
	"easy-chat/pkg/constant"
	muuid "easy-chat/pkg/m_uuid"
	"easy-chat/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type GroupCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupCreateLogic {
	return &GroupCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 群创建服务
func (l *GroupCreateLogic) GroupCreate(in *social.GroupCreateRequest) (*social.GroupCreateResponse, error) {
	// todo: add your logic here and delete this line
	// 需要开启事务
	groupID, err := muuid.GenUUIDv7()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewError("gen group id failed"), "gen group id failed: %v", err)
	}

	isVerify := 0
	if in.IsVerify {
		isVerify = 1
	} else {
		isVerify = 0
	}
	err = l.svcCtx.GroupsModel.CustomTx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		// 插入群组
		_, err := l.svcCtx.GroupsModel.InsertTx(ctx, session, &socialmodels.Groups{
			Id:           groupID,
			CreateUserId: in.UserId,
			Name:         in.GroupName,
			Avatar:       sql.NullString{String: in.GroupAvatar},
			Type:         int64(in.GroupType),
			IsVerify:     int64(isVerify),
		})
		if err != nil {
			return err
		}
		// 插入群成员
		_, err = l.svcCtx.GroupMembersModel.InsertTx(ctx, session, &socialmodels.GroupMembers{
			GroupId:    groupID,
			UserId:     in.UserId,
			JoinSource: 0,
			RoleLevel:  int64(constant.Owner),
			JoinTime:   time.Now(),
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBError(), "create group failed: %v, user_id: %s", err, in.UserId)
	}

	return &social.GroupCreateResponse{
		GroupId: groupID,
	}, nil
}
