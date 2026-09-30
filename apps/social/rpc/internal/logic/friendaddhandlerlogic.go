package logic

import (
	"context"
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

var (
	ErrFriendRequestNotExist    = xerr.New(xerr.SERVER_COMMON_ERROR, "好友请求不存在")
	ErrFriendRequestStatusError = xerr.New(xerr.SERVER_COMMON_ERROR, "好友请求已经处理")
	ErrFriendAuthError          = xerr.New(xerr.SERVER_COMMON_ERROR, "好友请求处理权限错误")
)

type FriendAddHandlerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendAddHandlerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendAddHandlerLogic {
	return &FriendAddHandlerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 好友添加处理服务
func (l *FriendAddHandlerLogic) FriendAddHandler(in *social.FriendAddHandlerRequest) (*social.FriendAddHandlerResponse, error) {
	// todo: add your logic here and delete this line
	// 1. 检查好友请求是否存在
	println("in.Id:", in.Id)
	req, err := l.svcCtx.FriendsRequestsModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == socialmodels.ErrNotFound {
			return nil, errors.WithStack(ErrFriendRequestNotExist)
		}
		return nil, errors.Wrapf(xerr.NewDBError(), "查询好友请求失败 err %v, id %v", err, in.Id)
	}
	println("req:", req)
	if req.FriendId != in.UserId {
		println(111111111111)
		println("req.UserId:", req.UserId)
		println("req.FriendId:", req.FriendId)
		println("req.ReqStatus:", req.ReqStatus)

		println("in.UserId:", in.UserId)
		return nil, errors.WithStack(ErrFriendAuthError)
	}

	// 2. 检查好友请求状态是否为待处理
	if req.ReqStatus != int64(constant.NoHandle) {
		return nil, errors.WithStack(ErrFriendRequestStatusError)
	}

	// 3. 处理好友请求

	if in.ReqStatus == int32(constant.Pass) {
		// 已同意
		err = l.svcCtx.FriendsRequestsModel.CustomTx(l.ctx, func(ctx context.Context, conn sqlx.Session) error {
			// 1. 更新好友请求状态为已同意
			req.ReqStatus = int64(constant.Pass)
			req.UpdateAt = time.Now()
			err := l.svcCtx.FriendsRequestsModel.UpdateTx(ctx, conn, req)
			if err != nil {
				return err
			}

			// 2. 插入 u1->u2
			user1 := socialmodels.Friends{
				UserId:    req.UserId,
				FriendId:  req.FriendId,
				AddSource: 0,
			}
			user2 := socialmodels.Friends{
				UserId:    req.FriendId,
				FriendId:  req.UserId,
				AddSource: 1,
			}

			_, err = l.svcCtx.FriendsModel.InsertsTx(ctx, conn, &user1, &user2)
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return nil, errors.Wrapf(xerr.NewDBError(), "更新好友请求失败 err %v, id %v", err, req.Id)
		}
	} else if in.ReqStatus == int32(constant.Refuse) {
		// 已拒绝
		req.ReqStatus = int64(constant.Refuse)
		req.UpdateAt = time.Now()
		err = l.svcCtx.FriendsRequestsModel.Update(l.ctx, req)

		if err != nil {
			return nil, errors.Wrapf(xerr.NewDBError(), "更新好友请求失败 err %v, id %v", err, req.Id)
		}
	} else {
		return nil, errors.WithStack(ErrFriendAuthError)
	}

	return &social.FriendAddHandlerResponse{}, nil
}
