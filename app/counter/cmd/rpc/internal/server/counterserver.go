// Code generated by goctl. DO NOT EDIT.
// Source: counter.proto

package server

import (
	"context"

	"github.com/me2seeks/echo-hub/app/counter/cmd/rpc/internal/logic"
	"github.com/me2seeks/echo-hub/app/counter/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/counter/cmd/rpc/pb"
)

type CounterServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedCounterServer
}

func NewCounterServer(svcCtx *svc.ServiceContext) *CounterServer {
	return &CounterServer{
		svcCtx: svcCtx,
	}
}

func (s *CounterServer) GetContentCounter(ctx context.Context, in *pb.GetContentCounterRequest) (*pb.GetContentCounterResponse, error) {
	l := logic.NewGetContentCounterLogic(ctx, s.svcCtx)
	return l.GetContentCounter(in)
}

func (s *CounterServer) GetUserCounter(ctx context.Context, in *pb.GetUserCounterRequest) (*pb.GetUserCounterResponse, error) {
	l := logic.NewGetUserCounterLogic(ctx, s.svcCtx)
	return l.GetUserCounter(in)
}
