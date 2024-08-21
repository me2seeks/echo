package svc

import (
	"github.com/me2seeks/echo-hub/app/interaction/cmd/api/internal/config"
	"github.com/me2seeks/echo-hub/app/interaction/cmd/rpc/interaction"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/usercenter"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	InteractionRPC interaction.Interaction
	UsercenterRPC  usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		InteractionRPC: interaction.NewInteraction(zrpc.MustNewClient(c.InteractionRPCConf)),
		UsercenterRPC:  usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRPCConf)),
	}
}
