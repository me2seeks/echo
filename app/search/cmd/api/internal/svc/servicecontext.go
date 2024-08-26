package svc

import (
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/content"
	"github.com/me2seeks/echo-hub/app/search/cmd/api/internal/config"
	"github.com/me2seeks/echo-hub/app/search/cmd/rpc/search"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	SearchRPC  search.Search
	ContentRPC content.Content
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		SearchRPC:  search.NewSearch(zrpc.MustNewClient(c.SearchRPCConf)),
		ContentRPC: content.NewContent(zrpc.MustNewClient(c.ContentRPCConf)),
	}
}
