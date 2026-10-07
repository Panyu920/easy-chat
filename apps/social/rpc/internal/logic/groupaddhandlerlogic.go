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
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type GroupAddHandlerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var (
	ErrGroupRequestNotExist    = xerr.New(xerr.SERVER_COMMON_ERROR, "群加入处理请求不存在")
	ErrGroupRequestStatusError = xerr.New(xerr.SERVER_COMMON_ERROR, "群加入处理请求状态错误")
	ErrGroupRequestHasHanled   = xerr.New(xerr.SERVER_COMMON_ERROR, "群加入处理请求已处理")
	ErrGroupRequestAuthError   = xerr.New(xerr.SERVER_COMMON_ERROR, "群加入处理请求权限错误")
	ErrGroupIDInvalid          = xerr.New(xerr.SERVER_COMMON_ERROR, "群ID无效")
)

func NewGroupAddHandlerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupAddHandlerLogic {
	return &GroupAddHandlerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 群加入处理服务
func (l *GroupAddHandlerLogic) GroupAddHandler(in *social.GroupAddHandlerRequest) (*social.GroupAddHandlerResponse, error) {
	// todo: add your logic here and delete this line
	// 1. 检查群加入处理请求是否存在
	println("in.Id", in.Id)
	req, err := l.svcCtx.GroupRequestsModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == socialmodels.ErrNotFound {
			return nil, errors.WithStack(ErrGroupRequestNotExist)
		}
		return nil, errors.Wrapf(xerr.NewDBError(), "查询群加入处理请求失败 err %v, id %v", err, in.Id)
	}
	if req.GroupId != in.GroupId {
		println("req.GroupId", req.GroupId)
		println("in.GroupId", in.GroupId)
		return nil, errors.WithStack(ErrGroupIDInvalid)
	}

	// 2. 检查群加入处理请求状态是否为待处理
	if req.ReqStatus != int64(constant.NoHandle) {
		return nil, errors.WithStack(ErrGroupRequestHasHanled)
	}
	// 3. 处理人是否为管理员或群主
	member, err := l.svcCtx.GroupMembersModel.FindOneByGroupIdUserId(l.ctx, in.GroupId, in.HandlerUserId)
	if err != nil {
		if err == socialmodels.ErrNotFound {
			return nil, errors.WithStack(ErrGroupRequestAuthError)
		}
		return nil, errors.Wrapf(xerr.NewDBError(), "查询群成员失败 err %v, id %v", err, req.HandlerUserId)
	}

	if member.RoleLevel != int64(constant.Admin) && member.RoleLevel != int64(constant.Owner) {
		return nil, errors.WithStack(ErrGroupRequestAuthError)
	}

	// 4. 处理群加入处理请求
	if in.ReqStatus == int32(constant.Pass) {
		// 同意
		err = l.svcCtx.GroupRequestsModel.CustomTx(l.ctx, func(ctx context.Context, conn sqlx.Session) error {
			// 1. 更新群加入处理请求状态为已同意
			req.ReqStatus = int64(constant.Pass)
			req.UpdateAt = time.Now()
			req.HandlerUserId = sql.NullString{String: in.HandlerUserId, Valid: true}
			err := l.svcCtx.GroupRequestsModel.UpdateTx(ctx, conn, req)
			if err != nil {
				return err
			}
			// 2. 插入群成员
			groupMember := socialmodels.GroupMembers{
				GroupId:       req.GroupId,
				UserId:        req.UserId,
				JoinTime:      time.Now(),
				JoinSource:    0,
				RoleLevel:     int64(constant.NormalMember),
				HandlerUserId: sql.NullString{String: in.HandlerUserId, Valid: true},
			}

			_, err = l.svcCtx.GroupMembersModel.InsertTx(ctx, conn, &groupMember)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return nil, errors.Wrapf(xerr.NewDBError(), "更新群加入处理请求失败 err %v, id %v", err, req.Id)
		}
	} else if in.ReqStatus == int32(constant.Refuse) {
		// 已拒绝
		req.ReqStatus = int64(constant.Refuse)
		req.UpdateAt = time.Now()
		req.HandlerUserId = sql.NullString{String: in.HandlerUserId, Valid: true}
		err = l.svcCtx.GroupRequestsModel.Update(l.ctx, req)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewDBError(), "更新群加入处理请求失败 err %v, id %v", err, req.Id)
		}
	} else {
		return nil, errors.WithStack(ErrGroupRequestAuthError)
	}

	return &social.GroupAddHandlerResponse{
		GroupId:   req.GroupId,
		ReqUserId: req.UserId,
	}, nil
}
