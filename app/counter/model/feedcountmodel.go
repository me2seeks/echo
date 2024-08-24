package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FeedCountModel = (*customFeedCountModel)(nil)

type (
	// FeedCountModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFeedCountModel.
	FeedCountModel interface {
		feedCountModel
	}

	customFeedCountModel struct {
		*defaultFeedCountModel
	}
)

// NewFeedCountModel returns a model for the database table.
func NewFeedCountModel(conn sqlx.SqlConn, c cache.CacheConf) FeedCountModel {
	return &customFeedCountModel{
		defaultFeedCountModel: newFeedCountModel(conn, c),
	}
}
