package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FeedLikesModel = (*customFeedLikesModel)(nil)

type (
	// FeedLikesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFeedLikesModel.
	FeedLikesModel interface {
		feedLikesModel
	}

	customFeedLikesModel struct {
		*defaultFeedLikesModel
	}
)

// NewFeedLikesModel returns a model for the database table.
func NewFeedLikesModel(conn sqlx.SqlConn, c cache.CacheConf) FeedLikesModel {
	return &customFeedLikesModel{
		defaultFeedLikesModel: newFeedLikesModel(conn, c),
	}
}
