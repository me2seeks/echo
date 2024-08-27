package svc

import (
	"github.com/me2seeks/echo-hub/app/counter/cmd/consumer/internal/config"
	"github.com/me2seeks/echo-hub/app/counter/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config              config.Config
	CommentCounterModel model.CommentCountModel
	FeedCounterModel    model.FeedCountModel
	UserStateModel      model.UserStateModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:              c,
		CommentCounterModel: model.NewCommentCountModel(sqlConn, c.Cache),
		FeedCounterModel:    model.NewFeedCountModel(sqlConn, c.Cache),
		UserStateModel:      model.NewUserStateModel(sqlConn, c.Cache),
	}
}
