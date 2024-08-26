package logic

import (
	"context"

	"github.com/me2seeks/echo-hub/app/search/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/search/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUsersLogic {
	return &SearchUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUsersLogic) SearchUsers(in *pb.SearchReq) (*pb.SearchUsersResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchUsersResp{}, nil
}
