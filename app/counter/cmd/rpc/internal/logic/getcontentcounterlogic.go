package logic

import (
	"context"

	"github.com/me2seeks/echo-hub/app/counter/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/counter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContentCounterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetContentCounterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContentCounterLogic {
	return &GetContentCounterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetContentCounterLogic) GetContentCounter(in *pb.GetContentCounterRequest) (*pb.GetContentCounterResponse, error) {
	// TODO use redis hash

	return &pb.GetContentCounterResponse{}, nil
}
