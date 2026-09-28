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

type GroupMemberListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupMemberListLogic {
	return &GroupMemberListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 群成员列表服务
func (l *GroupMemberListLogic) GroupMemberList(in *social.GroupMemberListRequest) (*social.GroupMemberListResponse, error) {
	// todo: add your logic here and delete this line
	groupMembers, err := l.svcCtx.GroupMembersModel.ListMembersByGroupId(l.ctx, in.GroupId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBError(), "list group members failed: %v, group_id: %s", err, in.GroupId)
	}
	socialGroupMembers := make([]*social.GroupMember, 0, len(groupMembers))
	for _, member := range groupMembers {
		socialGroupMembers = append(socialGroupMembers, ConvertDBGroupMembersToSocialGroupMember(member))
	}

	return &social.GroupMemberListResponse{
		GroupMembers: socialGroupMembers,
	}, nil
}

func ConvertDBGroupMembersToSocialGroupMember(member *socialmodels.GroupMembers) *social.GroupMember {
	return &social.GroupMember{
		Id:        int64(member.Id),
		GroupId:   member.GroupId,
		UserId:    member.UserId,
		JoinTime:  int32(member.JoinTime.Unix()),
		RoleLevel: int32(member.RoleLevel),
	}
}
