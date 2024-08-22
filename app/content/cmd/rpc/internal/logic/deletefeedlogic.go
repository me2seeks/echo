package logic

import (
	"context"

	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFeedLogic {
	return &DeleteFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteFeedLogic) DeleteFeed(in *pb.DeleteFeedReq) (*pb.DeleteFeedResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DeleteFeedResp{}, nil
}
