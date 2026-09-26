// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package user

import (
	"context"

	"easy-chat/apps/user/api/internal/svc"
	"easy-chat/apps/user/api/internal/types"
	"easy-chat/apps/user/rpc/userservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登入
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	// 1. 调用user rpc登入用户
	loginResp, err := l.svcCtx.UserService.Login(l.ctx, &userservice.LoginRequest{
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		return
	}
	// 2. 返回登入结果
	resp = &types.LoginResp{
		Token:  loginResp.Token,
		Expire: int64(loginResp.ExpiredAt),
	}
	return resp, nil
}
