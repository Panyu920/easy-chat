package logic

import (
	"context"
	"errors"

	"easy-chat/apps/user/models"
	"easy-chat/apps/user/rpc/internal/svc"
	"easy-chat/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func ConvertDbUserToRpcUser(dbUser *models.Users) *user.User {
	return &user.User{
		Id:       dbUser.Id,
		Phone:    dbUser.Phone,
		Nickname: dbUser.Nickname,
		Avatar:   dbUser.Avatar,
		Sex:      int32(dbUser.Sex),
	}
}

func NewFindUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLogic {
	return &FindUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FindUser is the RPC method that finds a user.
func (l *FindUserLogic) FindUser(in *user.FindUserRequest) (*user.FindUserResponse, error) {
	// todo: add your logic here and delete this line
	var users []models.Users
	var err error

	if in.Phone != "" {
		res, err := l.svcCtx.UsersModel.FindOneByPhone(l.ctx, in.Phone)
		if err != nil {
			if err == models.ErrNotFound {
				return nil, errors.New("用户不存在")
			}
			return nil, err
		}
		users = append(users, *res)
	} else if in.Nickname != "" {
		users, err = l.svcCtx.UsersModel.ListByNickname(l.ctx, in.Nickname)
		if err != nil {
			if err == models.ErrNotFound {
				return nil, errors.New("用户不存在")
			}
			return nil, err
		}
	} else if len(in.Ids) > 0 {
		users, err = l.svcCtx.UsersModel.ListByIds(l.ctx, in.Ids)
		if err != nil {
			if err == models.ErrNotFound {
				return nil, errors.New("用户不存在")
			}
			return nil, err
		}
	}

	// 转换为 RPC 用户列表
	rpcUsers := make([]*user.User, 0, len(users))
	for _, dbUser := range users {
		rpcUsers = append(rpcUsers, ConvertDbUserToRpcUser(&dbUser))
	}

	return &user.FindUserResponse{
		Users: rpcUsers,
	}, nil
}
