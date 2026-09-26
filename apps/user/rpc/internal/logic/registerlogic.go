package logic

import (
	"context"

	"easy-chat/apps/user/models"
	"easy-chat/apps/user/rpc/internal/svc"
	"easy-chat/apps/user/rpc/user"
	"easy-chat/pkg/encrypt"
	"easy-chat/pkg/wuid"
	"easy-chat/pkg/xerr"

	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrPhoneExist = xerr.New(xerr.REQUEST_PARAM_ERROR, "手机号已存在")
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Register is the RPC method that registers a user.
func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// todo: add your logic here and delete this line

	// 1.检查手机号是否存在
	userInfo, err := l.svcCtx.UsersModel.FindOneByPhone(l.ctx, in.Phone)
	// 错误
	if err != nil && err != models.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewDBError(), "查询用户失败: %v, phone: %s", err, in.Phone)
	}

	// 手机号已存在
	if userInfo != nil {
		return nil, errors.WithStack(ErrPhoneExist)
	}

	// 2.注册用户
	userId := wuid.GenerateUserID(l.svcCtx.Config.Mysql.DSN)
	hashPassword, err := encrypt.GneratePasswordHash(in.Password)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewInternalError(), "加密密码失败: %v, password: %s", err, in.Password)
	}
	userParam := &models.Users{
		Id:       userId,
		Avatar:   in.Avatar,
		Nickname: in.Nickname,
		Phone:    in.Phone,
		Password: hashPassword,
		Sex:      int64(in.Sex),
	}
	_, err = l.svcCtx.UsersModel.Insert(l.ctx, userParam)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBError(), "插入用户失败: %v, user id: %s", err, userId)
	}
	return &user.RegisterResponse{
		Id: userId,
	}, nil
}
