package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ContentCountModel = (*customContentCountModel)(nil)

type (
	// ContentCountModel is an interface to be customized, add more methods here,
	// and implement the added methods in customContentCountModel.
	ContentCountModel interface {
		contentCountModel
	}

	customContentCountModel struct {
		*defaultContentCountModel
	}
)

// NewContentCountModel returns a model for the database table.
func NewContentCountModel(conn sqlx.SqlConn, c cache.CacheConf) ContentCountModel {
	return &customContentCountModel{
		defaultContentCountModel: newContentCountModel(conn, c),
	}
}
