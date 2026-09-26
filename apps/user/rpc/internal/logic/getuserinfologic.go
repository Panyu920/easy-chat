package logic

import (
	"context"

	"github.com/pkg/errors"

	"easy-chat/apps/user/models"
	"easy-chat/apps/user/rpc/internal/svc"
	"easy-chat/apps/user/rpc/user"
	"easy-chat/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetUserInfo is the RPC method that returns the user info.
func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	// todo: add your logic here and delete this line
	userInfo, err := l.svcCtx.UsersModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == models.ErrNotFound {
			return nil, errors.WithStack(ErrUserNotFound)
		}
		return nil, errors.Wrapf(xerr.NewDBError(), "查询用户失败: %v, by id: %s", err, in.Id)
	}

	return &user.GetUserInfoResponse{
		User: ConvertDbUserToRpcUser(userInfo),
	}, nil
}
