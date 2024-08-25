package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CommentLikesModel = (*customCommentLikesModel)(nil)

type (
	// CommentLikesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentLikesModel.
	CommentLikesModel interface {
		commentLikesModel
	}

	customCommentLikesModel struct {
		*defaultCommentLikesModel
	}
)

// NewCommentLikesModel returns a model for the database table.
func NewCommentLikesModel(conn sqlx.SqlConn, c cache.CacheConf) CommentLikesModel {
	return &customCommentLikesModel{
		defaultCommentLikesModel: newCommentLikesModel(conn, c),
	}
}
