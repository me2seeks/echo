package logic

import (
	"context"

	"github.com/me2seeks/echo-hub/app/counter/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/counter/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/common/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContentCounterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetContentCounterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContentCounterLogic {
	return &GetContentCounterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetContentCounterLogic) GetContentCounter(in *pb.GetContentCounterRequest) (*pb.GetContentCounterResponse, error) {
	// TODO use redis hash
	if !in.IsComment {
		feedCount, err := l.svcCtx.FeedCounterModel.FindOne(l.ctx, in.ID)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "find feed counter error: %v", err)
		}
		return &pb.GetContentCounterResponse{
			CommentCount: feedCount.CommentCount,
			LikeCount:    feedCount.LikeCount,
			ViewCount:    feedCount.ViewCount,
		}, nil
	}
	commentCount, err := l.svcCtx.CommentCounterModel.FindOne(l.ctx, in.ID)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "find comment counter error: %v", err)
	}

	return &pb.GetContentCounterResponse{
		CommentCount: commentCount.CommentCount,
		LikeCount:    commentCount.LikeCount,
		ViewCount:    commentCount.ViewCount,
	}, nil
}
