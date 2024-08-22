package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysNotificationsModel = (*customSysNotificationsModel)(nil)

type (
	// SysNotificationsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysNotificationsModel.
	SysNotificationsModel interface {
		sysNotificationsModel
	}

	customSysNotificationsModel struct {
		*defaultSysNotificationsModel
	}
)

// NewSysNotificationsModel returns a model for the database table.
func NewSysNotificationsModel(conn sqlx.SqlConn, c cache.CacheConf) SysNotificationsModel {
	return &customSysNotificationsModel{
		defaultSysNotificationsModel: newSysNotificationsModel(conn, c),
	}
}
