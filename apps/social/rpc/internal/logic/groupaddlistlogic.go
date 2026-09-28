package logic

import (
	"context"

	"easy-chat/apps/social/rpc/internal/svc"
	"easy-chat/apps/social/rpc/social"
	"easy-chat/apps/social/socialmodels"
	"easy-chat/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GroupAddListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupAddListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupAddListLogic {
	return &GroupAddListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 群加入列表服务
func (l *GroupAddListLogic) GroupAddList(in *social.GroupAddListRequest) (*social.GroupAddListResponse, error) {
	// todo: add your logic here and delete this line
	groupRequests, err := l.svcCtx.GroupRequestsModel.ListGroupRequests(l.ctx, in.GroupId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBError(), "list group requests failed err %v, group_id %s", err, in.GetGroupId())
	}
	socialRequests := make([]*social.GroupRequest, 0, len(groupRequests))
	for _, request := range groupRequests {
		socialRequests = append(socialRequests, ConvertDBRequestToSocialRequest(request))
	}
	return &social.GroupAddListResponse{
		GroupRequests: socialRequests,
	}, nil
}

func ConvertDBRequestToSocialRequest(dbRequest *socialmodels.GroupRequests) *social.GroupRequest {
	return &social.GroupRequest{
		Id:        dbRequest.Id,
		GroupId:   dbRequest.GroupId,
		UserId:    dbRequest.UserId,
		ReqMsg:    dbRequest.ReqMsg.String,
		ReqStatus: int32(dbRequest.ReqStatus),
		ReqTime:   dbRequest.ReqTime.Unix(),
	}

}
