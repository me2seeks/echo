package logic

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/me2seeks/echo-hub/app/counter/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/counter/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/app/counter/cmd/rpc/counter"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContentCounterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// get content counter
func NewGetContentCounterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContentCounterLogic {
	return &GetContentCounterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetContentCounterLogic) GetContentCounter(req *types.GetContentCounterReq) (*types.GetContentCounterResp, error) {
	counterResp, err := l.svcCtx.CounterRPC.GetContentCounter(l.ctx, &counter.GetContentCounterRequest{
		ID:        req.ID,
		IsComment: req.IsComment,
	})
	if err != nil {
		return nil, err
	}
	var resp types.GetContentCounterResp
	_ = copier.Copy(&resp, counterResp)

	return &resp, nil
}
