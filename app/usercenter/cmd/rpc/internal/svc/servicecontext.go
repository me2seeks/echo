package svc

import (
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/internal/config"
	"github.com/me2seeks/echo-hub/app/usercenter/model"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	RedisClient          *redis.Redis
	UserModel            model.UserModel
	UserAuthModel        model.UserAuthModel
	UserRelationModel    model.UserRelationModel
	UserLastRequestModel model.UserLastRequestModel

	KqPusherCounterEventClient *kq.Pusher
	KqPusherEsEventClient      *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		UserModel:            model.NewUserModel(sqlConn, c.Cache),
		UserAuthModel:        model.NewUserAuthModel(sqlConn, c.Cache),
		UserRelationModel:    model.NewUserRelationModel(sqlConn, c.Cache),
		UserLastRequestModel: model.NewUserLastRequestModel(sqlConn, c.Cache),

		KqPusherCounterEventClient: kq.NewPusher(c.KqPusherCounterEventConf.Brokers, c.KqPusherCounterEventConf.Topic),
		KqPusherEsEventClient:      kq.NewPusher(c.KqPusherEsEventConf.Brokers, c.KqPusherEsEventConf.Topic),
	}
}
