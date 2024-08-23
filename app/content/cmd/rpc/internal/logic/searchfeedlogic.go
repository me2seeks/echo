package logic

import (
	"context"

	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchFeedLogic {
	return &SearchFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchFeedLogic) SearchFeed(in *pb.SearchFeedReq) (*pb.SearchFeedResp, error) {
	// TODO es

	return &pb.SearchFeedResp{}, nil
}
