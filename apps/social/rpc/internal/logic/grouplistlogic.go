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

type GroupListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupListLogic {
	return &GroupListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 群列表服务
func (l *GroupListLogic) GroupList(in *social.GroupListRequest) (*social.GroupListResponse, error) {
	// todo: add your logic here and delete this line
	groupMembers, err := l.svcCtx.GroupMembersModel.ListGroupsByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBError(), "list groups failed: %v, user_id: %s", err, in.UserId)
	}
	groupIds := make([]string, 0, len(groupMembers))
	for _, member := range groupMembers {
		groupIds = append(groupIds, member.GroupId)
	}
	groups, err := l.svcCtx.GroupsModel.ListGroupsByIds(l.ctx, groupIds)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBError(), "list groups failed: %v, user_id: %s", err, in.UserId)
	}
	socialGroups := make([]*social.Group, 0, len(groups))
	for _, group := range groups {
		socialGroups = append(socialGroups, ConvertDBGroupToSocialGroup(group))
	}
	return &social.GroupListResponse{
		Groups: socialGroups,
	}, nil
}

func ConvertDBGroupToSocialGroup(dbGroup *socialmodels.Groups) *social.Group {
	return &social.Group{
		Id:                 dbGroup.Id,
		GroupName:          dbGroup.Name,
		GroupAvatar:        dbGroup.Avatar.String,
		GroupType:          int32(dbGroup.Type),
		IsVerify:           dbGroup.IsVerify == 1,
		CreateTime:         dbGroup.CreateAt.Unix(),
		CreateUserId:       dbGroup.CreateUserId,
		GroupStatus:        int32(dbGroup.Status),
		Notification:       dbGroup.Notification.String,
		NotificationUserId: dbGroup.NotificationUserId.String,
	}
}
