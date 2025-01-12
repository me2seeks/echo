package svc

import (
	"github.com/me2seeks/echo-hub/app/interaction/cmd/rpc/internal/config"
	"github.com/me2seeks/echo-hub/app/interaction/model"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config            config.Config
	RedisClient       *redis.Redis
	FeedLikesModel    model.FeedLikesModel
	CommentLikesModel model.CommentLikesModel
	KqPusherClient    *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		FeedLikesModel:    model.NewFeedLikesModel(sqlConn, c.Cache),
		CommentLikesModel: model.NewCommentLikesModel(sqlConn, c.Cache),
		KqPusherClient:    kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
	}
}
