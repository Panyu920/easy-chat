package logic

import (
	"context"

	"easy-chat/apps/social/rpc/internal/svc"
	"easy-chat/apps/social/rpc/social"
	"easy-chat/apps/social/socialmodels"
	"easy-chat/pkg/constant"
	"easy-chat/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrUserNotGroupMember = xerr.New(xerr.REQUEST_PARAM_ERROR, "用户不是群成员")
	ErrUserNotGroupAdmin  = xerr.New(xerr.REQUEST_PARAM_ERROR, "用户不是群管理员")
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
	// 1. 检查用户是否为群主或群管理员
	memberInfo, err := l.svcCtx.GroupMembersModel.FindOneByGroupIdUserId(l.ctx, in.GroupId, in.UserId)
	if err != nil {
		if socialmodels.ErrNotFound == err {
			return nil, errors.WithStack(ErrUserNotGroupMember)
		}
		return nil, errors.Wrapf(xerr.NewDBError(), "find group member failed err %v, group_id %s, user_id %s", err, in.GetGroupId(), in.GetUserId())
	}
	if memberInfo.RoleLevel != int64(constant.Admin) && memberInfo.RoleLevel != int64(constant.Owner) {
		return nil, errors.WithStack(ErrUserNotGroupAdmin)
	}
	// 2. 查询群加入列表
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
