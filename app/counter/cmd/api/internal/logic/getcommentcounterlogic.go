package logic

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/me2seeks/echo-hub/app/counter/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/counter/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/app/counter/cmd/rpc/counter"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentCounterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// get comment counter
func NewGetCommentCounterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentCounterLogic {
	return &GetCommentCounterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentCounterLogic) GetCommentCounter(req *types.GetContentCounterReq) (*types.GetContentCounterResp, error) {
	getContentCounterResp, err := l.svcCtx.CounterRPC.GetContentCounter(l.ctx, &counter.GetContentCounterRequest{
		ID:        req.ID,
		IsComment: true,
	})
	if err != nil {
		return nil, err
	}
	var resp *types.GetContentCounterResp
	_ = copier.Copy(&resp, getContentCounterResp)

	return resp, nil
}
