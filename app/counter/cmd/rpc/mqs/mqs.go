package mqs

import (
	"context"

	"github.com/me2seeks/echo-hub/app/counter/cmd/rpc/internal/config"
	"github.com/me2seeks/echo-hub/app/counter/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

//nolint:revive
func Consumers(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {
	return []service.Service{
		// Listening for changes in consumption flow status
		kq.MustNewQueue(c.KqConsumerConf, NewCounterEvent(ctx, svcContext)),
	}
}