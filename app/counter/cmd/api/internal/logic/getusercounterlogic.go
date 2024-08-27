package logic

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/me2seeks/echo-hub/app/counter/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/counter/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/app/counter/cmd/rpc/counter"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserCounterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// get user counter
func NewGetUserCounterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserCounterLogic {
	return &GetUserCounterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserCounterLogic) GetUserCounter(req *types.GetUserCounterReq) (*types.GetUserCounterResp, error) {
	counterResp, err := l.svcCtx.CounterRPC.GetUserCounter(l.ctx, &counter.GetUserCounterRequest{
		UserId: req.UserID,
	})
	if err != nil {
		return nil, err
	}
	var resp types.GetUserCounterResp
	_ = copier.Copy(&resp, counterResp)

	return &resp, nil
}
