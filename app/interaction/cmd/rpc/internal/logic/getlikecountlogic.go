package logic

import (
	"context"

	"github.com/me2seeks/echo-hub/app/interaction/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/interaction/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/common/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLikeCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLikeCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLikeCountLogic {
	return &GetLikeCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLikeCountLogic) GetLikeCount(in *pb.GetLikeCountReq) (*pb.GetLikeCountResp, error) {
	var err error
	resp := &pb.GetLikeCountResp{}
	if in.IsComment {
		resp.Count, err = l.svcCtx.CommentLikesModel.FindCount(l.ctx, l.svcCtx.CommentLikesModel.SelectBuilder().Where("content_id = ?", in.Id), "id")
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "GetLikeCount find comment like count err:%v", err)
		}
	} else {
		resp.Count, err = l.svcCtx.FeedLikesModel.FindCount(l.ctx, l.svcCtx.FeedLikesModel.SelectBuilder().Where("content_id = ?", in.Id), "id")
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "GetLikeCount find feed like count err:%v", err)
		}
	}

	return resp, nil
}
