package logic

import (
	"context"

	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/app/content/model"

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
	err := l.svcCtx.FeedsModel.DeleteSoft(l.ctx, nil, &model.Feeds{
		Id: in.Id,
	})
	if err != nil {
		return nil, err
	}

	return &pb.DeleteFeedResp{}, nil
}