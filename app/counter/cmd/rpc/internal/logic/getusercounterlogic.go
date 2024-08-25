package logic

import (
	"context"

	"github.com/me2seeks/echo-hub/app/counter/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/counter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserCounterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserCounterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserCounterLogic {
	return &GetUserCounterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserCounterLogic) GetUserCounter(in *pb.GetUserCounterRequest) (*pb.GetUserCounterResponse, error) {
	// TODO use redis hash

	return &pb.GetUserCounterResponse{}, nil
}
