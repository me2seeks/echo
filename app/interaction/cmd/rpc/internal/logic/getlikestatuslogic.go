package logic

import (
	"context"

	"github.com/me2seeks/echo-hub/app/interaction/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/interaction/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLikeStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLikeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLikeStatusLogic {
	return &GetLikeStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLikeStatusLogic) GetLikeStatus(in *pb.GetLikeStatusReq) (*pb.GetLikeStatusResp, error) {
	resp := &pb.GetLikeStatusResp{
		IsLiked: false,
	}
	if !in.IsComment {
		feedLike, err := l.svcCtx.FeedLikesModel.FindOneByUserIdContentId(l.ctx, in.UserID, in.ContentID)
		if err == nil && feedLike != nil {
			resp.IsLiked = true
		}
	} else {
		commentLike, err := l.svcCtx.CommentLikesModel.FindOneByUserIdContentId(l.ctx, in.UserID, in.ContentID)
		if err == nil && commentLike != nil {
			resp.IsLiked = true
		}
	}
	return resp, nil
}
