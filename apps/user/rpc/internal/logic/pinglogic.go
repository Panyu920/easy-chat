package logic

import (
	"context"

	"easy-chat/apps/user/rpc/internal/svc"
	"easy-chat/apps/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Ping is the RPC method that returns the pong.
func (l *PingLogic) Ping(in *user.PingRequest) (*user.PingResponse, error) {
	// todo: add your logic here and delete this line

	return &user.PingResponse{
		Pong: "pong",
	}, nil
}
