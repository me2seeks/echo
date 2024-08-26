package logic

import (
	"context"

	"github.com/me2seeks/echo-hub/app/search/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/search/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchContentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchContentLogic {
	return &SearchContentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchContentLogic) SearchContent(in *pb.SearchReq) (*pb.SearchContentResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchContentResp{}, nil
}
