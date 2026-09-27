package svc

import (
	"easy-chat/apps/social/rpc/internal/config"
	"easy-chat/apps/social/socialmodels"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	socialmodels.FriendsModel
	socialmodels.FriendsRequestsModel
	socialmodels.GroupsModel
	socialmodels.GroupMembersModel
	socialmodels.GroupRequestsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DSN)
	return &ServiceContext{
		Config:               c,
		FriendsModel:         socialmodels.NewFriendsModel(conn, c.Cache),
		FriendsRequestsModel: socialmodels.NewFriendsRequestsModel(conn, c.Cache),
		GroupsModel:          socialmodels.NewGroupsModel(conn, c.Cache),
		GroupMembersModel:    socialmodels.NewGroupMembersModel(conn, c.Cache),
		GroupRequestsModel:   socialmodels.NewGroupRequestsModel(conn, c.Cache),
	}
}
