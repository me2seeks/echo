// Code generated by goctl. DO NOT EDIT.
// Source: counter.proto

package server

import (
	// "github.com/me2seeks/echo-hub/app/counter/cmd/rpc/internal/logic"
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