package mqs

// import (
// 	"context"
// 	"github.com/me2seeks/echo-hub/app/interaction/cmd/rpc/internal/svc"
// 	"github.com/zeromicro/go-zero/core/logx"
// )

// type RelationEvent struct {
// 	ctx context.Context
// 	svcCtx *svc.ServiceContext
// }

// func NewRelationEvent(ctx context.Context, svcCtx *svc.ServiceContext) *RelationEvent {
// 	return &RelationEvent{
// 		ctx: ctx,
// 		// svcCtx: svcCtx,
// 	}
// }

// func (l *RelationEvent) Consume(key, val string) error {
// 	logx.Infof("RelationEvent key :%s , val :%s", key, val)
// 	return nil
// }
