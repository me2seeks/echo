package logic

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/me2seeks/echo-hub/app/counter/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/counter/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/app/counter/cmd/rpc/counter"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFeedCounterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// get feed counter
func NewGetFeedCounterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFeedCounterLogic {
	return &GetFeedCounterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFeedCounterLogic) GetFeedCounter(req *types.GetContentCounterReq) (*types.GetContentCounterResp, error) {
	getContentCounterResp, err := l.svcCtx.CounterRPC.GetContentCounter(l.ctx, &counter.GetContentCounterRequest{
		ID:        req.ID,
		IsComment: false,
	})
	if err != nil {
		return nil, err
	}
	resp := &types.GetContentCounterResp{}
	_ = copier.Copy(resp, getContentCounterResp)

	return resp, nil
}
