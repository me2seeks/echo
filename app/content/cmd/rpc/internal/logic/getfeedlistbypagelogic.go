package logic

import (
	"context"

	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFeedListByPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFeedListByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFeedListByPageLogic {
	return &GetFeedListByPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFeedListByPageLogic) GetFeedListByPage(in *pb.GetFeedListByPageReq) (*pb.GetFeedListByPageResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetFeedListByPageResp{}, nil
}
