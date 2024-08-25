package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserStateModel = (*customUserStateModel)(nil)

type (
	// UserStateModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserStateModel.
	UserStateModel interface {
		userStateModel
	}

	customUserStateModel struct {
		*defaultUserStateModel
	}
)

// NewUserStateModel returns a model for the database table.
func NewUserStateModel(conn sqlx.SqlConn, c cache.CacheConf) UserStateModel {
	return &customUserStateModel{
		defaultUserStateModel: newUserStateModel(conn, c),
	}
}
