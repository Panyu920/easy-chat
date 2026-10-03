package svc

import (
	"easy-chat/apps/user/models"
	"easy-chat/apps/user/rpc/internal/config"
	"time"

	"easy-chat/pkg/constant"
	"easy-chat/pkg/token"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	models.UsersModel
	Redisx *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.Mysql.DSN)
	return &ServiceContext{
		Config:     c,
		UsersModel: models.NewUsersModel(mysqlConn, c.Cache),
		Redisx:     redis.MustNewRedis(c.Redisx),
	}
}

func (s *ServiceContext) SetSystemRootToken() error {
	token, err := token.GenerateToken(s.Config.Jwt.AccessSecret, constant.REDIS_KEY_USER_ID, time.Now().Unix(), 9999999999)
	if err != nil {
		return err
	}
	return s.Redisx.Set(constant.REDIS_KEY_ROOT_SYSTEM, token)
}
