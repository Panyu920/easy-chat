// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package user

import (
	"context"

	"easy-chat/apps/user/api/internal/svc"
	"easy-chat/apps/user/api/internal/types"
	"easy-chat/apps/user/rpc/userservice"
	"easy-chat/pkg/token"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line
	// 1. 从token中获取用户id
	userID, err := token.GetUserID(l.ctx)
	if err != nil {
		return
	}
	// 2. 调用user rpc获取用户信息
	userInfo, err := l.svcCtx.UserService.GetUserInfo(l.ctx, &userservice.GetUserInfoRequest{
		Id: userID,
	})
	if err != nil {
		return
	}
	// 3. 返回用户信息
	resp = &types.UserInfoResp{
		Info: ConvertDBUserToRespUser(userInfo.User),
	}

	return
}

func ConvertDBUserToRespUser(dbUser *userservice.User) types.User {
	return types.User{
		Id:       dbUser.Id,
		Phone:    dbUser.Phone,
		Nickname: dbUser.Nickname,
		Sex:      byte(dbUser.Sex),
		Avatar:   dbUser.Avatar,
	}
}
