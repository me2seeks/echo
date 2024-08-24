package mqs

import (
	"context"

	"github.com/me2seeks/echo-hub/app/counter/cmd/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type CounterEvent struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCounterEvent(ctx context.Context, svcCtx *svc.ServiceContext) *CounterEvent {
	return &CounterEvent{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CounterEvent) Consume(ctx context.Context, key, val string) error {
	logx.Infof("CounterEvent key :%s , val :%s", key, val)
	return nil
}
