package svc

import (
	"github.com/me2seeks/echo-hub/app/counter/cmd/api/internal/config"
	"github.com/me2seeks/echo-hub/app/counter/cmd/rpc/counter"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	CounterRPC counter.Counter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		CounterRPC: counter.NewCounter(zrpc.MustNewClient(c.CounterRPCConf)),
	}
}
