package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserLastRequestModel = (*customUserLastRequestModel)(nil)

type (
	// UserLastRequestModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserLastRequestModel.
	UserLastRequestModel interface {
		userLastRequestModel
	}

	customUserLastRequestModel struct {
		*defaultUserLastRequestModel
	}
)

// NewUserLastRequestModel returns a model for the database table.
func NewUserLastRequestModel(conn sqlx.SqlConn, c cache.CacheConf) UserLastRequestModel {
	return &customUserLastRequestModel{
		defaultUserLastRequestModel: newUserLastRequestModel(conn, c),
	}
}
