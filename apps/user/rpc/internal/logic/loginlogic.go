package logic

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"easy-chat/apps/user/models"
	"easy-chat/apps/user/rpc/internal/svc"
	"easy-chat/apps/user/rpc/user"
	"easy-chat/pkg/encrypt"
	"easy-chat/pkg/token"
	"easy-chat/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrUserHasRegistered = xerr.New(xerr.REQUEST_PARAM_ERROR, "用户已注册")
	ErrPasswordIncorrect = xerr.New(xerr.REQUEST_PARAM_ERROR, "密码错误")
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Login is the RPC method that logs a user in.
func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// todo: add your logic here and delete this line
	// 1. 校验手机号是否存在
	userInfo, err := l.svcCtx.UsersModel.FindOneByPhone(l.ctx, in.Phone)
	if err != nil && err != models.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewDBError(), "查询用户失败: %v, by phone: %s", err, in.Phone)
	}

	if userInfo == nil {
		return nil, errors.WithStack(ErrUserHasRegistered)
	}

	// 2. 校验密码是否正确
	if !encrypt.ValidatePasswordHash(userInfo.Password, in.Password) {
		return nil, errors.WithStack(ErrPasswordIncorrect)
	}

	// 3. 登录成功，返回token
	now := time.Now().Unix()
	accessToken, err := token.GenerateToken(l.svcCtx.Config.Jwt.AccessSecret, userInfo.Id, now, l.svcCtx.Config.Jwt.AccessExpire)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewInternalError(), "生成token失败: %v, user id: %s", err, userInfo.Id)
	}
	return &user.LoginResponse{
		Token:     accessToken,
		ExpiredAt: uint64(now + l.svcCtx.Config.Jwt.AccessExpire),
	}, nil
}
