package svc

import (
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/api/internal/config"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/usercenter"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	UsercenterRPC usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UsercenterRPC: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRPCConf)),
	}
}
