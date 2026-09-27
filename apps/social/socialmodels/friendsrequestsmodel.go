package socialmodels

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FriendsRequestsModel = (*customFriendsRequestsModel)(nil)

type (
	// FriendsRequestsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFriendsRequestsModel.
	FriendsRequestsModel interface {
		friendsRequestsModel
	}

	customFriendsRequestsModel struct {
		*defaultFriendsRequestsModel
	}
)

// NewFriendsRequestsModel returns a model for the database table.
func NewFriendsRequestsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) FriendsRequestsModel {
	return &customFriendsRequestsModel{
		defaultFriendsRequestsModel: newFriendsRequestsModel(conn, c, opts...),
	}
}
