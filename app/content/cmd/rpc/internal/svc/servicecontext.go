package svc

import (
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/internal/config"
	"github.com/me2seeks/echo-hub/app/content/model"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	RedisClient    *redis.Redis
	CommentsModel  model.CommentsModel
	FeedsModel     model.FeedsModel
	KqPusherClient *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		CommentsModel:  model.NewCommentsModel(sqlConn, c.Cache),
		FeedsModel:     model.NewFeedsModel(sqlConn, c.Cache),
		KqPusherClient: kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
	}
}
