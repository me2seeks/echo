package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	BaseURL string
	Mysql   struct {
		DataSource string
	}
	Cache                    cache.CacheConf
	KqPusherCounterEventConf struct {
		Brokers []string
		Topic   string
	}
	KqPusherEsEventConf struct {
		Brokers []string
		Topic   string
	}
}
