package svc

import (
	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/config"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/content"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/usercenter"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	ContentRPC    content.Content
	UsercenterRPC usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ContentRPC: content.NewContent(zrpc.MustNewClient(c.ContentRPCConf)),
		// UsercenterRPC: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRPCConf)),
	}
}
