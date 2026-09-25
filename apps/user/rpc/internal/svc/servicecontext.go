package svc

import (
	"easy-chat/apps/user/models"
	"easy-chat/apps/user/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	models.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.Mysql.DSN)
	return &ServiceContext{
		Config:     c,
		UsersModel: models.NewUsersModel(mysqlConn, c.Cache),
	}
}
