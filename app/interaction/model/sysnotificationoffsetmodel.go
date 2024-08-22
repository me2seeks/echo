package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysNotificationOffsetModel = (*customSysNotificationOffsetModel)(nil)

type (
	// SysNotificationOffsetModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysNotificationOffsetModel.
	SysNotificationOffsetModel interface {
		sysNotificationOffsetModel
	}

	customSysNotificationOffsetModel struct {
		*defaultSysNotificationOffsetModel
	}
)

// NewSysNotificationOffsetModel returns a model for the database table.
func NewSysNotificationOffsetModel(conn sqlx.SqlConn, c cache.CacheConf) SysNotificationOffsetModel {
	return &customSysNotificationOffsetModel{
		defaultSysNotificationOffsetModel: newSysNotificationOffsetModel(conn, c),
	}
}
