package logic

import (
	"context"

	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFeedLogic {
	return &UpdateFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateFeedLogic) UpdateFeed(in *pb.UpdateFeedReq) (*pb.UpdateFeedResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateFeedResp{}, nil
}
