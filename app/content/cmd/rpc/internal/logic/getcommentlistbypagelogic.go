package logic

import (
	"context"

	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentListByPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentListByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListByPageLogic {
	return &GetCommentListByPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentListByPageLogic) GetCommentListByPage(in *pb.GetCommentListByPageReq) (*pb.GetCommentListByPageResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetCommentListByPageResp{}, nil
}
