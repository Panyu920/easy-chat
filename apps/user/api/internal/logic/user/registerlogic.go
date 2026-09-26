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

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户注册
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line
	// 1. 调用user rpc注册用户
	registerResp, err := l.svcCtx.UserService.Register(l.ctx, &userservice.RegisterRequest{
		Phone:    req.Phone,
		Password: req.Password,
		Nickname: req.Nickname,
		Sex:      int32(req.Sex),
		Avatar:   req.Avatar,
	})

	if err != nil {
		return
	}
	// 2. 返回注册结果
	resp = &types.RegisterResp{
		Id: registerResp.Id,
	}
	return resp, nil
}
